//Package blockshot provides functions for iterating over the SKALE blockchain block-by-block
//to perform a list of extraction tasks over each block.
package blockshot

import (
	"log"
	"math/big"
	"time"

	"github.com/ChorusOne/pontus-internal/blockchain"
	"github.com/ChorusOne/pontus-internal/constants"
	"github.com/ChorusOne/pontus-internal/contract"
)

//TODO: Replace global state vars
const sleepSeconds = 6

var noBlocksIterated = big.NewInt(0).Sub(constants.GenesisBlockNumber[constants.NetActive], big.NewInt(1))

//Task is a type alias for extraction tasks
type Task = func(*big.Int) error

//BlockIterator iterates over the blocks and performs the supplied extraction tasks
func BlockIterator(tasks []Task) {
	contract.InitContracts()
	chainHeight := blockchain.GetLatestBlockNumber()

	for {
		nextBlockNumber := big.NewInt(0)
		nextBlockNumber.Add(LatestIteratedBlockNumber(), big.NewInt(1))

		// Sleep if (and until) chainHeight < nextBlockNumber
		for chainHeight.Cmp(nextBlockNumber) < 0 {
			log.Println("Chain Height :", chainHeight.String())
			log.Println(" | Next Block Number : ", nextBlockNumber.String())
			log.Printf("Chain height exceeded. Sleeping for %d seconds \n", sleepSeconds)
			time.Sleep(sleepSeconds * time.Second)
			chainHeight = blockchain.GetLatestBlockNumber()
		}

		// Iterate through all tasks for this block number
		for _, task := range tasks {
			err := task(nextBlockNumber)
			if err != nil {
				log.Fatal(err)
			}
		}
		SetIterationStatusComplete(nextBlockNumber)
	}
}
