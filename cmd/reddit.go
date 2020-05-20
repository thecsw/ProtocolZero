package cmd

import (
	"fmt"
	"os"

	"github.com/thecsw/ProtocolZero/reddit"
)

func cleanReddit() {
	fmt.Println("\n\tReddit Protocol Zero")
	fmt.Println("\nPlease enter the fields below:")

	clientID := getPass("Client ID:")
	clientSecret := getPass("Client Secret:")
	username := getPass("Username:")
	password := getPass("Password:")

	// Login to reddit
	fmt.Print("\nLogging in to Reddit...")
	r, err := reddit.Login(clientID, clientSecret, username, password)
	if err != nil {
		fmt.Println("Failed to login to reddit!", err)
		os.Exit(1)
	}
	fmt.Println(" [DONE]")

	// Ask one more time
	fmt.Println("\nThis will start deleting all your posts and comments!")
	confirmation := getPass("Are you sure? Type in 'y' or 'Y' if yes:")
	if confirmation != "y" && confirmation != "Y" {
		fmt.Println("\nConfirmation not received. Bailing out.")
		os.Exit(0)
	}

	// Start the wipe out
	fmt.Println("WIPING OUT IN PROGRESS...")
	// Let it be buffer sized 16
	posts := make(chan string, 16)
	// Set higher capacity so it can grow like a heap
	results := make([]string, 0, 128)
	// Set off the handler to start deleting posts
	r.ClearPosts(posts)
	for post := range posts {
		// Break if we're done
		if post == "QUIT" {
			break
		}
		results = append(results, post)
		clear()
		show(results, "DELETED", 16)
	}
	close(posts)
	// Set off the handler to start deleting comments
	comments := make(chan string, 16)
	r.ClearComments(comments)
	for comment := range comments {
		// Break if we're done
		if comment == "QUIT" {
			break
		}
		results = append(results, comment)
		clear()
		show(results, "DELETED", 16)
	}
	fmt.Println("\n DONE! Your reddit account should be clean.")
}
