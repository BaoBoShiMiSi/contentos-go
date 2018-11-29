package blocklog

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/coschain/contentos-go/common"
)

const indexSize = 8
const blockLenSize = 4
const maxPayloadLen = 1024 * 1024 * 256

/*BLog is an external append only log of the blocks. Blocks should only be written
 * to the log after they're irreversible as the log is append only. There is a secondary
 * index file of only block positions that enables O(1) random access lookup by block number.
 *
 * A block data in the BLog is formatted as len+payload, len is a uint32
 *
 * +---------+----------------+---------+----------------+-----+------------+-------------------+
 * | Block 1 | Pos of Block 1 | Block 2 | Pos of Block 2 | ... | Head Block | Pos of Head Block |
 * +---------+----------------+---------+----------------+-----+------------+-------------------+
 *
 * +----------------+----------------+-----+-------------------+
 * | Pos of Block 1 | Pos of Block 2 | ... | Pos of Head Block |
 * +----------------+----------------+-----+-------------------+
 *
 *
 * Blocks can be accessed at random via block number through the index file. Seek to 8 * (block_num - 1)
 * to find the position of the block in the main file.
 *
 * The main file is the only file that needs to persist. The index file can be reconstructed during a
 * linear scan of the main file.
 */
type BLog struct {
	logFile   *os.File
	indexFile *os.File
}

// Open opens the block log & index file
func (bl *BLog) Open(dir string) (err error) {
	_, err = os.Stat(dir)

	if err != nil {
		os.Mkdir(dir, os.ModePerm)
	}

	bl.logFile, err = os.OpenFile(dir+"/block.bin", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	bl.indexFile, err = os.OpenFile(dir+"/block.index", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}

	logInfo, err := bl.logFile.Stat()
	if err != nil {
		return err
	}
	indexInfo, err := bl.indexFile.Stat()
	if err != nil {
		return err
	}

	if logInfo.Size() != 0 {
		if indexInfo.Size() != 0 {
			indexByte := make([]byte, indexSize)
			lastIdxFromLogFile, err := bl.readLastIndex(indexByte, true)
			if err != nil {
				return err
			}
			lastIdxFromIndexFile, err := bl.readLastIndex(indexByte, false)
			if err != nil {
				return err
			}
			if lastIdxFromIndexFile != lastIdxFromLogFile {
				bl.reindex()
			}
		} else {
			bl.reindex()
		}
	} else if indexInfo.Size() != 0 {
		bl.indexFile.Truncate(0)
	}

	return
}

// Remove remove log and index file
func (bl *BLog) Remove(dir string) {
	bl.logFile.Close()
	bl.indexFile.Close()
	os.Remove(dir + "/block.bin")
	os.Remove(dir + "/block.index")
}

// Append appends a common.SignedBlock to the BLog
func (bl *BLog) Append(sb common.ISignedBlock) error {
	logFileOffset, _ := bl.logFile.Seek(0, 2)
	bl.indexFile.Seek(0, 2)
	// TODO: check index cnt and sb block num

	payload, err := sb.Marshall()
	if err != nil {
		return fmt.Errorf("BLOG Append: %s", err.Error())
	}
	if len(payload) > maxPayloadLen {
		return errors.New("BLog Append: common.SignedBlock too big")
	}

	// append payload len to log file
	lenByte := make([]byte, blockLenSize)
	binary.LittleEndian.PutUint32(lenByte, uint32(len(payload)))
	_, err = bl.logFile.Write(lenByte)
	if err != nil {
		return fmt.Errorf("BLOG Append: %s", err.Error())
	}

	// append payload to log file
	_, err = bl.logFile.Write(payload)
	if err != nil {
		return fmt.Errorf("BLOG Append: %s", err.Error())
	}

	// append index to log file
	indexByte := make([]byte, indexSize)
	binary.LittleEndian.PutUint64(indexByte, uint64(logFileOffset))
	_, err = bl.logFile.Write(indexByte)
	if err != nil {
		return err
	}

	// append index to index file
	_, err = bl.indexFile.Write(indexByte)
	if err != nil {
		return err
	}

	return nil
}

// Size...
func (bl *BLog) Size() int64 {
	idxInfo, err := bl.indexFile.Stat()
	if err != nil {
		panic(err)
	}

	return idxInfo.Size() / indexSize
}

// Empty returns true if it contains no block
func (bl *BLog) Empty() bool {
	logInfo, err := bl.logFile.Stat()
	if err != nil {
		panic(err)
	}

	return logInfo.Size() == 0
}

// ReadBlock reads a block at blockNum, blockNum start at 0
func (bl *BLog) ReadBlock(sb common.ISignedBlock, blockNum int64) error {
	indexOffset := blockNum * indexSize
	// read index
	indexByte := make([]byte, indexSize)
	_, err := bl.indexFile.ReadAt(indexByte, indexOffset)
	if err != nil {
		return fmt.Errorf("BLOG ReadBlock: %s", err.Error())
	}
	offset := binary.LittleEndian.Uint64(indexByte)
	return bl.readBlock(sb, int64(offset))
}

func (bl *BLog) readBlock(sb common.ISignedBlock, idx int64) error {
	// read payload len
	payloadLenByte := make([]byte, blockLenSize)
	var payloadLen uint32
	_, err := bl.logFile.ReadAt(payloadLenByte, idx)
	if err != nil {
		return fmt.Errorf("BLOG readBlock: %s", err.Error())
	}
	payloadLen = binary.LittleEndian.Uint32(payloadLenByte)

	// read payload
	payloadByte := make([]byte, payloadLen)
	_, err = bl.logFile.ReadAt(payloadByte, idx+blockLenSize)
	if err != nil {
		return err
	}

	err = sb.Unmarshall(payloadByte)
	if err != nil {
		return fmt.Errorf("BLOG readBlock: %s", err.Error())
	}

	return nil
}

func (bl *BLog) reindex() (err error) {
	if bl.indexFile != nil {
		// TODO: error log
		bl.indexFile.Truncate(0)
	} else {
		return nil
	}

	var offset, end int64
	indexByte := make([]byte, indexSize)

	end, err = bl.readLastIndex(indexByte, true)
	if err != nil {
		return err
	}

	for offset < end {
		// read payload len
		payloadLenByte := make([]byte, blockLenSize)
		var payloadLen uint32
		_, err = bl.logFile.Read(payloadLenByte)
		if err != nil {
			return err
		}
		payloadLen = binary.LittleEndian.Uint32(payloadLenByte)

		// read payload
		payloadByte := make([]byte, payloadLen)
		_, err = bl.logFile.Read(payloadByte)
		if err != nil {
			return err
		}

		// read index
		_, err = bl.logFile.Read(indexByte)
		if err != nil {
			return err
		}

		// append index to indexFile
		_, err = bl.indexFile.Write(indexByte)
		if err != nil {
			return err
		}

		offset = int64(binary.LittleEndian.Uint32(indexByte))
	}
	return nil
}

func (bl *BLog) readLastIndex(indexByte []byte, isLogFile bool) (int64, error) {
	var file *os.File
	if isLogFile {
		file = bl.logFile
	} else {
		file = bl.indexFile
	}
	file.Seek(-indexSize, 2)
	_, err := file.Read(indexByte)
	if err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(indexByte)), nil
}
