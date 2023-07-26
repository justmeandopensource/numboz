/*
Copyright © 2023 Venkat Nagappan
*/
package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/justmeandopensource/numboz/cmd"
)

func main() {

	// handle os interrupt via channel
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		fmt.Println("\nGoodbye!")
		os.Exit(2)
	}()

	cmd.Execute()
}
