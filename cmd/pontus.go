package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ChorusOne/pontus-internal/blockshot"
	"github.com/ChorusOne/pontus-internal/config"
	"github.com/ChorusOne/pontus-internal/rest"

	"github.com/ChorusOne/pontus-internal/connection"
)

func main() {
	// Begin.
	log.Print("Pontus - Fortuna, favor us.")
	timestamp := time.Now().Unix()

	// Load environment variables
	config.LoadEnv()

	// Configure logging.
	// TODO: Consider paths, modularize
	var logDirFlag = flag.String("logDir", "../", "Path for Pontus log file.")
	file, err := os.OpenFile(fmt.Sprintf("%spontus_%d.log", *logDirFlag, timestamp), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file for Pontus.")
	}
	log.SetOutput(file)
	defer file.Close()

	// Create ethereum blockchain connection
	connection.InitEthClients()

	// Create database connection.
	connection.InitPostgresDB()

	// The following statement flushes the Pontus database.
	// Comment out if you wish to resume the iterator without flushing the database
	// TODO : Configify this
	blockshot.ResetDatabase()

	// Start non blocking goroutine for REST serving.
	go rest.StartAPI()

	blockshot.BlockIterator(
		[]blockshot.Task{
			//blockshot.ProcessTransactions,
			blockshot.TakeDailySnapshots,
		})
}
