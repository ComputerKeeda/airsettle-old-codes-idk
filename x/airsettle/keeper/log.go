package keeper

import (
	"fmt"
	"os"
)

func Log(text string) {
	// Specify the path to the log file.
	logFilePath := "air.log"

	// Open the log file in append mode.
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// String to append to the log file.
	logEntry := text + "\n"

	// Write the log entry to the log file.
	_, err = logFile.WriteString(logEntry)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
		return
	}

	fmt.Println("Log entry added successfully!")
}

func LogCreateChainid(chainid string) {
	logFilePath := "test/chainid.test.air"
	logFile, _ := os.OpenFile(logFilePath, os.O_WRONLY|os.O_TRUNC, 0644) //logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer logFile.Close()
	_, _ = logFile.WriteString(chainid)
}
