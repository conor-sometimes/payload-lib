package git

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5" // Import the go-git library
)

func Pull(localPath string) {
	// Open the repository
	repo, err := git.PlainOpen(localPath)
	if err != nil {
		log.Fatalf("Failed to open the repository: %v", err)
	}

	// Get the working tree for the repository
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Failed to get the worktree: %v", err)
	}

	// Pull the latest changes from the origin
	err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Println("Already up-to-date.")
		} else {
			log.Fatalf("Failed to pull the latest changes: %v", err)
		}
	} else {
		fmt.Println("Successfully pulled the latest changes.")
	}
}
