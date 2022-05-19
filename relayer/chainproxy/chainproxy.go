package chainproxy

import (
	"context"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/relayer/sentry"
	"github.com/lavanet/lava/relayer/sigs"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

type NodeMessage interface {
	GetServiceApi() *spectypes.ServiceApi
	Send(ctx context.Context) (*pairingtypes.RelayReply, error)
	RequestedBlock() int64
}

type ChainProxy interface {
	Start(context.Context) error
	GetSentry() *sentry.Sentry
	ParseMsg(string, []byte) (NodeMessage, error)
	PortalStart(context.Context, *btcec.PrivateKey, string)
}

func GetChainProxy(nodeUrl string, nConns uint, sentry *sentry.Sentry) (ChainProxy, error) {
	switch sentry.ApiInterface {
	case "jsonrpc":
		return NewJrpcChainProxy(nodeUrl, nConns, sentry), nil
	case "tendermintrpc":
		return NewtendermintRpcChainProxy(nodeUrl, nConns, sentry), nil
	case "rest":
		return NewRestChainProxy(nodeUrl, sentry), nil
	}
	return nil, fmt.Errorf("chain proxy for apiInterface (%s) not found", sentry.ApiInterface)
}

func SendRelay(
	ctx context.Context,
	cp ChainProxy,
	privKey *btcec.PrivateKey,
	url string,
	req string,
) (*pairingtypes.RelayReply, error) {

	//
	// Unmarshal request
	nodeMsg, err := cp.ParseMsg(url, []byte(req))
	if err != nil {
		return nil, err
	}

	//
	//
	reply, err := cp.GetSentry().SendRelay(ctx, func(clientSession *sentry.ClientSession) (*pairingtypes.RelayReply, error) {
		//client session is locked here
		err := CheckComputeUnits(clientSession, nodeMsg.GetServiceApi().ComputeUnits)
		if err != nil {
			return nil, err
		}

		relayRequest := &pairingtypes.RelayRequest{
			Provider:     clientSession.Client.Acc,
			ApiUrl:       url,
			Data:         []byte(req),
			SessionId:    uint64(clientSession.SessionId),
			ChainID:      cp.GetSentry().ChainID,
			CuSum:        clientSession.CuSum,
			BlockHeight:  cp.GetSentry().GetBlockHeight(),
			RelayNum:     clientSession.RelayNum,
			RequestBlock: nodeMsg.RequestedBlock(),
		}

		sig, err := sigs.SignRelay(privKey, []byte(relayRequest.String()))
		if err != nil {
			return nil, err
		}
		relayRequest.Sig = sig

		c := *clientSession.Client.Client
		reply, err := c.Relay(ctx, relayRequest)
		if err != nil {
			return nil, err
		}
		serverKey, err := sigs.RecoverPubKeyFromRelayReply(reply)
		if err != nil {
			return nil, err
		}
		serverAddr, err := sdk.AccAddressFromHex(serverKey.Address().String())
		if err != nil {
			return nil, err
		}
		if serverAddr.String() != clientSession.Client.Acc {
			return nil, fmt.Errorf("server address mismatch in reply (%s) (%s)", serverAddr.String(), clientSession.Client.Acc)
		}

		return reply, nil
	})

	return reply, err
}

func CheckComputeUnits(clientSession *sentry.ClientSession, apiCu uint64) error {
	clientSession.Client.SessionsLock.Lock()
	defer clientSession.Client.SessionsLock.Unlock()

	if clientSession.Client.UsedComputeUnits+apiCu > clientSession.Client.MaxComputeUnits {
		return fmt.Errorf("used all the available compute units")
	}

	clientSession.CuSum += apiCu
	clientSession.Client.UsedComputeUnits += apiCu
	clientSession.RelayNum += 1
	return nil
}
