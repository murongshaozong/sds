package serv

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/stratosnet/sds/relay/stratoschain/grpc"
	"github.com/stratosnet/sds/utils"
)

const (
	DefaultMsg = "Request Accepted"
)

type CmdResult struct {
	Msg string
}

type relayCmd struct {
}

func RelayAPI() *relayCmd {
	return &relayCmd{}
}

func (api *relayCmd) Sync(ctx context.Context, param []string) (CmdResult, error) {
	if len(param) != 1 || len(param[0]) == 0 {
		utils.ErrorLog("wrong number of arguments")
		return CmdResult{Msg: ""}, errors.New("wrong number of arguments")
	}
	txHash := param[0]
	err := grpc.QueryTxByHash(txHash)
	if err != nil {
		errMsg := fmt.Sprintf("error when calling grpc.QueryTxByHash for txHash[%v], reason: %v", txHash, err.Error())
		utils.DebugLogf(errMsg)
		return CmdResult{Msg: ""}, errors.New(errMsg)
	}
	return CmdResult{Msg: DefaultMsg}, nil
}
