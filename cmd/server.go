// // This file just for testing.
package main

import (
	"sync"

	app "github.com/Mahmoud-Emad/envserver/app"
)

func main() {
	server := app.NewServer("0.0.0.0", "8080")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := server.Serve(); err != nil {
		}
	}()
	wg.Wait()
}
