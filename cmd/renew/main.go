package main

import (
	"fmt"
	"main/pkg/config"
	"os"
	vault "github.com/hashicorp/vault/api"
)


func main(){
	conf := config.GetConfig()

	client, err := vault.NewClient(conf)
	if err != nil {
		fmt.Printf("Unable to get a client %v",err)
		os.Exit(1)
	}

	secret, err := client.Logical().Read("secret/data/mysecret")

	if err != nil {
		fmt.Printf("Unable to read to secret %v\n",err)
		os.Exit(1)
	}

	if secret == nil {
		fmt.Println("The secret is empty")
		os.Exit(1)
	}

	fmt.Printf("The secret is %#v \n",secret.Data["data"].(map[string]interface{}))
	renew, err := secret.TokenIsRenewable()

	if err != nil {
		fmt.Printf("Unable to read to token status %v\n",err)
		os.Exit(1)
	}

	fmt.Printf("The token is %v \n",renew)

	data := map[string]interface{}{
		"data": map[string]interface{}{
			"Program": true,
		"Env": "Playground",
		},
	}

	secret, err = client.Logical().Write("secret/data/new-secret",data)

	if err != nil {
		fmt.Printf("Unable to write to secret %v\n",err)
		os.Exit(1)
	}
}