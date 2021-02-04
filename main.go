package main

import (
	"fmt"
	"github.com/StanislavDimitrenco/bot-reminder/cmd"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// listen signal termination
	terminateC := make(chan os.Signal)
	stopC := make(chan struct{})
	signal.Notify(terminateC, os.Interrupt, syscall.SIGTERM)
	go terminationListening(terminateC, stopC)

	cmd.Run()

	<-stopC

}

func terminationListening(terminateC chan os.Signal, stopC chan<- struct{}) {
	<-terminateC

	fmt.Println("Bot was stopped.")

	// TODO: close db connections, sockets, net connections

	stopC <- struct{}{}
}
