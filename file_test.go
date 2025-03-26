// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acmeaisdk_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/ACME-AI-Co/go"
	"github.com/ACME-AI-Co/go/internal/testutil"
	"github.com/ACME-AI-Co/go/option"
)

func TestFileFileNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Files.FileNew(context.TODO(), acmeaisdk.FileFileNewParams{
		File:        acmeaisdk.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		Description: acmeaisdk.F("description"),
		ProcessingOptions: acmeaisdk.F(acmeaisdk.FileFileNewParamsProcessingOptions{
			Language: acmeaisdk.F("language"),
			Ocr:      acmeaisdk.F(true),
		}),
	})
	if err != nil {
		var apierr *acmeaisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileFileSearchWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Files.FileSearch(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		acmeaisdk.FileFileSearchParams{
			Query:           acmeaisdk.F("query"),
			ContextSize:     acmeaisdk.F(int64(0)),
			IncludeMetadata: acmeaisdk.F(true),
			MaxResults:      acmeaisdk.F(int64(1)),
		},
	)
	if err != nil {
		var apierr *acmeaisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileFileslistWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Files.Fileslist(context.TODO(), acmeaisdk.FileFileslistParams{
		Limit:     acmeaisdk.F(int64(1)),
		Offset:    acmeaisdk.F(int64(0)),
		SortBy:    acmeaisdk.F(acmeaisdk.FileFileslistParamsSortByUploadTime),
		SortOrder: acmeaisdk.F(acmeaisdk.FileFileslistParamsSortOrderAsc),
		Status:    acmeaisdk.F(acmeaisdk.FileFileslistParamsStatusPending),
	})
	if err != nil {
		var apierr *acmeaisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
