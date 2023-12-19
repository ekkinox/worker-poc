package main

import (
	"fmt"
	"time"
)

func printObservation(title string, pool *WorkerPool) {
	fmt.Printf("****\n%s:\n", title)

	for name, execution := range pool.Report() {
		fmt.Println("")
		fmt.Printf("%s: %s (%d events)\n", name, execution.Status(), len(execution.Events()))
		for _, event := range execution.Events() {
			fmt.Printf("- message: %s (at: %s)\n", event.Message(), event.Timestamp().Format(time.DateTime))
		}
	}

	fmt.Println("****\n\n")
}
