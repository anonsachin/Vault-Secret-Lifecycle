package main

import (
	"fmt"
	"main/pkg/autorenew"
	"main/pkg/config"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	conf := config.GetConfig()

	client, err := vault.NewClient(conf)
	if err != nil {
		fmt.Printf("Unable to get a client %v", err)
		os.Exit(1)
	}

	renew := client.Token()

	fmt.Printf("The token is %#v \n", renew)

	auth, err := client.Logical().Write("auth/token/renew-self", map[string]interface{}{
		// "token": renew,
	})

	if err != nil {
		fmt.Printf("Unable to renew %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("The new token is %v \n", auth.Auth.ClientToken)

	self, err := client.Auth().Token().Lookup(auth.Auth.ClientToken)

	if err != nil {
		fmt.Printf("Unable to self lookup %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("The token data is %v \n", self.Data)

	data := map[string]interface{}{
		"data": map[string]interface{}{
			"Program": true,
			"Env":     "Playground",
		},
	}

	secret, err := client.Logical().Write("secret/data/new-secret", data)

	if err != nil {
		fmt.Printf("Unable to write to secret %v\n", err)
		os.Exit(1)
	}

	secret, err = client.Logical().Read("secret/data/new-secret")

	if err != nil {
		fmt.Printf("Unable to read to secret %v\n", err)
		os.Exit(1)
	}

	if secret == nil {
		fmt.Println("The secret is empty")
		os.Exit(1)
	}

	fmt.Printf("The secret is %#v \n", secret.Data["data"].(map[string]interface{}))

	if !auth.Auth.Renewable {
		fmt.Println("The token is not renewable error ")
	}

	watcher, err := client.NewLifetimeWatcher(&vault.LifetimeWatcherInput{
		Secret: auth,
	})

	autorenew.Token(watcher)
}
