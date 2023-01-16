package lavaprotocol

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/relayer/sigs"
	"github.com/lavanet/lava/utils"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

func VerifyRelayReply(reply *pairingtypes.RelayReply, relayRequest *pairingtypes.RelayRequest, addr string) error {
	serverKey, err := sigs.RecoverPubKeyFromRelayReply(reply, relayRequest)
	if err != nil {
		return err
	}
	serverAddr, err := sdk.AccAddressFromHex(serverKey.Address().String())
	if err != nil {
		return err
	}
	if serverAddr.String() != addr {
		return utils.LavaFormatError("reply server address mismatch ", ProviderFinzalizationDataError, &map[string]string{"parsed Address": serverAddr.String(), "expected address": addr})
	}

	return nil
}

func VerifyFinalizationData(reply *pairingtypes.RelayReply, relayRequest *pairingtypes.RelayRequest, addr string, latestSessionBlock int64, blockDistanceForfinalization uint32) (finalizedBlocks map[int64]string, errRet error) {
	strAdd, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, err
	}
	serverKey, err := sigs.RecoverPubKeyFromResponseFinalizationData(reply, relayRequest, strAdd)
	if err != nil {
		return nil, err
	}

	serverAddr, err := sdk.AccAddressFromHex(serverKey.Address().String())
	if err != nil {
		return nil, err
	}

	if serverAddr.String() != addr {
		return nil, utils.LavaFormatError("reply server address mismatch in finalization data ", ProviderFinzalizationDataError, &map[string]string{"parsed Address": serverAddr.String(), "expected address": addr})
	}

	finalizedBlocks = map[int64]string{} // TODO:: define struct in relay response
	err = json.Unmarshal(reply.FinalizedBlocksHashes, &finalizedBlocks)
	if err != nil {
		return nil, utils.LavaFormatError("failed in unmarshalling finalized blocks data", ProviderFinzalizationDataError, &map[string]string{"FinalizedBlocksHashes": string(reply.FinalizedBlocksHashes), "errMsg": err.Error()})
	}
	latestBlock := reply.LatestBlock
	err = verifyFinalizationDataIntegrity(latestBlock, latestSessionBlock, finalizedBlocks, blockDistanceForfinalization, addr)
	if err != nil {
		return nil, err
	}
	return
}

func verifyFinalizationDataIntegrity(latestBlock int64, latestSessionBlock int64, finalizedBlocks map[int64]string, blockDistanceForfinalization uint32, providerAddr string) error {
	sorted := make([]int64, len(finalizedBlocks))
	idx := 0
	maxBlockNum := int64(0)
	for blockNum := range finalizedBlocks {
		if !spectypes.IsFinalizedBlock(blockNum, latestBlock, blockDistanceForfinalization) {
			return utils.LavaFormatError("Simulation: provider returned non finalized block reply for reliability", ProviderFinzalizationDataAccountabilityError, &map[string]string{"blockNum": strconv.FormatInt(blockNum, 10), "latestBlock": strconv.FormatInt(latestBlock, 10), "Provider": providerAddr, "finalizedBlocks": fmt.Sprintf("%+v", finalizedBlocks)})
		}

		sorted[idx] = blockNum

		if blockNum > maxBlockNum {
			maxBlockNum = blockNum
		}
		idx++
		// TODO: check blockhash length and format
	}

	// check for consecutive blocks
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	for index := range sorted {
		if index != 0 && sorted[index]-1 != sorted[index-1] {
			// log.Println("provider returned non consecutive finalized blocks reply.\n Provider: %s", providerAcc)
			return utils.LavaFormatError("Simulation: provider returned non consecutive finalized blocks reply", ProviderFinzalizationDataAccountabilityError, &map[string]string{"curr block": strconv.FormatInt(sorted[index], 10), "prev block": strconv.FormatInt(sorted[index-1], 10), "Provider": providerAddr, "finalizedBlocks": fmt.Sprintf("%+v", finalizedBlocks)})
		}
	}

	// check that latest finalized block address + 1 points to a non finalized block
	if spectypes.IsFinalizedBlock(maxBlockNum+1, latestBlock, blockDistanceForfinalization) {
		return utils.LavaFormatError("Simulation: provider returned finalized hashes for an older latest block", ProviderFinzalizationDataAccountabilityError, &map[string]string{
			"maxBlockNum": strconv.FormatInt(maxBlockNum, 10),
			"latestBlock": strconv.FormatInt(latestBlock, 10), "Provider": providerAddr, "finalizedBlocks": fmt.Sprintf("%+v", finalizedBlocks),
		})
	}

	// New reply should have blocknum >= from block same provider
	if latestSessionBlock > latestBlock {
		return utils.LavaFormatError("Simulation: Provider supplied an older latest block than it has previously", ProviderFinzalizationDataAccountabilityError, &map[string]string{
			"session.LatestBlock": strconv.FormatInt(latestSessionBlock, 10),
			"latestBlock":         strconv.FormatInt(latestBlock, 10), "Provider": providerAddr,
		})
	}

	return nil
}
