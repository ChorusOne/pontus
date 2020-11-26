package snapshot

import (
	"log"
	"math/big"
	"time"

	"github.com/ChorusOne/pontus-internal/atlas"
	"github.com/ChorusOne/pontus-internal/blockchain"
	"github.com/ChorusOne/pontus-internal/constants"
	"github.com/ChorusOne/pontus-internal/contract"
)

const haltOnIntegrityCheckFailure = true

// GenerateAndStoreGenesisSnapshot writes details of the SKALE contract origin
func GenerateAndStoreGenesisSnapshot(blockNumber *big.Int) {
	log.Println("Attempting to snapshot genesis block")

	var snapshot Snapshot = GenesisSnapshotAccountLevel()
	DumpSnapshotData(snapshot, blockNumber)

	blockHeader := blockchain.GetHeader(blockNumber)
	log.Printf("Genesis block number: %s Block timestamp: %d\n", blockNumber, blockHeader.Time)

	CreateSnapshotStatusRow(blockNumber, blockHeader.Time)
	log.Println("Genesis block snapshot complete")

}

// GenerateAndStoreBlockSnapshots creates and writes database entries for SKALE data on Ethereum
// TODO split?
func GenerateAndStoreBlockSnapshots(blockNumber *big.Int) {
	log.Println("Attempting to snapshot block number: ", blockNumber)
	tStart := time.Now()
	previousSnapshotBlockNumber := fetchPreviousSnapshotBlockNumber()

	DelegationsAccountLevel(blockNumber, previousSnapshotBlockNumber)
	RewardsAccountLevel(blockNumber, previousSnapshotBlockNumber)

	var snapshot Snapshot = AccountLevel(blockNumber, previousSnapshotBlockNumber)
	var systemSnapshot SystemSnapshot = SystemLevel(blockNumber)

	integrity := CheckIntegrity(snapshot, systemSnapshot)

	if !integrity && haltOnIntegrityCheckFailure {
		log.Fatalln("Integrity Check Failed. Halting Pontus. ")

	}

	tEnd := time.Now()
	log.Println("SnapshotAccount Level Time : ", tEnd.Sub(tStart))

	DumpSnapshotData(snapshot, blockNumber)
	DumpSystemSnapshotData(systemSnapshot, blockNumber)

	blockHeader := blockchain.GetHeader(blockNumber)

	CreateSnapshotStatusRow(blockNumber, blockHeader.Time)
	log.Println("Snapshot Complete For Block #", blockNumber)

}

// GenesisSnapshotAccountLevel writes genesis block account details
func GenesisSnapshotAccountLevel() Snapshot {
	var snapshot Snapshot = make(Snapshot)

	stateAddressesWithData := atlas.GetBlockTransferAddresses(
		constants.GenesisBlockNumber[constants.NetActive],
		constants.GenesisBlockNumber[constants.NetActive])

	//Set All State Dump Columns From Storage
	for address := range stateAddressesWithData {
		row := NewSnapshotRow()
		row.SkaleTokenBalance = contract.AccountSkaleBalance(address,
			constants.GenesisBlockNumber[constants.NetActive])
		row.SkaleTokenDelegatedBalance = contract.AccountDelegatedSkaleBalance(address,
			constants.GenesisBlockNumber[constants.NetActive])
		row.SkaleTokenLockedBalance = contract.AccountLockedSkaleBalance(address,
			constants.GenesisBlockNumber[constants.NetActive])
		row.SkaleTokenSlashedBalance = contract.AccountSlashedSkaleBalance(address,
			constants.GenesisBlockNumber[constants.NetActive])
		row.SkaleTokenRewards = big.NewInt(0) //TODO: Fix when sorted with SKALE team
		//AddColumnWork
		snapshot[address] = row
	}

	return snapshot
}

// AccountLevel creates a snapshot row containing details for each SKALE account address
func AccountLevel(blockNumber *big.Int, previousSnapshotBlockNumber *big.Int) Snapshot {
	var snapshot = make(Snapshot)
	previousSnapshotAddresses := FetchSnapshotAddresses(previousSnapshotBlockNumber)
	presentStateAddresses := atlas.GetBlockTransferAddresses(previousSnapshotBlockNumber, blockNumber)

	for address := range AddressUnion(previousSnapshotAddresses, presentStateAddresses) {
		snapshot[address] = GenerateAccountRow(address, blockNumber)
	}
	return snapshot
}

// GenerateAccountRow adds relevant SKALE column data from the Ethereum blockchain.
func GenerateAccountRow(address string, blockNumber *big.Int) SnapshotRow {
	row := NewSnapshotRow()
	row.SkaleTokenBalance = contract.AccountSkaleBalance(address, blockNumber)
	row.SkaleTokenLockedBalance = contract.AccountLockedSkaleBalance(address, blockNumber)
	row.SkaleTokenDelegatedBalance = contract.AccountDelegatedSkaleBalance(address, blockNumber)
	row.SkaleTokenSlashedBalance = contract.AccountSlashedSkaleBalance(address, blockNumber)
	row.SkaleTokenRewards = big.NewInt(0) // TODO: replace once this is sorted with SKALE team
	return row
}

// SystemLevel returns the total supply of SKALE tokens a given block number.
func SystemLevel(blockNumber *big.Int) SystemSnapshot {
	var ss SystemSnapshot
	ss.SkaleTokenSupply = contract.SystemSkaleTokenSupply(blockNumber)
	//AddSystemColumnWork
	return ss
}

// DelegationsAccountLevel extracts delegation event info
func DelegationsAccountLevel(blockNumber *big.Int, previousSnapshotBlockNumber *big.Int) {
	//Get Delegation event info between previousSnapshotBlockNumber+1 and this block inclusive
	var fromBlock *big.Int = big.NewInt(0)
	fromBlock.Add(previousSnapshotBlockNumber, big.NewInt(1))
	atlas.GetDelegationEventsInfo(fromBlock, blockNumber)
	/* 	vid := big.NewInt(9)
	   	contract.GetValidatorDelegations(vid) */
}

// RewardsAccountLevel extracts delegation event info
func RewardsAccountLevel(blockNumber *big.Int, previousSnapshotBlockNumber *big.Int) {
	//Get Reward event info between previousSnapshotBlockNumber+1 and this block inclusive
	var fromBlock *big.Int = big.NewInt(0)
	fromBlock.Add(previousSnapshotBlockNumber, big.NewInt(1))
	atlas.GetRewardEventsInfo(fromBlock, blockNumber)
}
