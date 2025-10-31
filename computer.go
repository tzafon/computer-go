// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package computer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/computer-go/internal/apijson"
	shimjson "github.com/stainless-sdks/computer-go/internal/encoding/json"
	"github.com/stainless-sdks/computer-go/internal/requestconfig"
	"github.com/stainless-sdks/computer-go/option"
	"github.com/stainless-sdks/computer-go/packages/param"
	"github.com/stainless-sdks/computer-go/packages/respjson"
)

// ComputerService contains methods and other services that help with interacting
// with the computer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputerService] method instead.
type ComputerService struct {
	Options []option.RequestOption
}

// NewComputerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputerService(opts ...option.RequestOption) (r ComputerService) {
	r = ComputerService{}
	r.Options = opts
	return
}

// Create a new browser or desktop automation session with configurable timeout.
// Returns endpoints for executing actions, streaming events, and viewing
// screencast.
func (r *ComputerService) New(ctx context.Context, body ComputerNewParams, opts ...option.RequestOption) (res *ComputerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "computers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the current status and metadata of a computer instance
func (r *ComputerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ComputerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all active computers for the user's organization
func (r *ComputerService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ComputerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "computers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Execute a single action such as screenshot, click, type, navigate, scroll, debug
// or other computer use actions
func (r *ComputerService) ExecuteAction(ctx context.Context, id string, body ComputerExecuteActionParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/execute", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute a batch of actions in sequence, stopping on first error
func (r *ComputerService) ExecuteBatch(ctx context.Context, id string, body ComputerExecuteBatchParams, opts ...option.RequestOption) (res *ComputerExecuteBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/batch", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extend the timeout for a computer session and verify it is still running
func (r *ComputerService) KeepAlive(ctx context.Context, id string, opts ...option.RequestOption) (res *ComputerKeepAliveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/keepalive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Navigate the browser to a specified URL
func (r *ComputerService) Navigate(ctx context.Context, id string, body ComputerNavigateParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/navigate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stream real-time events using Server-Sent Events (SSE)
func (r *ComputerService) StreamEvents(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Terminate and clean up a computer instance, stopping the session and recording
// metrics
func (r *ComputerService) Terminate(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ActionResult struct {
	ErrorMessage string         `json:"error_message"`
	RequestID    string         `json:"request_id"`
	Result       map[string]any `json:"result"`
	Status       string         `json:"status"`
	Timestamp    string         `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorMessage respjson.Field
		RequestID    respjson.Field
		Result       respjson.Field
		Status       respjson.Field
		Timestamp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResult) RawJSON() string { return r.JSON.raw }
func (r *ActionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerResponse struct {
	ID        string            `json:"id"`
	CreatedAt string            `json:"created_at"`
	Endpoints map[string]string `json:"endpoints"`
	Status    string            `json:"status"`
	Type      string            `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Endpoints   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteBatchResponse map[string]any

type ComputerKeepAliveResponse map[string]any

type ComputerNewParams struct {
	ContextID param.Opt[string] `json:"context_id,omitzero"`
	// "browser"|"desktop"|"code" etc
	Kind           param.Opt[string] `json:"kind,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	// TODO: implement
	Display ComputerNewParamsDisplay `json:"display,omitzero"`
	// TODO: implement
	Stealth any `json:"stealth,omitzero"`
	paramObj
}

func (r ComputerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TODO: implement
type ComputerNewParamsDisplay struct {
	Height param.Opt[int64]   `json:"height,omitzero"`
	Scale  param.Opt[float64] `json:"scale,omitzero"`
	Width  param.Opt[int64]   `json:"width,omitzero"`
	paramObj
}

func (r ComputerNewParamsDisplay) MarshalJSON() (data []byte, err error) {
	type shadow ComputerNewParamsDisplay
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerNewParamsDisplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteActionParams struct {
	Body any
	paramObj
}

func (r ComputerExecuteActionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ComputerExecuteActionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type ComputerExecuteBatchParams struct {
	Body any
	paramObj
}

func (r ComputerExecuteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ComputerExecuteBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type ComputerNavigateParams struct {
	Body any
	paramObj
}

func (r ComputerNavigateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ComputerNavigateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
