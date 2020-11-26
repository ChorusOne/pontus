package snapshot

import (
	"log"
	"math/big"

	"github.com/ChorusOne/pontus-internal/connection"
)

// DumpSystemSnapshotData writes a Pontus system snapshot.
func DumpSystemSnapshotData(ss SystemSnapshot, blockNumber *big.Int) error {

	insertStatement := "INSERT INTO system_snapshot_data VALUES ($1, $2)" //AddSystemColumnWork

	_, err := connection.DBCLIENT.Exec(insertStatement,
		blockNumber.String(),
		ss.SkaleTokenSupply.String())
	//AddSystemColumnWork
	//Order Matters

	if err != nil {
		return err
	}

	return nil

}

// FetchAllSystemSnapshots returns Pontus system data from completed snapshots.
func FetchAllSystemSnapshots() ([]*SystemSnapshotRow, error) {

	allSnapshots := make([]*SystemSnapshotRow, 0)

	fetchStatement := `SELECT system_snapshot_data.block_number, 
		snapshot_blocks.block_date, system_snapshot_data.skale_token_supply 
		FROM system_snapshot_data
		INNER JOIN snapshot_blocks
		ON system_snapshot_data.block_number = snapshot_blocks.block_number
		ORDER BY system_snapshot_data.block_number DESC`

	rows, err := connection.DBCLIENT.Query(fetchStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ssr SystemSnapshotRow
		err := rows.Scan(
			&ssr.BlockNumber,
			&ssr.SnapshotDate,
			&ssr.SkaleTokenSupply)
		//AddColumnWork
		//Order Matters
		if err != nil {
			return nil, err
		}
		allSnapshots = append(allSnapshots, &ssr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allSnapshots, nil

}

// DumpSnapshotData writes a completed snapshot for a block containing SKALE data.
func DumpSnapshotData(snapshot Snapshot, blockNumber *big.Int) {
	//AddColumnWork
	insertStatement := "INSERT INTO snapshot_data VALUES ($1, $2, $3, $4, $5, $6, $7)" //AddColumnWork

	for address := range snapshot {
		snapshotRow := snapshot[address]
		_, err := connection.DBCLIENT.Exec(insertStatement,
			blockNumber.String(),
			address,
			snapshotRow.SkaleTokenBalance.String(),
			snapshotRow.SkaleTokenLockedBalance.String(),
			snapshotRow.SkaleTokenDelegatedBalance.String(),
			snapshotRow.SkaleTokenSlashedBalance.String(),
			snapshotRow.SkaleTokenRewards.String())
		//AddColumnWork //Order Matters
		if err != nil {
			log.Printf("Could not write snapshot to database: %s\n", err)
		}

	}
}

// FetchSnapshotAddresses returns stored addresses involved in a SKALE transaction.
func FetchSnapshotAddresses(blockNumber *big.Int) map[string]bool {

	addresses := make(map[string]bool)
	fetchStatement := "SELECT address from snapshot_data where block_number < $1"

	rows, err := connection.DBCLIENT.Query(fetchStatement, blockNumber.String())
	if err != nil {
		log.Printf("Database query error: %s\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var address string
		err := rows.Scan(&address)
		if err != nil {
			log.Printf("Database scan error: %s\n", err)
		}
		addresses[address] = true
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Database iteration error: %s\n", err)
	}

	return addresses

}

// FetchAccountSnapshot returns historical SKALE snapshot data for a given account
func FetchAccountSnapshot(account string) (AccountSnapshot, error) {

	var accountSnapshot = make(AccountSnapshot)
	//AddColumnWork
	//Order Matters
	fetchStatement := `SELECT block_number, skale_token_balance, skale_token_locked_balance, 
	                    skale_token_delegated_balance, skale_token_slashed_balance
						FROM snapshot_data 
						WHERE address ILIKE $1`

	rows, err := connection.DBCLIENT.Query(fetchStatement, account)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var asr AccountSnapshotRow
		var blockNumber string
		err := rows.Scan(
			&blockNumber,
			&asr.SkaleTokenBalance,
			&asr.SkaleTokenLockedBalance,
			&asr.SkaleTokenDelegatedBalance,
			&asr.SkaleTokenSlashedBalance)
		//AddColumnWork
		//Order Matters
		if err != nil {
			return nil, err
		}
		accountSnapshot[blockNumber] = &asr
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return accountSnapshot, nil

}

// FetchSnapshotData returns SKALE data for a given Ethereum block.
func FetchSnapshotData(blockNumber *big.Int) (Snapshot, error) {

	var snapshot Snapshot
	snapshot = make(map[string]SnapshotRow)
	//AddColumnWork // Order Matters
	fetchStatement := "SELECT address, skale_token_balance, skale_token_locked_balance, skale_token_delegated_balance, skale_token_slashed_balance, skale_token_rewards FROM snapshot_data WHERE block_number = $1"

	rows, err := connection.DBCLIENT.Query(fetchStatement, blockNumber.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var spr SnapshotPostgresRow
		err := rows.Scan(
			&spr.Address,
			&spr.SkaleTokenBalance,
			&spr.SkaleTokenLockedBalance,
			&spr.SkaleTokenDelegatedBalance,
			&spr.SkaleTokenSlashedBalance,
			&spr.SkaleTokenRewards)
		//AddColumnWork
		//Order Matters
		if err != nil {
			return nil, err
		}
		AddRowToSnapshot(&snapshot, &spr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return snapshot, nil

}

// CreateSnapshotStatusRow writes Pontus system status data for blocks containing SKALE data.
func CreateSnapshotStatusRow(blockNumber *big.Int, blockTimestamp uint64) {

	_, err := connection.DBCLIENT.Exec("INSERT INTO snapshot_blocks VALUES($1, $2, $3, $4)",
		blockNumber.String(),
		blockTimestamp,
		true,
		UTCTime(blockTimestamp).Format(("02-01-2006")),
	)

	if err != nil {
		log.Fatalln(err)
	}

}

func fetchPreviousSnapshotBlockNumber() *big.Int {

	var blockNumber int64
	stmt := `SELECT block_number FROM snapshot_blocks 
			WHERE snapshot_done = TRUE 
			ORDER BY block_number DESC LIMIT 1`

	err := connection.DBCLIENT.QueryRow(stmt).Scan(&blockNumber)

	if err != nil {
		log.Fatalln(err)
	}

	return big.NewInt(blockNumber)
}

// FlushAllSnapshots performs a psql quick remove on the snapshot_data table.
func FlushAllSnapshots() {

	_, err := connection.DBCLIENT.Exec("TRUNCATE TABLE snapshot_data")
	if err != nil {
		log.Fatalln(err)
	}

}

// AddSnapshotColumn appends a column to the snapshot_data table.
func AddSnapshotColumn(columnName string) {

	stmt := "ALTER TABLE snapshot_data ADD " + columnName + " text NOT NULL"
	_, err := connection.DBCLIENT.Exec(stmt)
	if err != nil {
		log.Fatalln(err)
	}

}

// DatesOfCompletedSnapshots returns a slice of data representing completed blocks.
func DatesOfCompletedSnapshots() ([]*SnapshotDate, error) {

	fetchStatement := `SELECT block_number, block_date 
					 	FROM snapshot_blocks 
					 	WHERE snapshot_done = TRUE
						ORDER BY block_number ASC`

	rows, err := connection.DBCLIENT.Query(fetchStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dates := make([]*SnapshotDate, 0, 100)

	for rows.Next() {
		var snapshotdate SnapshotDate
		err := rows.Scan(&snapshotdate.BlockNumber, &snapshotdate.Date)

		if err != nil {
			return nil, err
		}

		dates = append(dates, &snapshotdate)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return dates, nil

}
