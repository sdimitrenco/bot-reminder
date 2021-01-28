package main

import (
	"fmt"
	"os"
)

// terminationListening listen signals channel for termination and cleanup application before closing
func terminationListening(terminateC chan os.Signal, stopC chan<- struct{}) {
	<-terminateC

	fmt.Println("Bot was stopped.")

	// TODO: close db connections, sockets, net connections

	stopC <- struct{}{}
}
