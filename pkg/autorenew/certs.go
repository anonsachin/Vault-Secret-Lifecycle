package autorenew

import (
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func Certs(name string, certs *vault.Secret, interrupt chan os.Signal, client *vault.Client, certPath string, certsData map[string]interface{}) {
	fmt.Printf("Strating auto renew of secret %s\n", name)

	secretWatcher, err := client.NewLifetimeWatcher(&vault.LifetimeWatcherInput{
		Secret: certs,
	})

	if err != nil {
		fmt.Printf("Unable to create watcher %v\n", err)
		return
	}

	// Handling accidental Closes
	defer secretWatcher.Stop()

	for {
		// Starting the watcher
		go secretWatcher.Start()

		// Handling signals
		select {
		// When certificate expires
		case err := <-secretWatcher.DoneCh():
			if err != nil {
				// Error when trying to renew.
				fmt.Printf("Unexpected error occured %#v \n", err)
				return
			}
			// create new certs
			secretWatcher.Stop()

			// Creating new certs
			fmt.Println("Creating new certs")
			certs, err = client.Logical().Write(certPath, certsData)

			if err != nil {
				fmt.Printf("Unable to write to secret %v\n", err)
				return
			}

			if certs == nil {
				fmt.Println("The certs are empty.")
				return
			}

			fmt.Printf("The new certs serial number is %s \n", certs.Data["serial_number"])

			// New watcher for cewrts
			secretWatcher, err = client.NewLifetimeWatcher(&vault.LifetimeWatcherInput{
				Secret: certs,
			})

			if err != nil {
				fmt.Printf("Unable to create watcher %v\n", err)
				return
			}

			fmt.Println("New Watcher created")

		case close := <-interrupt:
			fmt.Printf("We are closing the renewal of %s %v \n", name, close)
			return
		}
	}
}
