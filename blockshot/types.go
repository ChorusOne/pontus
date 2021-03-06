package blockshot

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/ChorusOne/pontus-internal/blockchain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"math/big"

	_ "github.com/lib/pq"
)

// Nice Guide : https://www.alexedwards.net/blog/using-postgresql-jsonb

//Transaction struct contains relevant (to Pontus) details of a SKALE transaction
type Transaction struct {
	Nonce               uint64          `json:"nonce"`
	GasLimit            uint64          `json:"gasLimit"`
	GasPrice            *big.Int        `json:"gasPrice"`
	GasUsed             uint64          `json:"gasUsed"`
	FeeCurrency         *common.Address `json:"feeCurrency"`
	GatewayFeeRecipient *common.Address `json:"gatewayFeeRecipient"`
	GatewayFee          *big.Int        `json:"gatewayFee"`
	To                  *common.Address `json:"to"`
	Val                 *big.Int        `json:"value"`

	Raw *types.Transaction `json:"raw"`
}

//TransactionFrom creates a Transaction object given a *types.Transaction object and gas details
func TransactionFrom(tx *types.Transaction, gasReceipt *blockchain.GasReceipt) *Transaction {
	transaction := &Transaction{
		Raw:      tx,
		Nonce:    tx.Nonce(),
		GasLimit: tx.Gas(),
		GasPrice: tx.GasPrice(),
		//FeeCurrency:         tx.FeeCurrency(),
		//GatewayFeeRecipient: tx.GatewayFeeRecipient(),
		//GatewayFee:          tx.GatewayFee(),
		To:  tx.To(),
		Val: tx.Value(),

		GasUsed: gasReceipt.GasUsed,
	}

	return transaction

}

//EventLog contains relevant (to Pontus) details of a SKALE event log
type EventLog struct {
	BlockLogIndex uint
	TxHash        common.Hash
	TxLogIndex    uint
	// This will be vLog.Topics[0].Hex()
	Topic   string
	Details *types.Log
}

//EventLogFrom generates an EventLog object given a *types.Log object
func EventLogFrom(vLog *types.Log) *EventLog {

	eventLog := &EventLog{
		BlockLogIndex: vLog.Index,
		TxLogIndex:    vLog.TxIndex,
		TxHash:        vLog.TxHash,
		Topic:         vLog.Topics[0].Hex(),
		Details:       vLog,
	}

	return eventLog

}

// Value returns the JSON-encoded representation of the Transaction struct.
func (tx Transaction) Value() (driver.Value, error) {
	return json.Marshal(tx)
}

// Scan decodes a JSON-encoded value into a Transaction struct
func (tx *Transaction) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &tx)
}

//Tag represents a Pontus tag on a SKALE transaction
type Tag struct {
	Name       string `json:"eventname"`
	prettyName string
	Source     string            `json:"source"`
	Parameters map[string]string `json:"parameters"`
}

//NewTag is a constructor for Tag
func NewTag(name string, prettyName string, source string) *Tag {
	return &Tag{
		Name:       name,
		prettyName: prettyName,
		Source:     source,
		Parameters: make(map[string]string)}

}

//EventLogs is a Pontus wrapper over types.Log to define Value() and Scan()
type EventLogs []*types.Log

// Value returns the JSON-encoded representation of EventLogs
func (eLogs EventLogs) Value() (driver.Value, error) {
	return json.Marshal(eLogs)
}

// Scan decodes a JSON-encoded value to Tags
func (eLogs *EventLogs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &eLogs)
}

//Tags is a type representation for a slice of Tag pointers
type Tags []*Tag

// Value returns the JSON-encoded representation of Tags
func (tags Tags) Value() (driver.Value, error) {
	return json.Marshal(tags)
}

// Scan decodes a JSON-encoded value to Tags
func (tags *Tags) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &tags)
}

// Value returns the JSON-encoded representation of the Tag struct.
func (tag Tag) Value() (driver.Value, error) {
	return json.Marshal(tag)
}

// Scan decodes a JSON-encoded value into a Tag struct
func (tag *Tag) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &tag)
}

//TransactionDetailed wraps relevant details of a transaction
//including its event logs + tags generated by Pontus for the transaction
type TransactionDetailed struct {
	BlockNumber string      `json:"blockNumber"`
	Timestamp   string      `json:"timestamp"`
	TxHash      string      `json:"hash"`
	Details     Transaction `json:"details"`
	From        string      `json:"from"`
	To          string      `json:"to"`
	ELogs       EventLogs   `json:"logs"`
	Tags        Tags        `json:"tags"`
}