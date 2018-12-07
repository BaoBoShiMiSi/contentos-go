package commands

import (
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet/mock"
	"github.com/coschain/contentos-go/rpc/mock_grpcpb"
	"github.com/coschain/contentos-go/rpc/pb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplyWithoutBeneficiaries(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	myassert := assert.New(t)
	cmd := ReplyCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"initminer", "Lorem Ipsum", "10000000000000000"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		reply_op := op.GetOp7()
		myassert.Equal(reply_op.Content, "Lorem Ipsum")
		myassert.Equal(reply_op.ParentUuid, uint64(10000000000000000))
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestReplyWithBeneficiaries(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	myassert := assert.New(t)
	cmd := ReplyCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"initminer", "Lorem Ipsum", "10000000000000000",
		"-b", "Alice=5,Bob=5"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		reply_op := op.GetOp7()
		if reply_op.Beneficiaries[0].Name.Value == "Alice" {
			myassert.Equal(reply_op.Beneficiaries[1].Name.Value, "Bob")
			myassert.Equal(reply_op.Beneficiaries[1].Weight, uint32(5))
		} else {
			myassert.Equal(reply_op.Beneficiaries[1].Name.Value, "Alice")
			myassert.Equal(reply_op.Beneficiaries[1].Weight, uint32(5))
		}
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}
