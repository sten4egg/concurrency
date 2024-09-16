package main

import (
	"bufio"
	"concurrency/db"
	"concurrency/db/storage"
	"fmt"
	"os"
	"strings"
)

func main() {

	if err := db.InitLogger(); err != nil {
		fmt.Println("Failed to initialize logger:", err)
		os.Exit(1)
	}
	defer db.SyncLogger()

	engine := storage.NewStorage()
	pipeline := db.NewPipeline(engine)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "EXIT" || input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		output, err := pipeline.Process(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(output)
		}
	}
}
