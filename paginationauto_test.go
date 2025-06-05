// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acmeaisdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/ACME-AI-Co/go"
	"github.com/ACME-AI-Co/go/internal/testutil"
	"github.com/ACME-AI-Co/go/option"
)

func TestAutoPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := acmeaisdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	iter := client.Files.FileslistAutoPaging(context.TODO(), acmeaisdk.FileFileslistParams{
		Limit:  acmeaisdk.F(int64(20)),
		Offset: acmeaisdk.F(int64(20)),
	})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		file := iter.Current()
		t.Logf("%+v\n", file.FileID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
