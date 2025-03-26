package main

import (
	"context"
	"fmt"
	"github.com/ACME-AI-Co/go/option"
	"os"
	"time"

	"github.com/ACME-AI-Co/go"
	"github.com/ACME-AI-Co/go/internal/param"
)

func main() {
	ctx := context.Background()
	client := acmeaisdk.NewClient(option.WithBearerToken("test"))

	file, err := os.Open("birds.csv")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	// Create file upload params
	params := acmeaisdk.FileFileNewParams{
		File: param.Field[string]{
			Value: file,
		},
	}

	// Upload the file
	fileResponse, err := client.Files.FileNew(context.Background(), params)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}

	var filesResponse *acmeaisdk.FileFileslistResponse

	for {
		// List files to check status
		filesResponse, err = client.Files.Fileslist(ctx, acmeaisdk.FileFileslistParams{})
		if err != nil {
			fmt.Printf("Error listing files: %v\n", err)
			return
		}

		var found bool
		for _, f := range filesResponse.Files {
			if f.FileID == fileResponse.FileID && f.Status == acmeaisdk.FileFileslistResponseFilesStatusCompleted {
				found = true
				break
			}
		}

		if found {
			break
		}

		// Wait before checking again
		time.Sleep(1 * time.Second)
	}

	// Search for "chickadee" in the file
	searchParams := acmeaisdk.FileFileSearchParams{}

	searchResults, err := client.Files.FileSearch(ctx, fileResponse.FileID, searchParams)
	if err != nil {
		fmt.Printf("Error searching file: %v\n", err)
		return
	}

	// Print the search results
	fmt.Printf("%+v\n", searchResults)
}
