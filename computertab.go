// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package computer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/computer-go/internal/apijson"
	"github.com/stainless-sdks/computer-go/internal/requestconfig"
	"github.com/stainless-sdks/computer-go/option"
	"github.com/stainless-sdks/computer-go/packages/param"
	"github.com/stainless-sdks/computer-go/packages/respjson"
)

// ComputerTabService contains methods and other services that help with
// interacting with the computer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputerTabService] method instead.
type ComputerTabService struct {
	Options []option.RequestOption
}

// NewComputerTabService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputerTabService(opts ...option.RequestOption) (r ComputerTabService) {
	r = ComputerTabService{}
	r.Options = opts
	return
}

// Create a new tab, optionally navigating to a URL (browser sessions only)
func (r *ComputerTabService) New(ctx context.Context, id string, body ComputerTabNewParams, opts ...option.RequestOption) (res *ComputerTabNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/tabs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of all open tabs with their IDs, URLs, titles, and main tab status
// (browser sessions only)
func (r *ComputerTabService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *ComputerTabListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/tabs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Close a specific tab by ID. Cannot close the last remaining tab (browser
// sessions only)
func (r *ComputerTabService) Delete(ctx context.Context, tabID string, body ComputerTabDeleteParams, opts ...option.RequestOption) (res *ComputerTabDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if tabID == "" {
		err = errors.New("missing required tab_id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/tabs/%s", body.ID, tabID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Switch the main/active tab to a different tab by ID (browser sessions only)
func (r *ComputerTabService) Switch(ctx context.Context, tabID string, body ComputerTabSwitchParams, opts ...option.RequestOption) (res *ComputerTabSwitchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if tabID == "" {
		err = errors.New("missing required tab_id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/tabs/%s/switch", body.ID, tabID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ComputerTabNewResponse struct {
	ErrorMessage  string                            `json:"error_message"`
	ExecutedTabID string                            `json:"executed_tab_id"`
	PageContext   ComputerTabNewResponsePageContext `json:"page_context"`
	RequestID     string                            `json:"request_id"`
	Result        map[string]any                    `json:"result"`
	Status        string                            `json:"status"`
	Timestamp     string                            `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorMessage  respjson.Field
		ExecutedTabID respjson.Field
		PageContext   respjson.Field
		RequestID     respjson.Field
		Result        respjson.Field
		Status        respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerTabNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerTabNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabNewResponsePageContext struct {
	DeviceScaleFactor float64 `json:"device_scale_factor"`
	IsMainTab         bool    `json:"is_main_tab"`
	PageHeight        int64   `json:"page_height"`
	PageWidth         int64   `json:"page_width"`
	ScrollX           float64 `json:"scroll_x"`
	ScrollY           float64 `json:"scroll_y"`
	TabID             string  `json:"tab_id"`
	Title             string  `json:"title"`
	URL               string  `json:"url"`
	ViewportHeight    int64   `json:"viewport_height"`
	ViewportWidth     int64   `json:"viewport_width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceScaleFactor respjson.Field
		IsMainTab         respjson.Field
		PageHeight        respjson.Field
		PageWidth         respjson.Field
		ScrollX           respjson.Field
		ScrollY           respjson.Field
		TabID             respjson.Field
		Title             respjson.Field
		URL               respjson.Field
		ViewportHeight    respjson.Field
		ViewportWidth     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerTabNewResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerTabNewResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabListResponse struct {
	ErrorMessage  string                             `json:"error_message"`
	ExecutedTabID string                             `json:"executed_tab_id"`
	PageContext   ComputerTabListResponsePageContext `json:"page_context"`
	RequestID     string                             `json:"request_id"`
	Result        map[string]any                     `json:"result"`
	Status        string                             `json:"status"`
	Timestamp     string                             `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorMessage  respjson.Field
		ExecutedTabID respjson.Field
		PageContext   respjson.Field
		RequestID     respjson.Field
		Result        respjson.Field
		Status        respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerTabListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerTabListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabListResponsePageContext struct {
	DeviceScaleFactor float64 `json:"device_scale_factor"`
	IsMainTab         bool    `json:"is_main_tab"`
	PageHeight        int64   `json:"page_height"`
	PageWidth         int64   `json:"page_width"`
	ScrollX           float64 `json:"scroll_x"`
	ScrollY           float64 `json:"scroll_y"`
	TabID             string  `json:"tab_id"`
	Title             string  `json:"title"`
	URL               string  `json:"url"`
	ViewportHeight    int64   `json:"viewport_height"`
	ViewportWidth     int64   `json:"viewport_width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceScaleFactor respjson.Field
		IsMainTab         respjson.Field
		PageHeight        respjson.Field
		PageWidth         respjson.Field
		ScrollX           respjson.Field
		ScrollY           respjson.Field
		TabID             respjson.Field
		Title             respjson.Field
		URL               respjson.Field
		ViewportHeight    respjson.Field
		ViewportWidth     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerTabListResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerTabListResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabDeleteResponse struct {
	ErrorMessage  string                               `json:"error_message"`
	ExecutedTabID string                               `json:"executed_tab_id"`
	PageContext   ComputerTabDeleteResponsePageContext `json:"page_context"`
	RequestID     string                               `json:"request_id"`
	Result        map[string]any                       `json:"result"`
	Status        string                               `json:"status"`
	Timestamp     string                               `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorMessage  respjson.Field
		ExecutedTabID respjson.Field
		PageContext   respjson.Field
		RequestID     respjson.Field
		Result        respjson.Field
		Status        respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerTabDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerTabDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabDeleteResponsePageContext struct {
	DeviceScaleFactor float64 `json:"device_scale_factor"`
	IsMainTab         bool    `json:"is_main_tab"`
	PageHeight        int64   `json:"page_height"`
	PageWidth         int64   `json:"page_width"`
	ScrollX           float64 `json:"scroll_x"`
	ScrollY           float64 `json:"scroll_y"`
	TabID             string  `json:"tab_id"`
	Title             string  `json:"title"`
	URL               string  `json:"url"`
	ViewportHeight    int64   `json:"viewport_height"`
	ViewportWidth     int64   `json:"viewport_width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceScaleFactor respjson.Field
		IsMainTab         respjson.Field
		PageHeight        respjson.Field
		PageWidth         respjson.Field
		ScrollX           respjson.Field
		ScrollY           respjson.Field
		TabID             respjson.Field
		Title             respjson.Field
		URL               respjson.Field
		ViewportHeight    respjson.Field
		ViewportWidth     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerTabDeleteResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerTabDeleteResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabSwitchResponse struct {
	ErrorMessage  string                               `json:"error_message"`
	ExecutedTabID string                               `json:"executed_tab_id"`
	PageContext   ComputerTabSwitchResponsePageContext `json:"page_context"`
	RequestID     string                               `json:"request_id"`
	Result        map[string]any                       `json:"result"`
	Status        string                               `json:"status"`
	Timestamp     string                               `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorMessage  respjson.Field
		ExecutedTabID respjson.Field
		PageContext   respjson.Field
		RequestID     respjson.Field
		Result        respjson.Field
		Status        respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerTabSwitchResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerTabSwitchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabSwitchResponsePageContext struct {
	DeviceScaleFactor float64 `json:"device_scale_factor"`
	IsMainTab         bool    `json:"is_main_tab"`
	PageHeight        int64   `json:"page_height"`
	PageWidth         int64   `json:"page_width"`
	ScrollX           float64 `json:"scroll_x"`
	ScrollY           float64 `json:"scroll_y"`
	TabID             string  `json:"tab_id"`
	Title             string  `json:"title"`
	URL               string  `json:"url"`
	ViewportHeight    int64   `json:"viewport_height"`
	ViewportWidth     int64   `json:"viewport_width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceScaleFactor respjson.Field
		IsMainTab         respjson.Field
		PageHeight        respjson.Field
		PageWidth         respjson.Field
		ScrollX           respjson.Field
		ScrollY           respjson.Field
		TabID             respjson.Field
		Title             respjson.Field
		URL               respjson.Field
		ViewportHeight    respjson.Field
		ViewportWidth     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerTabSwitchResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerTabSwitchResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabNewParams struct {
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r ComputerTabNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerTabNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerTabNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTabDeleteParams struct {
	ID string `path:"id,required" json:"-"`
	paramObj
}

type ComputerTabSwitchParams struct {
	ID string `path:"id,required" json:"-"`
	paramObj
}
