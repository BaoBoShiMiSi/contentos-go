package rpc

/*
import (
	"bytes"
	"context"
	"fmt"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/common/logging"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/rpc/pb"
	"hash/crc32"
	"math"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

var asc grpcpb.ApiServiceClient

func TestMain(m *testing.M) {
	logging.Init("logs	", "debug", 0)

	os.RemoveAll("/Users/eagle/.coschain/cosd/db")

	addr := fmt.Sprintf("127.0.0.1:%d", uint32(8888))
	conn, err := Dial(addr)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer conn.Close()

	asc = grpcpb.NewApiServiceClient(conn)

	exitCode := m.Run()
	asc = nil

	os.Exit(exitCode)
}

func TestGRPCApi_GetAccountByName(t *testing.T) {
	req := &grpcpb.GetAccountByNameRequest{AccountName: &prototype.AccountName{Value: "Jack"}}
	resp := &grpcpb.AccountResponse{}
	resp, err := asc.GetAccountByName(context.Background(), req)

	if err != nil {
		t.Errorf("GetAccountByName failed: err:[%v], resp:[%x]", err, resp)
	} else {
		t.Logf("GetAccountByName detail: %v", resp.AccountName)
	}
}

func TestGPRCApi_GetFollowerListByName(t *testing.T) {
	req := &grpcpb.GetFollowerListByNameRequest{
		Limit: 100,
		Start: &prototype.FollowerCreatedOrder{
			Account:     &prototype.AccountName{Value: "Jack"},
			CreatedTime: &prototype.TimePointSec{UtcSeconds: uint32(time.Now().Second())},
			Follower:    &prototype.AccountName{Value: "Bob"},
		},
	}
	resp := &grpcpb.GetFollowerListByNameResponse{}
	resp, err := asc.GetFollowerListByName(context.Background(), req)
	if err != nil {
		t.Errorf("GetFollowerListByName failed: %v", err)
	} else {
		t.Logf("GetFollowerListByName detail: %v", resp.FollowerList)
	}
}

func TestGPRCApi_GetFollowingListByName(t *testing.T) {
	req := &grpcpb.GetFollowingListByNameRequest{
		Limit: 100,
		Start: &prototype.FollowingCreatedOrder{
			Account:     &prototype.AccountName{Value: "Jack"},
			CreatedTime: &prototype.TimePointSec{UtcSeconds: uint32(time.Now().Second())},
			Following:   &prototype.AccountName{Value: "Bob"},
		},
	}
	resp := &grpcpb.GetFollowingListByNameResponse{}
	resp, err := asc.GetFollowingListByName(context.Background(), req)
	if err != nil {
		t.Errorf("GetFollowingListByName failed: %v", err)
	} else {
		t.Logf("GetFollowingListByName detail: %v", resp.FollowingList)
	}
}

func TestGPRCApi_GetWitnessList(t *testing.T) {
	req := &grpcpb.GetWitnessListRequest{Limit: 100}
	resp := &grpcpb.GetWitnessListResponse{}
	resp, err := asc.GetWitnessList(context.Background(), req)
	if err != nil {
		t.Errorf("GetWitnessList failed: %v", err)
	} else {
		t.Logf("GetWitnessList detail: %v", resp.WitnessList)
	}
}

func TestGRPCApi_GetPostListByCreated(t *testing.T) {
	req := &grpcpb.GetPostListByCreatedRequest{}
	resp := &grpcpb.GetPostListByCreatedResponse{}

	resp, err := asc.GetPostListByCreated(context.Background(), req)
	if err != nil {
		t.Errorf("GetPostListByCreated failed: %v", err)
	} else {
		t.Logf("GetPostListByCreated detail: %v", resp.PostList)
	}
}

func TestGRPCApi_GetReplyListByPostId(t *testing.T) {
	req := &grpcpb.GetReplyListByPostIdRequest{}
	resp := &grpcpb.GetReplyListByPostIdResponse{}

	resp, err := asc.GetReplyListByPostId(context.Background(), req)
	if err != nil {
		t.Errorf("GetReplyListByPostId failed: %v", err)
	} else {
		t.Logf("GetReplyListByPostId detail: %v", resp.ReplyList)
	}
}

func TestGRPCApi_GetChainState(t *testing.T) {
	req := &grpcpb.NonParamsRequest{}
	resp := &grpcpb.GetChainStateResponse{}

	resp, err := asc.GetChainState(context.Background(), req)
	if err != nil {
		t.Errorf("GetChainState failed: %v", err)
	} else {
		t.Logf("GetChainState detail: %v", resp.Props)
	}
}

func TestGRPCApi_GetBlockTransactionsByNum(t *testing.T) {
	req := &grpcpb.GetBlockTransactionsByNumRequest{}
	resp := &grpcpb.GetBlockTransactionsByNumResponse{}

	resp, err := asc.GetBlockTransactionsByNum(context.Background(), req)
	if err != nil {
		t.Errorf("GetChainState failed: %v", err)
	} else {
		t.Logf("GetChainState detail: %v", resp.Transactions)
	}
}

func TestGRPCApi_GetTrxById(t *testing.T) {
	req := &grpcpb.GetTrxByIdRequest{}
	resp := &grpcpb.GetTrxByIdResponse{}

	resp, err := asc.GetTrxById(context.Background(), req)
	if err != nil {
		t.Errorf("GetTrxById failed: %v", err)
	} else {
		t.Logf("GetTrxById detail: %v", resp.Trx)
	}
}

var (
	BOB          = "BobName"
	ALICE        = "AliceName"
	pubkeyWIFBOB = "COS6Ezgyx3RQP5YjwBRf7higSytEVwELBCzK6xgB9orvpMuaLregA"
	prikeyWIFBOB = "YLC5nMjxPWvMPzDW9dC3d5UEamZwWffZpjWCmFq1Mk99EpQ1D"

	pubkeyWIFAlice = "COS65V8VdcvE4sF6qXtXs6k74TCi3rJrA5Lc5EqkH9Rh8YS3D2WT7"
	prikeyWIFAlice = "y9i4xUWGpbHQqfFjE1wL8LA2oevjhJtoej1KbMMJdoH9gnbhZ"
)

func TestGRPCApi_BroadcastTrx(t *testing.T) {
	//if test account is created in current db, pls comment out createAccount method
	pushTrx(t, createAccountTxReq(t))
	time.Sleep(time.Second * 5)

	pushTrx(t, createUnfollowTxReq(t))
	getFollowerList(t)
	pushTrx(t, createFollowTxReq(t))
	getFollowerList(t)

	uuid, postReq := createPostTxReq(t)
	pushTrx(t, postReq)
	time.Sleep(time.Second * 5)
	pushTrx(t, createRelayTxReq(t, uuid))

	getPostList(t)
	getRelyList(t, uuid)
}

func getPostList(t *testing.T) {
	req := &grpcpb.GetPostListByCreatedRequest{
		Start:&prototype.PostCreatedOrder{
			Created:&prototype.TimePointSec{UtcSeconds:math.MaxUint32},
			ParentId:0,
		},
		End:&prototype.PostCreatedOrder{
			Created:&prototype.TimePointSec{UtcSeconds: 0},
			ParentId:0,
		},
		Limit: 100,
	}
	resp := &grpcpb.GetPostListByCreatedResponse{}

	resp, err := asc.GetPostListByCreated(context.Background(), req)
	if err != nil {
		t.Errorf("GetPostListByCreated failed: %v", err)
	} else {
		t.Logf("GetPostListByCreated detail: %v", resp.PostList)
	}
}

func getRelyList(t *testing.T, parentId uint64) {
	req := &grpcpb.GetReplyListByPostIdRequest{
		Start:&prototype.ReplyCreatedOrder{
			ParentId:parentId,
			Created:&prototype.TimePointSec{UtcSeconds: math.MaxUint32},
		},
		End:&prototype.ReplyCreatedOrder{
			ParentId:parentId,
			Created:&prototype.TimePointSec{UtcSeconds: 0},
		},
		Limit: 100,
	}
	resp := &grpcpb.GetReplyListByPostIdResponse{}

	resp, err := asc.GetReplyListByPostId(context.Background(), req)
	if err != nil {
		t.Errorf("GetReplyListByPostId failed: %v", err)
	} else {
		t.Logf("GetReplyListByPostId detail: %v", resp.ReplyList)
	}
}

func createPostTxReq(t *testing.T) (uuid uint64, req *grpcpb.BroadcastTrxRequest) {
	title := "title_" + RandomString(15, "Aa0")
	uuid = GenerateUUID(BOB + title)
	post_op := &prototype.PostOperation{
		Uuid:          uuid,
		Owner:         &prototype.AccountName{Value: BOB},
		Title:         title,
		Content:       "content" + RandomString(100, "Aa0"),
		Tags:          []string{"abc"},
		Beneficiaries: []*prototype.BeneficiaryRouteType{},
	}
	return uuid, generateSignedTxResp(t, BOB, post_op)
}

func createRelayTxReq(t *testing.T, parentId uint64) (req *grpcpb.BroadcastTrxRequest) {
	content := "reply_content_" + RandomString(119, "Aa0")
	uuid := GenerateUUID(ALICE + content)
	reply_op := &prototype.ReplyOperation{
		Uuid:          uuid,
		Owner:         &prototype.AccountName{Value: ALICE},
		Content:       content,
		ParentUuid:    parentId,
		Beneficiaries: []*prototype.BeneficiaryRouteType{},
	}
	return generateSignedTxResp(t, ALICE, reply_op)
}

func getFollowerList(t *testing.T) {
	req := &grpcpb.GetFollowerListByNameRequest{
		Limit: 100,
	}
	resp := &grpcpb.GetFollowerListByNameResponse{}
	resp, err := asc.GetFollowerListByName(context.Background(), req)
	if err != nil {
		t.Errorf("GetFollowerListByName failed: %v", err)
	} else {
		t.Logf("GetFollowerListByName detail: %v", resp.FollowerList)
	}
}

func createFollowTxReq(t *testing.T) *grpcpb.BroadcastTrxRequest {
	fOP := &prototype.FollowOperation{
		Account:  &prototype.AccountName{Value: BOB},
		FAccount: &prototype.AccountName{Value: ALICE},
		Cancel:   false,
	}

	return generateSignedTxResp(t, BOB, fOP)
}

func createUnfollowTxReq(t *testing.T) *grpcpb.BroadcastTrxRequest {
	unfOP := &prototype.FollowOperation{
		Account:  &prototype.AccountName{Value: BOB},
		FAccount: &prototype.AccountName{Value: ALICE},
		Cancel:   true,
	}

	return generateSignedTxResp(t, BOB, unfOP)
}

func createAccountTxReq(t *testing.T) *grpcpb.BroadcastTrxRequest {

	pubkeyA, _ := prototype.PublicKeyFromWIF(pubkeyWIFBOB)
	pubkeyB, _ := prototype.PublicKeyFromWIF(pubkeyWIFAlice)

	keysA := prototype.NewAuthorityFromPubKey(pubkeyA)
	keysB := prototype.NewAuthorityFromPubKey(pubkeyB)

	acoA := &prototype.AccountCreateOperation{
		Fee:            prototype.NewCoin(1),
		Creator:        &prototype.AccountName{Value: constants.INIT_MINER_NAME},
		NewAccountName: &prototype.AccountName{Value: BOB},
		Owner:          keysA,
	}

	acoB := &prototype.AccountCreateOperation{
		Fee:            prototype.NewCoin(1),
		Creator:        &prototype.AccountName{Value: constants.INIT_MINER_NAME},
		NewAccountName: &prototype.AccountName{Value: ALICE},
		Owner:          keysB,
	}

	return generateSignedTxResp(t, constants.INIT_MINER_NAME, acoA, acoB)
}

func pushTrx(t *testing.T, req *grpcpb.BroadcastTrxRequest) {
	resp := &grpcpb.BroadcastTrxResponse{}

	resp, err := asc.BroadcastTrx(context.Background(), req)
	if err != nil || resp.Invoice.Status != 200 {
		t.Errorf("BroadcastTrx failed: err:[%v], status:[%d]", err, resp.Invoice.Status)
	} else {
		t.Logf("BroadcastTrx detail: resp: [%v]", resp)
	}
}

func generateSignedTxResp(t *testing.T, creator string, ops ...interface{}) *grpcpb.BroadcastTrxRequest {
	var creatorPrikey *prototype.PrivateKeyType
	switch creator {
	case constants.INIT_MINER_NAME:
		creatorPrikey, _ = prototype.PrivateKeyFromWIF(constants.INITMINER_PRIKEY)
	case BOB:
		creatorPrikey, _ = prototype.PrivateKeyFromWIF(prikeyWIFBOB)
	case ALICE:
		creatorPrikey, _ = prototype.PrivateKeyFromWIF(prikeyWIFAlice)
	default:
		creatorPrikey, _ = prototype.PrivateKeyFromWIF(constants.INITMINER_PRIKEY)
	}

	currTime := time.Now().Unix()

	tx := &prototype.Transaction{RefBlockNum: 0, RefBlockPrefix: 0, Expiration: &prototype.TimePointSec{UtcSeconds: uint32(currTime) + constants.TRX_MAX_EXPIRATION_TIME}}

	for _, op := range ops {
		tx.AddOperation(op)
	}

	signTx := prototype.SignedTransaction{Trx: tx}
	signTx.Signatures = append(signTx.Signatures, &prototype.SignatureType{Sig: signTx.Sign(creatorPrikey, prototype.ChainId{Value: 0})})

	if err := signTx.Validate(); err != nil {
		t.Error(err)
	}

	return &grpcpb.BroadcastTrxRequest{Transaction: &signTx}
}

func GenerateNewKey() (string, string, error) {
	privKey, err := prototype.GenerateNewKey()
	if err != nil {
		return "", "", err
	}
	pubKey, err := privKey.PubKey()
	if err != nil {
		return "", "", err
	}
	privKeyStr := privKey.ToWIF()
	pubKeyStr := pubKey.ToWIF()
	return pubKeyStr, privKeyStr, nil
}

func GenerateUUID(content string) uint64 {
	crc32q := crc32.MakeTable(0xD5828281)
	randContent := content + string(rand.Intn(1e5))
	return uint64(time.Now().Unix()*1e9) + uint64(crc32.Checksum([]byte(randContent), crc32q))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomString(randLength int, randType string) (result string) {

	//eg:
	//RandomString(8, "A")
	//RandomString(8, "a0")
	//RandomString(20, "Aa0")


	var num string = "0123456789"
	var lower string = "abcdefghijklmnopqrstuvwxyz"
	var upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := bytes.Buffer{}
	if strings.Contains(randType, "0") {
		b.WriteString(num)
	}
	if strings.Contains(randType, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(randType, "A") {
		b.WriteString(upper)
	}
	var str = b.String()
	var strLen = len(str)
	if strLen == 0 {
		result = ""
		return
	}

	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < randLength; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	result = b.String()
	return
}
*/

/*
func TestHTTPApi_GetAccountByName(t *testing.T) {
	postValue := "{\"account_name\": {\"value\":\"jack's test info\"}}"
	http_client("POST", "http://127.0.0.1:8080/v1/user/get_account_by_name", postValue)
}

func TestHTTPApi_GetWitnessList(t *testing.T) {
	http_client("GET", "http://127.0.0.1:8080/v1/user/get_witness_list?page=1&size=5", "")
}

func http_client(rtype, url, reqJson string) error {
	req, err := http.NewRequest(rtype, url, bytes.NewBuffer([]byte(reqJson)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//logging.CLog().Println("response Body:", string(body))

	return nil
}

func TestMockGRPCApi_GetAccountByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mock_grpcpb.NewMockIDatabaseService(ctrl)

	db.EXPECT().Has(gomock.Any()).Return(true , nil)
	expected := &grpcpb.AccountResponse{AccountName: &prototype.AccountName{Value: "Jack"}}
	db.EXPECT().Get(gomock.Any()).Return(expected, nil)

	cs := mock_grpcpb.NewMockIConsensus(ctrl)

	as := NewAPIService(cs, nil, db, nil)

	req := &grpcpb.GetAccountByNameRequest{AccountName:&prototype.AccountName{Value:"Jack"}}
	resp, err := as.GetAccountByName(context.Background(), req)

	if err != nil {
		t.Logf("GetAccountByName failed: %x", err)
	} else {
		t.Logf("GetAccountByName detail: %v", resp.AccountName)
	}

	//client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	//{
	//	req := &grpcpb.GetAccountByNameRequest{AccountName:&prototype.AccountName{Value:"Jack"}}
	//	resp := &grpcpb.AccountResponse{}
	//	expected := &grpcpb.AccountResponse{AccountName: &prototype.AccountName{Value: "Jack"}}
	//	client.EXPECT().GetAccountByName(gomock.Any(), gomock.Any()).Return(expected, nil)
	//
	//	resp, err := client.GetAccountByName(context.Background(), req)
	//	if err != nil {
	//		t.Logf("GetAccountByName failed: %x", err)
	//	} else {
	//		t.Logf("GetAccountByName detail: %v", resp.AccountName)
	//	}
	//}
}

func TestMockGPRCApi_GetFollowerListByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)

	{
		req := &grpcpb.GetFollowerListByNameRequest{}
		resp := &grpcpb.GetFollowerListByNameResponse{}

		expected := &grpcpb.GetFollowerListByNameResponse{}
		client.EXPECT().GetFollowerListByName(gomock.Any(), gomock.Any()).Return(expected, nil)

		resp, err := client.GetFollowerListByName(context.Background(), req)
		if err != nil {
			t.Logf("GetFollowerListByName failed: %x", err)
		} else {
			t.Logf("GetFollowerListByName detail: %v", resp.FollowerList)
		}
	}
}

func TestMockGRPCApi_GetFollowingListByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)

	{
		req := &grpcpb.GetFollowingListByNameRequest{}
		resp := &grpcpb.GetFollowingListByNameResponse{}

		expected := &grpcpb.GetFollowingListByNameResponse{}
		client.EXPECT().GetFollowingListByName(gomock.Any(), gomock.Any()).Return(expected, nil)

		resp, err := client.GetFollowingListByName(context.Background(), req)
		if err != nil {
			t.Logf("GetFollowingListByName failed: %x", err)
		} else {
			t.Logf("GetFollowingListByName detail: %v", resp.FollowingList)
		}
	}
}

func TestMockGPRCApi_GetWitnessList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)

	{
		req := &grpcpb.GetWitnessListRequest{}
		resp := &grpcpb.GetWitnessListResponse{}

		expected := &grpcpb.GetWitnessListResponse{}
		client.EXPECT().GetWitnessList(gomock.Any(), gomock.Any()).Return(expected, nil)

		resp, err := client.GetWitnessList(context.Background(), req)
		if err != nil {
			t.Logf("GetWitnessListByName failed: %x", err)
		} else {
			t.Logf("GetWitnessListByName detail: %v", resp.WitnessList)
		}
	}
}

*/

/*
Installation
First you need to install ProtocolBuffers 3.0.0 or later.

mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
Then, go get -u as usual the following packages:

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/coschain/contentos-go --go_out=plugins=grpc:. grpc.proto
protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/coschain/contentos-go --grpc-gateway_out=logtostderr=true:. grpc.proto

protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/coschain/contentos-go --go_out=plugins=grpc:. grpc.proto
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/coschain/contentos-go --grpc-gateway_out=logtostderr=true:. grpc.proto

compile:
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/coschain/contentos-go --go_out=plugins=grpc:. grpc.proto
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/coschain/contentos-go --grpc-gateway_out=logtostderr=true:. grpc.proto

mockgen -source=grpc.pb.go > ../mock_grpcpb/mock_grpcpb.go

protoc --go_out=paths=source_relative:. prototype/*.proto
*/
/*
service ApiService {
    rpc QueryTableContent (GetTableContentRequest) returns (TableContentResponse) {
        option (google.api.http) = {
            post: "/v1/vm/get_table_content"
            body: "*"
        };
    }

    rpc GetAccountByName (GetAccountByNameRequest) returns (AccountResponse) {
        option (google.api.http) = {
            post: "/v1/user/get_account_by_name"
            body: "*"
        };
    }

    rpc GetAccountRewardByName (GetAccountRewardByNameRequest) returns (AccountRewardResponse) {
        option (google.api.http) = {
            post: "/v1/user/get_account_reward_by_name"
            body: "*"
        };
    }

    rpc GetFollowerListByName (GetFollowerListByNameRequest) returns (GetFollowerListByNameResponse) {
        option (google.api.http) = {
            post: "/v1/user/get_follower_list_by_name"
            body: "*"
        };
    }

    rpc GetFollowingListByName (GetFollowingListByNameRequest) returns (GetFollowingListByNameResponse) {
        option (google.api.http) = {
            post: "/v1/user/get_following_list_by_name"
            body: "*"
        };
    }

    rpc GetFollowCountByName (GetFollowCountByNameRequest) returns (GetFollowCountByNameResponse) {
        option (google.api.http) = {
            post: "/v1/user/get_follow_count_by_name"
            body: "*"
        };
    }

    rpc GetWitnessList (GetWitnessListRequest) returns (GetWitnessListResponse) {
        option (google.api.http) = {
            get: "/v1/user/get_witness_list"
        };
    }

    rpc GetPostListByCreated (GetPostListByCreatedRequest) returns (GetPostListByCreatedResponse) {
        option (google.api.http) = {
            post: "/v1/post/get_post_list_by_created"
            body: "*"
        };
    }

    rpc GetReplyListByPostId (GetReplyListByPostIdRequest) returns (GetReplyListByPostIdResponse) {
        option (google.api.http) = {
            post: "/v1/post/get_reply_list_by_post_id"
            body: "*"
        };
    }

    rpc GetBlockTransactionsByNum (GetBlockTransactionsByNumRequest) returns (GetBlockTransactionsByNumResponse) {
        option (google.api.http) = {
            post: "/v1/trx/get_block_transactions_by_num"
            body: "*"
        };
    }

    rpc GetTrxById (GetTrxByIdRequest) returns (GetTrxByIdResponse) {
        option (google.api.http) = {
            post: "/v1/trx/get_trx_by_id"
            body: "*"
        };
    }

    rpc GetChainState (NonParamsRequest) returns (GetChainStateResponse) {
        option (google.api.http) = {
            post: "/v1/trx/get_chain_state"
            body: "*"
        };
    }


    rpc BroadcastTrx (BroadcastTrxRequest) returns (BroadcastTrxResponse) {
        option (google.api.http) = {
            post: "/v1/trx/broadcast_trx"
            body: "*"
        };
    }
}

*/