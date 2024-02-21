package keeper

import (
	"encoding/binary"
	"sort"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/x/rewards/types"
)

// GetIprpcRewardsCurrent get the total number of IprpcReward
func (k Keeper) GetIprpcRewardsCurrent(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.IprpcRewardsCurrentPrefix)
	bz := store.Get(byteKey)

	// Current doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetIprpcRewardsCurrent set the total number of IprpcReward
func (k Keeper) SetIprpcRewardsCurrent(ctx sdk.Context, current uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.IprpcRewardsCurrentPrefix)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, current)
	store.Set(byteKey, bz)
}

// SetIprpcReward set a specific IprpcReward in the store
func (k Keeper) SetIprpcReward(ctx sdk.Context, iprpcReward types.IprpcReward) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IprpcRewardPrefix))
	b := k.cdc.MustMarshal(&iprpcReward)
	store.Set(GetIprpcRewardIDBytes(iprpcReward.Id), b)
}

// GetIprpcReward returns a IprpcReward from its id
func (k Keeper) GetIprpcReward(ctx sdk.Context, id uint64) (val types.IprpcReward, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IprpcRewardPrefix))
	b := store.Get(GetIprpcRewardIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveIprpcReward removes a IprpcReward from the store
func (k Keeper) RemoveIprpcReward(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IprpcRewardPrefix))
	store.Delete(GetIprpcRewardIDBytes(id))
}

// GetAllIprpcReward returns all IprpcReward
func (k Keeper) GetAllIprpcReward(ctx sdk.Context) (list []types.IprpcReward) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IprpcRewardPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.IprpcReward
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetIprpcRewardIDBytes returns the byte representation of the ID
func GetIprpcRewardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetIprpcRewardIDFromBytes returns ID in uint64 format from a byte array
func GetIprpcRewardIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

// PopIprpcReward gets the lowest id IprpcReward object and removes it
func (k Keeper) PopIprpcReward(ctx sdk.Context, advanceCurrent bool) (types.IprpcReward, bool) {
	current := k.GetIprpcRewardsCurrent(ctx)
	if advanceCurrent {
		k.SetIprpcRewardsCurrent(ctx, current+1)
	}
	return k.GetIprpcReward(ctx, current)
}

// AddSpecFunds adds funds for a specific spec for <duration> of months.
// This function is used by the fund-iprpc TX.
func (k Keeper) addSpecFunds(ctx sdk.Context, spec string, fund sdk.Coins, duration uint64) {
	startID := k.GetIprpcRewardsCurrent(ctx) + 1 // fund IPRPC only from the next month for <duration> months
	for i := startID; i < startID+duration; i++ {
		iprpcReward, found := k.GetIprpcReward(ctx, i)
		if found {
			// found IPRPC reward, find if spec exists
			specIndex := -1
			for i := 0; i < len(iprpcReward.SpecFunds); i++ {
				if iprpcReward.SpecFunds[i].Spec == spec {
					specIndex = i
				}
			}
			// update spec funds
			if specIndex >= 0 {
				iprpcReward.SpecFunds[specIndex].Fund = iprpcReward.SpecFunds[specIndex].Fund.Add(fund...)
			} else {
				iprpcReward.SpecFunds = append(iprpcReward.SpecFunds, types.Specfund{Spec: spec, Fund: fund})
			}
		} else {
			// did not find IPRPC reward -> create a new one
			iprpcReward.Id = i
			iprpcReward.SpecFunds = []types.Specfund{{Spec: spec, Fund: fund}}
		}
		k.SetIprpcReward(ctx, iprpcReward)
	}
}

// transferSpecFundsToNextMonth transfer the specFunds to the next month's IPRPC funds
// this function is used when there are no providers that should get the monthly IPRPC reward,
// so the reward transfers to the next month
func (k Keeper) transferSpecFundsToNextMonth(specFunds []types.Specfund, nextMonthSpecFunds []types.Specfund) []types.Specfund {
	mergedMap := make(map[string]sdk.Coins)

	// populate map with current spec funds
	for i, obj := range specFunds {
		mergedMap[obj.Spec] = specFunds[i].Fund
	}

	// update the merged map with the next month spec funds
	for i, obj := range nextMonthSpecFunds {
		if fund, ok := mergedMap[obj.Spec]; ok {
			mergedMap[obj.Spec] = fund.Add(nextMonthSpecFunds[i].Fund...)
		} else {
			mergedMap[obj.Spec] = obj.Fund
		}
	}

	// Convert map back to list and sort
	var mergedList []types.Specfund
	for spec, fund := range mergedMap {
		mergedList = append(mergedList, types.Specfund{Spec: spec, Fund: fund})
	}
	sort.Slice(mergedList, func(i, j int) bool { return mergedList[i].Spec < mergedList[j].Spec })

	return mergedList
}