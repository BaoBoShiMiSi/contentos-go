package cos

import (
	//"math/big"
	"os"
	"os/user"
	//"path/filepath"
	//"runtime"
	"time"

	"github.com/coschain/contentos-go/p2p/depend/common"
	"github.com/coschain/contentos-go/p2p/depend/common/hexutil"
	//"github.com/ethereum/go-ethereum/consensus/ethash"
	//"github.com/ethereum/go-ethereum/core"
	"github.com/coschain/contentos-go/p2p/cos/downloader"
)

// DefaultConfig contains default settings for use on the Contentos main net.
var DefaultConfig = Config{
	SyncMode: downloader.FastSync,
	/*
	Ethash: ethash.Config{
		CacheDir:       "ethash",
		CachesInMem:    2,
		CachesOnDisk:   3,
		DatasetsInMem:  1,
		DatasetsOnDisk: 2,
	},
	*/
	NetworkId:     1,
	LightPeers:    100,
	DatabaseCache: 768,
	TrieCache:     256,
	TrieTimeout:   60 * time.Minute,

	//TxPool: core.DefaultTxPoolConfig,
}

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		if user, err := user.Current(); err == nil {
			home = user.HomeDir
		}
	}
	/*
	if runtime.GOOS == "windows" {
		DefaultConfig.Ethash.DatasetDir = filepath.Join(home, "AppData", "Ethash")
	} else {
		DefaultConfig.Ethash.DatasetDir = filepath.Join(home, ".ethash")
	}
	*/
}

//go:generate gencodec -type Config -field-override configMarshaling -formats toml -out gen_config.go

type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Contentos main net block is used.
	//Genesis *core.Genesis `toml:",omitempty"`

	// Protocol options
	NetworkId uint64 // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode
	NoPruning bool

	// Light client options
	LightServ  int `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightPeers int `toml:",omitempty"` // Maximum number of LES client peers

	// Database options
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	TrieCache          int
	TrieTimeout        time.Duration

	// Mining-related options
	Etherbase    common.Address `toml:",omitempty"`
	MinerThreads int            `toml:",omitempty"`
	MinerNotify  []string       `toml:",omitempty"`
	ExtraData    []byte         `toml:",omitempty"`

	// Ethash options
	//Ethash ethash.Config

	// Transaction pool options
	//TxPool core.TxPoolConfig

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// Miscellaneous options
	DocRoot string `toml:"-"`
}

type configMarshaling struct {
	ExtraData hexutil.Bytes
}
