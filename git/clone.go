package git

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5" // Import the go-git library
)

func Clone(repoURL, localPath string) {
	// Check if the directory already exists
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		// Clone the repository if it doesn't exist
		fmt.Println("Cloning repository...")

		_, err := git.PlainClone(localPath, false, &git.CloneOptions{
			URL:      repoURL,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatalf("Failed to clone the repository: %v", err)
		}

		fmt.Println("Repository cloned successfully.")
	} else {
		fmt.Printf("Directory %s already exists.\n", localPath)
	}
}
