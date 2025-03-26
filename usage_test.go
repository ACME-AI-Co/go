// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acmeaisdk_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/acme-ai-sdk-go"
	"github.com/stainless-sdks/acme-ai-sdk-go/internal/testutil"
	"github.com/stainless-sdks/acme-ai-sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Files.FileNew(context.TODO(), acmeaisdk.FileFileNewParams{
		File: acmeaisdk.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response.FileID)
}
