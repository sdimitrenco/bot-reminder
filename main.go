package main

import (
	"context"
	"fmt"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/controlers"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/telegram"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var ctx context.Context

func main() {
	// listen signal termination
	terminateC := make(chan os.Signal)
	stopC := make(chan struct{})
	signal.Notify(terminateC, os.Interrupt, syscall.SIGTERM)
	go terminationListening(terminateC, stopC)

	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	//add database to context
	ctx = database.Boot(context.Background())

	//parser daily text
	go controlers.GetDailyText(ctx)

	//telegram server
	go telegram.Run(ctx)

	// Run listener
	<-stopC

}

func terminationListening(terminateC chan os.Signal, stopC chan<- struct{}) {
	<-terminateC

	fmt.Println("Bot was stopped.")

	// TODO: close db connections, sockets, net connections

	stopC <- struct{}{}
}
