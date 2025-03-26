// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acmeaisdk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/ACME-AI-Co/go/internal/apiform"
	"github.com/ACME-AI-Co/go/internal/apijson"
	"github.com/ACME-AI-Co/go/internal/apiquery"
	"github.com/ACME-AI-Co/go/internal/param"
	"github.com/ACME-AI-Co/go/internal/requestconfig"
	"github.com/ACME-AI-Co/go/option"
)

// FileService contains methods and other services that help with interacting with
// the acme-ai-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// Upload a file for processing with AI. The file will be analyzed and made
// searchable using natural language queries.
func (r *FileService) FileNew(ctx context.Context, body FileFileNewParams, opts ...option.RequestOption) (res *FileFileNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "files/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search for content within a processed file using natural language queries.
// Returns relevant passages and their context.
func (r *FileService) FileSearch(ctx context.Context, fileID string, query FileFileSearchParams, opts ...option.RequestOption) (res *FileFileSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("files/%s/search", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the processing status of files. Can be filtered by status and sorted by
// upload time.
func (r *FileService) Fileslist(ctx context.Context, query FileFileslistParams, opts ...option.RequestOption) (res *FileFileslistResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "files/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FileFileNewResponse struct {
	// Unique identifier for the file
	FileID string `json:"file_id" format:"uuid"`
	// Current processing status
	Status FileFileNewResponseStatus `json:"status"`
	// Time the file was uploaded
	UploadTime time.Time               `json:"upload_time" format:"date-time"`
	JSON       fileFileNewResponseJSON `json:"-"`
}

// fileFileNewResponseJSON contains the JSON metadata for the struct
// [FileFileNewResponse]
type fileFileNewResponseJSON struct {
	FileID      apijson.Field
	Status      apijson.Field
	UploadTime  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileFileNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileFileNewResponseJSON) RawJSON() string {
	return r.raw
}

// Current processing status
type FileFileNewResponseStatus string

const (
	FileFileNewResponseStatusPending    FileFileNewResponseStatus = "pending"
	FileFileNewResponseStatusProcessing FileFileNewResponseStatus = "processing"
)

func (r FileFileNewResponseStatus) IsKnown() bool {
	switch r {
	case FileFileNewResponseStatusPending, FileFileNewResponseStatusProcessing:
		return true
	}
	return false
}

type FileFileSearchResponse struct {
	// Unique identifier of the searched file
	FileID string `json:"file_id" format:"uuid"`
	// File metadata (only included if requested)
	Metadata FileFileSearchResponseMetadata `json:"metadata"`
	// The search query used
	Query   string                         `json:"query"`
	Results []FileFileSearchResponseResult `json:"results"`
	// Total number of results found
	TotalResults int64                      `json:"total_results"`
	JSON         fileFileSearchResponseJSON `json:"-"`
}

// fileFileSearchResponseJSON contains the JSON metadata for the struct
// [FileFileSearchResponse]
type fileFileSearchResponseJSON struct {
	FileID       apijson.Field
	Metadata     apijson.Field
	Query        apijson.Field
	Results      apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FileFileSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileFileSearchResponseJSON) RawJSON() string {
	return r.raw
}

// File metadata (only included if requested)
type FileFileSearchResponseMetadata struct {
	// User-provided description of the file
	Description string `json:"description"`
	// Unique identifier for the file
	FileID string `json:"file_id" format:"uuid"`
	// MIME type of the file
	FileType string `json:"file_type"`
	// Original name of the file
	Filename string `json:"filename"`
	// Number of pages (for documents)
	PageCount         int64                                           `json:"page_count"`
	ProcessingOptions FileFileSearchResponseMetadataProcessingOptions `json:"processing_options"`
	// Time the file was uploaded
	UploadTime time.Time `json:"upload_time" format:"date-time"`
	// Approximate word count
	WordCount int64                              `json:"word_count"`
	JSON      fileFileSearchResponseMetadataJSON `json:"-"`
}

// fileFileSearchResponseMetadataJSON contains the JSON metadata for the struct
// [FileFileSearchResponseMetadata]
type fileFileSearchResponseMetadataJSON struct {
	Description       apijson.Field
	FileID            apijson.Field
	FileType          apijson.Field
	Filename          apijson.Field
	PageCount         apijson.Field
	ProcessingOptions apijson.Field
	UploadTime        apijson.Field
	WordCount         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *FileFileSearchResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileFileSearchResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type FileFileSearchResponseMetadataProcessingOptions struct {
	// Language used for processing
	Language string `json:"language"`
	// Whether OCR was used
	Ocr  bool                                                `json:"ocr"`
	JSON fileFileSearchResponseMetadataProcessingOptionsJSON `json:"-"`
}

// fileFileSearchResponseMetadataProcessingOptionsJSON contains the JSON metadata
// for the struct [FileFileSearchResponseMetadataProcessingOptions]
type fileFileSearchResponseMetadataProcessingOptionsJSON struct {
	Language    apijson.Field
	Ocr         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileFileSearchResponseMetadataProcessingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileFileSearchResponseMetadataProcessingOptionsJSON) RawJSON() string {
	return r.raw
}

type FileFileSearchResponseResult struct {
	// Additional context information (document-type specific)
	AdditionalContext interface{} `json:"additional_context"`
	// Character ranges to highlight within the passage
	HighlightRanges []FileFileSearchResponseResultsHighlightRange `json:"highlight_ranges"`
	// Page number where the match was found (if applicable)
	PageNumber int64 `json:"page_number"`
	// Text passage containing the match with surrounding context
	Passage string `json:"passage"`
	// Relevance score of the result (0-1)
	RelevanceScore float64                          `json:"relevance_score"`
	JSON           fileFileSearchResponseResultJSON `json:"-"`
}

// fileFileSearchResponseResultJSON contains the JSON metadata for the struct
// [FileFileSearchResponseResult]
type fileFileSearchResponseResultJSON struct {
	AdditionalContext apijson.Field
	HighlightRanges   apijson.Field
	PageNumber        apijson.Field
	Passage           apijson.Field
	RelevanceScore    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *FileFileSearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileFileSearchResponseResultJSON) RawJSON() string {
	return r.raw
}

type FileFileSearchResponseResultsHighlightRange struct {
	// End index of highlight in passage
	End int64 `json:"end"`
	// Start index of highlight in passage
	Start int64                                           `json:"start"`
	JSON  fileFileSearchResponseResultsHighlightRangeJSON `json:"-"`
}

// fileFileSearchResponseResultsHighlightRangeJSON contains the JSON metadata for
// the struct [FileFileSearchResponseResultsHighlightRange]
type fileFileSearchResponseResultsHighlightRangeJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileFileSearchResponseResultsHighlightRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileFileSearchResponseResultsHighlightRangeJSON) RawJSON() string {
	return r.raw
}

type FileFileslistResponse struct {
	Files []FileFileslistResponseFile `json:"files"`
	// Maximum number of files returned
	Limit int64 `json:"limit"`
	// Number of files skipped
	Offset int64 `json:"offset"`
	// Total number of files matching the filter
	Total int64                     `json:"total"`
	JSON  fileFileslistResponseJSON `json:"-"`
}

// fileFileslistResponseJSON contains the JSON metadata for the struct
// [FileFileslistResponse]
type fileFileslistResponseJSON struct {
	Files       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileFileslistResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileFileslistResponseJSON) RawJSON() string {
	return r.raw
}

type FileFileslistResponseFile struct {
	// Time processing was completed (if applicable)
	CompletionTime time.Time `json:"completion_time" format:"date-time"`
	// Error message (if status is 'failed')
	Error string `json:"error"`
	// Unique identifier for the file
	FileID string `json:"file_id" format:"uuid"`
	// Size of the file in bytes
	FileSize int64 `json:"file_size"`
	// Original name of the file
	Filename string `json:"filename"`
	// Current processing status
	Status FileFileslistResponseFilesStatus `json:"status"`
	// Time the file was uploaded
	UploadTime time.Time                     `json:"upload_time" format:"date-time"`
	JSON       fileFileslistResponseFileJSON `json:"-"`
}

// fileFileslistResponseFileJSON contains the JSON metadata for the struct
// [FileFileslistResponseFile]
type fileFileslistResponseFileJSON struct {
	CompletionTime apijson.Field
	Error          apijson.Field
	FileID         apijson.Field
	FileSize       apijson.Field
	Filename       apijson.Field
	Status         apijson.Field
	UploadTime     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FileFileslistResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileFileslistResponseFileJSON) RawJSON() string {
	return r.raw
}

// Current processing status
type FileFileslistResponseFilesStatus string

const (
	FileFileslistResponseFilesStatusPending    FileFileslistResponseFilesStatus = "pending"
	FileFileslistResponseFilesStatusProcessing FileFileslistResponseFilesStatus = "processing"
	FileFileslistResponseFilesStatusCompleted  FileFileslistResponseFilesStatus = "completed"
	FileFileslistResponseFilesStatusFailed     FileFileslistResponseFilesStatus = "failed"
)

func (r FileFileslistResponseFilesStatus) IsKnown() bool {
	switch r {
	case FileFileslistResponseFilesStatusPending, FileFileslistResponseFilesStatusProcessing, FileFileslistResponseFilesStatusCompleted, FileFileslistResponseFilesStatusFailed:
		return true
	}
	return false
}

type FileFileNewParams struct {
	// The file to upload
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Optional description of the file
	Description       param.Field[string]                             `json:"description"`
	ProcessingOptions param.Field[FileFileNewParamsProcessingOptions] `json:"processing_options"`
}

func (r FileFileNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type FileFileNewParamsProcessingOptions struct {
	// Preferred language for processing
	Language param.Field[string] `json:"language"`
	// Enable OCR for image-based documents
	Ocr param.Field[bool] `json:"ocr"`
}

func (r FileFileNewParamsProcessingOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileFileSearchParams struct {
	// Natural language search query
	Query param.Field[string] `query:"query,required"`
	// Number of characters to include before and after the match
	ContextSize param.Field[int64] `query:"context_size"`
	// Whether to include file metadata in response
	IncludeMetadata param.Field[bool] `query:"include_metadata"`
	// Maximum number of results to return
	MaxResults param.Field[int64] `query:"max_results"`
}

// URLQuery serializes [FileFileSearchParams]'s query parameters as `url.Values`.
func (r FileFileSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileFileslistParams struct {
	// Maximum number of files to return
	Limit param.Field[int64] `query:"limit"`
	// Number of files to skip
	Offset param.Field[int64] `query:"offset"`
	// Field to sort by
	SortBy param.Field[FileFileslistParamsSortBy] `query:"sort_by"`
	// Sort order
	SortOrder param.Field[FileFileslistParamsSortOrder] `query:"sort_order"`
	// Filter by processing status
	Status param.Field[FileFileslistParamsStatus] `query:"status"`
}

// URLQuery serializes [FileFileslistParams]'s query parameters as `url.Values`.
func (r FileFileslistParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to sort by
type FileFileslistParamsSortBy string

const (
	FileFileslistParamsSortByUploadTime FileFileslistParamsSortBy = "upload_time"
	FileFileslistParamsSortByStatus     FileFileslistParamsSortBy = "status"
)

func (r FileFileslistParamsSortBy) IsKnown() bool {
	switch r {
	case FileFileslistParamsSortByUploadTime, FileFileslistParamsSortByStatus:
		return true
	}
	return false
}

// Sort order
type FileFileslistParamsSortOrder string

const (
	FileFileslistParamsSortOrderAsc  FileFileslistParamsSortOrder = "asc"
	FileFileslistParamsSortOrderDesc FileFileslistParamsSortOrder = "desc"
)

func (r FileFileslistParamsSortOrder) IsKnown() bool {
	switch r {
	case FileFileslistParamsSortOrderAsc, FileFileslistParamsSortOrderDesc:
		return true
	}
	return false
}

// Filter by processing status
type FileFileslistParamsStatus string

const (
	FileFileslistParamsStatusPending    FileFileslistParamsStatus = "pending"
	FileFileslistParamsStatusProcessing FileFileslistParamsStatus = "processing"
	FileFileslistParamsStatusCompleted  FileFileslistParamsStatus = "completed"
	FileFileslistParamsStatusFailed     FileFileslistParamsStatus = "failed"
)

func (r FileFileslistParamsStatus) IsKnown() bool {
	switch r {
	case FileFileslistParamsStatusPending, FileFileslistParamsStatusProcessing, FileFileslistParamsStatusCompleted, FileFileslistParamsStatusFailed:
		return true
	}
	return false
}
