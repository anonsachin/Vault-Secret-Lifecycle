package autorenew

import (
	"fmt"
	"os"
	"os/signal"

	vault "github.com/hashicorp/vault/api"
)

//Token handles the management of the token by renewing it 
// and keeping it valid and usable
func Token(tokenWatcher *vault.LifetimeWatcher) {
	fmt.Println("Strating auto renew of token")

	// Starting the go routines for managing renewals.
	go tokenWatcher.Start()
	defer tokenWatcher.Stop()

	// For handling ctrl+c.
	ctrlC := make(chan os.Signal)
	signal.Notify(ctrlC, os.Interrupt)

	for {
		select {
		// Renewal error channel.
		case err :=<- tokenWatcher.DoneCh():
			if err != nil {
				// Error when trying to renew.
				fmt.Printf("Unexpected error occured %#v \n",err)
				return
			}
			// Handles case when the token is no longer allowed to renew.
			fmt.Println("Failed to renew. Try re-login.")
			return
		// Renewal success channel.
		case renew := <- tokenWatcher.RenewCh():
			fmt.Printf("Successfully renewed: %#v \n", renew)
		// Ctrl + c handling.
		case close := <- ctrlC:
			fmt.Printf("We are closing the renewal %v \n",close)
			return
		}
	}
}
