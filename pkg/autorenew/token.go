package autorenew

import (
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
)

//Token handles the management of the tokens by renewing it
// and keeping it valid and usable
func Token(name string, secretWatcher *vault.LifetimeWatcher, interrupt chan os.Signal) {
	fmt.Printf("Strating auto renew of secret %s\n", name)

	// Starting the go routines for managing renewals.
	go secretWatcher.Start()
	defer secretWatcher.Stop()

	for {
		select {
		// Renewal error channel.
		case err := <-secretWatcher.DoneCh():
			if err != nil {
				// Error when trying to renew.
				fmt.Printf("Unexpected error occured %#v \n", err)
				return
			}
			// Handles case when the secret is no longer allowed to renew.
			fmt.Println("Failed to renew. Try re-login.")
			return
		// Renewal success channel.
		case renew := <-secretWatcher.RenewCh():
			fmt.Printf("Successfully renewed %s at: %s \n", name, renew.RenewedAt)
		// Ctrl + c handling.
		case close := <-interrupt:
			fmt.Printf("We are closing the renewal of %s %v \n", name, close)
			return
		}
	}
}
