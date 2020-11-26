package snapshot

import (
	"math/big"
)

type Address = string

type Snapshot map[Address]SnapshotRow

type SnapshotRow struct {
	SkaleTokenBalance          *big.Int
	SkaleTokenLockedBalance    *big.Int
	SkaleTokenDelegatedBalance *big.Int
	SkaleTokenSlashedBalance   *big.Int
	SkaleTokenRewards          *big.Int
	//AddColumnWork
}

// This type acts as a bridge between Postgres row and Golang SnapshotRow
type SnapshotPostgresRow struct {
	Address                    string
	SkaleTokenBalance          string
	SkaleTokenLockedBalance    string
	SkaleTokenDelegatedBalance string
	SkaleTokenSlashedBalance   string
	SkaleTokenRewards          string
	//AddColumnWork
}

func AddRowToSnapshot(snapshot *Snapshot, spr *SnapshotPostgresRow) {
	sr := NewSnapshotRow()
	//AddColumnWork
	sr.SkaleTokenBalance, _ = new(big.Int).SetString(spr.SkaleTokenBalance, 10)
	sr.SkaleTokenLockedBalance, _ = new(big.Int).SetString(spr.SkaleTokenLockedBalance, 10)
	sr.SkaleTokenDelegatedBalance, _ = new(big.Int).SetString(spr.SkaleTokenDelegatedBalance, 10)
	sr.SkaleTokenSlashedBalance, _ = new(big.Int).SetString(spr.SkaleTokenSlashedBalance, 10)
	sr.SkaleTokenRewards, _ = new(big.Int).SetString(spr.SkaleTokenRewards, 10)
	address := spr.Address
	(*snapshot)[address] = sr
}

//Use this to set trivial values
func NewSnapshotRow() SnapshotRow {
	var sr SnapshotRow
	sr.SkaleTokenBalance = big.NewInt(0)
	sr.SkaleTokenLockedBalance = big.NewInt(0)
	sr.SkaleTokenDelegatedBalance = big.NewInt(0)
	sr.SkaleTokenSlashedBalance = big.NewInt(0)
	sr.SkaleTokenRewards = big.NewInt(0)
	//AddColumnWork
	return sr
}

// ------ System Types ----

type SystemSnapshot struct {
	SkaleTokenSupply *big.Int
	//AddSystemColumnWork
}

// Types and Structs useful for API

// SnapshotInfo is a tuple representing a snapshot's date and block number
type SnapshotDate struct {
	Date        string
	BlockNumber string
}

type BlockNumber = string

// AccountSnapshot is helper type for fetching historical snapshot data of an account
type AccountSnapshot map[BlockNumber]*AccountSnapshotRow

// AccountSnapshotRow is helper type for fetching historical snapshot data of an account
type AccountSnapshotRow struct {
	//AddColumnWork
	SkaleTokenBalance          string
	SkaleTokenLockedBalance    string
	SkaleTokenDelegatedBalance string
	SkaleTokenSlashedBalance   string
}

// SystemSnapshotRow is a helper type for fetching historical snapshot of SKALE system data
type SystemSnapshotRow struct {
	BlockNumber      string
	SnapshotDate     string
	SkaleTokenSupply string
	//AddSystemColumnWork
}

type CommissionsSnapshot = []*CommissionsSnapshotRow

type CommissionsSnapshotRow struct {
	EpochBlockNumber *big.Int
	Validator        string
	ValidatorPayment *big.Int
	Group            string
	GroupPayment     *big.Int
}

type RelevantStateData struct {
	SkaleTokenBalance *big.Int
	//AddColumnWork
}
