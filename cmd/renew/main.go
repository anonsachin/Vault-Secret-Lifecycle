package main

import (
	"context"
	"fmt"
	"main/pkg/autorenew"
	"main/pkg/config"
	"os"
	"os/signal"
	"time"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	conf := config.GetConfig()

	client, err := vault.NewClient(conf)
	if err != nil {
		fmt.Printf("Unable to get a client %v", err)
		os.Exit(1)
	}

	//  Channel for interrupt handling
	ctrlC := make(chan os.Signal,2)
	signal.Notify(ctrlC, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	renew := client.Token()

	fmt.Printf("The token is %#v \n", renew)

	auth, err := client.Auth().Token().RenewSelf(0)

	if err != nil {
		fmt.Printf("Unable to renew %v\n", err)
		os.Exit(1)
	}

	watcher, err := client.NewLifetimeWatcher(&vault.LifetimeWatcherInput{
		Secret: auth,
	})

	if err != nil {
		fmt.Printf("Unable to create watcher for auth %v\n", err)
		os.Exit(1)
	}
	// Running token renewal on a different thread
	go autorenew.Token("Token", watcher, ctx)

	// Generating certs
	certsPath := "NewOrgCA/issue/client"
	certsData := map[string]interface{}{
		"ttl":         "60",
		"common_name": "vault.service.consul",
		"alt_names":   "localhost",
	}
	certs, err := client.Logical().Write(certsPath, certsData)

	if err != nil {
		fmt.Printf("Unable to write to secret %v\n", err)
		os.Exit(1)
	}

	if certs == nil {
		fmt.Println("The certs are empty.")
		os.Exit(1)
	}

	// Running the certs renewal
	go autorenew.Certs("Certificate", certs, ctx, client, certsPath, certsData)


	<- ctrlC
	fmt.Println("Defering cancel now")
	cancel()
	fmt.Println("Cancel called")
	// Just to allow the cancel output to be printed
	time.Sleep(1 *time.Millisecond)

}
