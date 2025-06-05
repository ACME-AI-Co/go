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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Files.Fileslist(context.TODO(), acmeaisdk.FileFileslistParams{
		Limit:  acmeaisdk.F(int64(20)),
		Offset: acmeaisdk.F(int64(20)),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, file := range page.Files {
		t.Logf("%+v\n", file.FileID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, file := range page.Files {
			t.Logf("%+v\n", file.FileID)
		}
	}
}
