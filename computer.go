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

// ComputerService contains methods and other services that help with interacting
// with the computer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputerService] method instead.
type ComputerService struct {
	Options []option.RequestOption
	Tabs    ComputerTabService
}

// NewComputerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputerService(opts ...option.RequestOption) (r ComputerService) {
	r = ComputerService{}
	r.Options = opts
	r.Tabs = NewComputerTabService(opts...)
	return
}

// Create a new automation session. Set kind to "browser" for web automation or
// "desktop" for OS-level automation. Defaults to "browser" if not specified.
func (r *ComputerService) New(ctx context.Context, body ComputerNewParams, opts ...option.RequestOption) (res *ComputerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "computers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the current status and metadata of a computer instance
func (r *ComputerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ComputerGetResponse, err error) {
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
func (r *ComputerService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ComputerListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "computers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Take a screenshot of the current browser viewport, optionally as base64.
// Optionally specify tab_id (browser sessions only)
func (r *ComputerService) CaptureScreenshot(ctx context.Context, id string, body ComputerCaptureScreenshotParams, opts ...option.RequestOption) (res *ComputerCaptureScreenshotResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/screenshot", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a left mouse click at the specified x,y coordinates. Coordinates are
// screenshot pixel positions - send exactly what you see in the
// screenshot/screencast image. If target is at pixel (500, 300) in the image, send
// x=500, y=300. Optionally specify tab_id (browser sessions only)
func (r *ComputerService) Click(ctx context.Context, id string, body ComputerClickParams, opts ...option.RequestOption) (res *ComputerClickResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/click", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Establish WebSocket for real-time bidirectional communication
func (r *ComputerService) ConnectWebsocket(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/ws", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Execute a shell command with optional timeout and output length limits.
// Optionally specify tab_id (browser sessions only). Deprecated: use /exec or
// /exec/sync instead.
func (r *ComputerService) Debug(ctx context.Context, id string, body ComputerDebugParams, opts ...option.RequestOption) (res *ComputerDebugResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/debug", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a double mouse click at the specified x,y coordinates. Coordinates are
// screenshot pixel positions. Optionally specify tab_id (browser sessions only)
func (r *ComputerService) DoubleClick(ctx context.Context, id string, body ComputerDoubleClickParams, opts ...option.RequestOption) (res *ComputerDoubleClickResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/double-click", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a click-and-drag action from (x1,y1) to (x2,y2). Coordinates are
// screenshot pixel positions. Optionally specify tab_id (browser sessions only)
func (r *ComputerService) Drag(ctx context.Context, id string, body ComputerDragParams, opts ...option.RequestOption) (res *ComputerDragResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/drag", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute a single action such as screenshot, click, type, navigate, scroll,
// debug, set_viewport, get_html_content or other computer use actions
func (r *ComputerService) ExecuteAction(ctx context.Context, id string, body ComputerExecuteActionParams, opts ...option.RequestOption) (res *ComputerExecuteActionResponse, err error) {
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

// Get the HTML content of the current browser page. Optionally specify tab_id
// (browser sessions only)
func (r *ComputerService) GetHTML(ctx context.Context, id string, body ComputerGetHTMLParams, opts ...option.RequestOption) (res *ComputerGetHTMLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/html", id)
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

// Press and hold a keyboard key. Use with key_up for complex interactions.
// Optionally specify tab_id (browser sessions only)
func (r *ComputerService) KeyDown(ctx context.Context, id string, body ComputerKeyDownParams, opts ...option.RequestOption) (res *ComputerKeyDownResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/key-down", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Release a keyboard key that was previously pressed with key_down. Optionally
// specify tab_id (browser sessions only)
func (r *ComputerService) KeyUp(ctx context.Context, id string, body ComputerKeyUpParams, opts ...option.RequestOption) (res *ComputerKeyUpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/key-up", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Press and hold the left mouse button at the specified x,y coordinates.
// Coordinates are screenshot pixel positions. Optionally specify tab_id (browser
// sessions only)
func (r *ComputerService) MouseDown(ctx context.Context, id string, body ComputerMouseDownParams, opts ...option.RequestOption) (res *ComputerMouseDownResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/mouse-down", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Release the left mouse button at the specified x,y coordinates. Coordinates are
// screenshot pixel positions. Optionally specify tab_id (browser sessions only)
func (r *ComputerService) MouseUp(ctx context.Context, id string, body ComputerMouseUpParams, opts ...option.RequestOption) (res *ComputerMouseUpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/mouse-up", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Navigate the browser to a specified URL. Optionally specify tab_id to navigate a
// specific tab (browser sessions only)
func (r *ComputerService) Navigate(ctx context.Context, id string, body ComputerNavigateParams, opts ...option.RequestOption) (res *ComputerNavigateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/navigate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Press a combination of keys (e.g., ["Control", "c"] for copy). Optionally
// specify tab_id (browser sessions only)
func (r *ComputerService) PressHotkey(ctx context.Context, id string, body ComputerPressHotkeyParams, opts ...option.RequestOption) (res *ComputerPressHotkeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/hotkey", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a right mouse click at the specified x,y coordinates. Coordinates are
// screenshot pixel positions. Optionally specify tab_id (browser sessions only)
func (r *ComputerService) RightClick(ctx context.Context, id string, body ComputerRightClickParams, opts ...option.RequestOption) (res *ComputerRightClickResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/right-click", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scroll at the specified x,y position by delta dx,dy. Coordinates are screenshot
// pixel positions. Optionally specify tab_id (browser sessions only)
func (r *ComputerService) ScrollViewport(ctx context.Context, id string, body ComputerScrollViewportParams, opts ...option.RequestOption) (res *ComputerScrollViewportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/scroll", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change the browser viewport dimensions and scale factor. Optionally specify
// tab_id (browser sessions only)
func (r *ComputerService) SetViewport(ctx context.Context, id string, body ComputerSetViewportParams, opts ...option.RequestOption) (res *ComputerSetViewportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/viewport", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stream real-time events using Server-Sent Events (SSE)
func (r *ComputerService) StreamEvents(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Stream only screencast frames (base64 JPEG images) using Server-Sent Events
// (SSE) for live browser viewing
func (r *ComputerService) StreamScreencast(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/screencast", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Terminate and clean up a computer instance, stopping the session and recording
// metrics
func (r *ComputerService) Terminate(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Type text into the currently focused element in the browser. Optionally specify
// tab_id (browser sessions only)
func (r *ComputerService) TypeText(ctx context.Context, id string, body ComputerTypeTextParams, opts ...option.RequestOption) (res *ComputerTypeTextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/type", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ComputerNewResponse struct {
	ID        string            `json:"id"`
	CreatedAt string            `json:"created_at"`
	Endpoints map[string]string `json:"endpoints"`
	Kind      string            `json:"kind"`
	Status    string            `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Endpoints   respjson.Field
		Kind        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerGetResponse struct {
	ID        string            `json:"id"`
	CreatedAt string            `json:"created_at"`
	Endpoints map[string]string `json:"endpoints"`
	Kind      string            `json:"kind"`
	Status    string            `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Endpoints   respjson.Field
		Kind        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerListResponse struct {
	ID        string            `json:"id"`
	CreatedAt string            `json:"created_at"`
	Endpoints map[string]string `json:"endpoints"`
	Kind      string            `json:"kind"`
	Status    string            `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Endpoints   respjson.Field
		Kind        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerCaptureScreenshotResponse struct {
	ErrorMessage  string                                       `json:"error_message"`
	ExecutedTabID string                                       `json:"executed_tab_id"`
	PageContext   ComputerCaptureScreenshotResponsePageContext `json:"page_context"`
	RequestID     string                                       `json:"request_id"`
	Result        map[string]any                               `json:"result"`
	Status        string                                       `json:"status"`
	Timestamp     string                                       `json:"timestamp"`
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
func (r ComputerCaptureScreenshotResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerCaptureScreenshotResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerCaptureScreenshotResponsePageContext struct {
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
func (r ComputerCaptureScreenshotResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerCaptureScreenshotResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerClickResponse struct {
	ErrorMessage  string                           `json:"error_message"`
	ExecutedTabID string                           `json:"executed_tab_id"`
	PageContext   ComputerClickResponsePageContext `json:"page_context"`
	RequestID     string                           `json:"request_id"`
	Result        map[string]any                   `json:"result"`
	Status        string                           `json:"status"`
	Timestamp     string                           `json:"timestamp"`
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
func (r ComputerClickResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerClickResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerClickResponsePageContext struct {
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
func (r ComputerClickResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerClickResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDebugResponse struct {
	ErrorMessage  string                           `json:"error_message"`
	ExecutedTabID string                           `json:"executed_tab_id"`
	PageContext   ComputerDebugResponsePageContext `json:"page_context"`
	RequestID     string                           `json:"request_id"`
	Result        map[string]any                   `json:"result"`
	Status        string                           `json:"status"`
	Timestamp     string                           `json:"timestamp"`
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
func (r ComputerDebugResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerDebugResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDebugResponsePageContext struct {
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
func (r ComputerDebugResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerDebugResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDoubleClickResponse struct {
	ErrorMessage  string                                 `json:"error_message"`
	ExecutedTabID string                                 `json:"executed_tab_id"`
	PageContext   ComputerDoubleClickResponsePageContext `json:"page_context"`
	RequestID     string                                 `json:"request_id"`
	Result        map[string]any                         `json:"result"`
	Status        string                                 `json:"status"`
	Timestamp     string                                 `json:"timestamp"`
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
func (r ComputerDoubleClickResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerDoubleClickResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDoubleClickResponsePageContext struct {
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
func (r ComputerDoubleClickResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerDoubleClickResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDragResponse struct {
	ErrorMessage  string                          `json:"error_message"`
	ExecutedTabID string                          `json:"executed_tab_id"`
	PageContext   ComputerDragResponsePageContext `json:"page_context"`
	RequestID     string                          `json:"request_id"`
	Result        map[string]any                  `json:"result"`
	Status        string                          `json:"status"`
	Timestamp     string                          `json:"timestamp"`
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
func (r ComputerDragResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerDragResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDragResponsePageContext struct {
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
func (r ComputerDragResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerDragResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteActionResponse struct {
	ErrorMessage  string                                   `json:"error_message"`
	ExecutedTabID string                                   `json:"executed_tab_id"`
	PageContext   ComputerExecuteActionResponsePageContext `json:"page_context"`
	RequestID     string                                   `json:"request_id"`
	Result        map[string]any                           `json:"result"`
	Status        string                                   `json:"status"`
	Timestamp     string                                   `json:"timestamp"`
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
func (r ComputerExecuteActionResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerExecuteActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteActionResponsePageContext struct {
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
func (r ComputerExecuteActionResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerExecuteActionResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteBatchResponse map[string]any

type ComputerGetHTMLResponse struct {
	ErrorMessage  string                             `json:"error_message"`
	ExecutedTabID string                             `json:"executed_tab_id"`
	PageContext   ComputerGetHTMLResponsePageContext `json:"page_context"`
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
func (r ComputerGetHTMLResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerGetHTMLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerGetHTMLResponsePageContext struct {
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
func (r ComputerGetHTMLResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerGetHTMLResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerKeepAliveResponse map[string]any

type ComputerKeyDownResponse struct {
	ErrorMessage  string                             `json:"error_message"`
	ExecutedTabID string                             `json:"executed_tab_id"`
	PageContext   ComputerKeyDownResponsePageContext `json:"page_context"`
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
func (r ComputerKeyDownResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerKeyDownResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerKeyDownResponsePageContext struct {
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
func (r ComputerKeyDownResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerKeyDownResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerKeyUpResponse struct {
	ErrorMessage  string                           `json:"error_message"`
	ExecutedTabID string                           `json:"executed_tab_id"`
	PageContext   ComputerKeyUpResponsePageContext `json:"page_context"`
	RequestID     string                           `json:"request_id"`
	Result        map[string]any                   `json:"result"`
	Status        string                           `json:"status"`
	Timestamp     string                           `json:"timestamp"`
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
func (r ComputerKeyUpResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerKeyUpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerKeyUpResponsePageContext struct {
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
func (r ComputerKeyUpResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerKeyUpResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerMouseDownResponse struct {
	ErrorMessage  string                               `json:"error_message"`
	ExecutedTabID string                               `json:"executed_tab_id"`
	PageContext   ComputerMouseDownResponsePageContext `json:"page_context"`
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
func (r ComputerMouseDownResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerMouseDownResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerMouseDownResponsePageContext struct {
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
func (r ComputerMouseDownResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerMouseDownResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerMouseUpResponse struct {
	ErrorMessage  string                             `json:"error_message"`
	ExecutedTabID string                             `json:"executed_tab_id"`
	PageContext   ComputerMouseUpResponsePageContext `json:"page_context"`
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
func (r ComputerMouseUpResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerMouseUpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerMouseUpResponsePageContext struct {
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
func (r ComputerMouseUpResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerMouseUpResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerNavigateResponse struct {
	ErrorMessage  string                              `json:"error_message"`
	ExecutedTabID string                              `json:"executed_tab_id"`
	PageContext   ComputerNavigateResponsePageContext `json:"page_context"`
	RequestID     string                              `json:"request_id"`
	Result        map[string]any                      `json:"result"`
	Status        string                              `json:"status"`
	Timestamp     string                              `json:"timestamp"`
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
func (r ComputerNavigateResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerNavigateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerNavigateResponsePageContext struct {
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
func (r ComputerNavigateResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerNavigateResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerPressHotkeyResponse struct {
	ErrorMessage  string                                 `json:"error_message"`
	ExecutedTabID string                                 `json:"executed_tab_id"`
	PageContext   ComputerPressHotkeyResponsePageContext `json:"page_context"`
	RequestID     string                                 `json:"request_id"`
	Result        map[string]any                         `json:"result"`
	Status        string                                 `json:"status"`
	Timestamp     string                                 `json:"timestamp"`
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
func (r ComputerPressHotkeyResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerPressHotkeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerPressHotkeyResponsePageContext struct {
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
func (r ComputerPressHotkeyResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerPressHotkeyResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerRightClickResponse struct {
	ErrorMessage  string                                `json:"error_message"`
	ExecutedTabID string                                `json:"executed_tab_id"`
	PageContext   ComputerRightClickResponsePageContext `json:"page_context"`
	RequestID     string                                `json:"request_id"`
	Result        map[string]any                        `json:"result"`
	Status        string                                `json:"status"`
	Timestamp     string                                `json:"timestamp"`
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
func (r ComputerRightClickResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerRightClickResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerRightClickResponsePageContext struct {
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
func (r ComputerRightClickResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerRightClickResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerScrollViewportResponse struct {
	ErrorMessage  string                                    `json:"error_message"`
	ExecutedTabID string                                    `json:"executed_tab_id"`
	PageContext   ComputerScrollViewportResponsePageContext `json:"page_context"`
	RequestID     string                                    `json:"request_id"`
	Result        map[string]any                            `json:"result"`
	Status        string                                    `json:"status"`
	Timestamp     string                                    `json:"timestamp"`
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
func (r ComputerScrollViewportResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerScrollViewportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerScrollViewportResponsePageContext struct {
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
func (r ComputerScrollViewportResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerScrollViewportResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerSetViewportResponse struct {
	ErrorMessage  string                                 `json:"error_message"`
	ExecutedTabID string                                 `json:"executed_tab_id"`
	PageContext   ComputerSetViewportResponsePageContext `json:"page_context"`
	RequestID     string                                 `json:"request_id"`
	Result        map[string]any                         `json:"result"`
	Status        string                                 `json:"status"`
	Timestamp     string                                 `json:"timestamp"`
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
func (r ComputerSetViewportResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerSetViewportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerSetViewportResponsePageContext struct {
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
func (r ComputerSetViewportResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerSetViewportResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTypeTextResponse struct {
	ErrorMessage  string                              `json:"error_message"`
	ExecutedTabID string                              `json:"executed_tab_id"`
	PageContext   ComputerTypeTextResponsePageContext `json:"page_context"`
	RequestID     string                              `json:"request_id"`
	Result        map[string]any                      `json:"result"`
	Status        string                              `json:"status"`
	Timestamp     string                              `json:"timestamp"`
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
func (r ComputerTypeTextResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerTypeTextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTypeTextResponsePageContext struct {
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
func (r ComputerTypeTextResponsePageContext) RawJSON() string { return r.JSON.raw }
func (r *ComputerTypeTextResponsePageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerNewParams struct {
	// If true (default), kill session after inactivity
	AutoKill  param.Opt[bool]   `json:"auto_kill,omitzero"`
	ContextID param.Opt[string] `json:"context_id,omitzero"`
	// "browser" (default) or "desktop"
	Kind           param.Opt[string]        `json:"kind,omitzero"`
	TimeoutSeconds param.Opt[int64]         `json:"timeout_seconds,omitzero"`
	Display        ComputerNewParamsDisplay `json:"display,omitzero"`
	Stealth        any                      `json:"stealth,omitzero"`
	paramObj
}

func (r ComputerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type ComputerCaptureScreenshotParams struct {
	Base64 param.Opt[bool]   `json:"base64,omitzero"`
	TabID  param.Opt[string] `json:"tab_id,omitzero"`
	paramObj
}

func (r ComputerCaptureScreenshotParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerCaptureScreenshotParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerCaptureScreenshotParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerClickParams struct {
	TabID param.Opt[string]  `json:"tab_id,omitzero"`
	X     param.Opt[float64] `json:"x,omitzero"`
	Y     param.Opt[float64] `json:"y,omitzero"`
	paramObj
}

func (r ComputerClickParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerClickParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerClickParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDebugParams struct {
	Command         param.Opt[string] `json:"command,omitzero"`
	MaxOutputLength param.Opt[int64]  `json:"max_output_length,omitzero"`
	TabID           param.Opt[string] `json:"tab_id,omitzero"`
	TimeoutSeconds  param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	paramObj
}

func (r ComputerDebugParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerDebugParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerDebugParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDoubleClickParams struct {
	TabID param.Opt[string]  `json:"tab_id,omitzero"`
	X     param.Opt[float64] `json:"x,omitzero"`
	Y     param.Opt[float64] `json:"y,omitzero"`
	paramObj
}

func (r ComputerDoubleClickParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerDoubleClickParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerDoubleClickParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerDragParams struct {
	TabID param.Opt[string]  `json:"tab_id,omitzero"`
	X1    param.Opt[float64] `json:"x1,omitzero"`
	X2    param.Opt[float64] `json:"x2,omitzero"`
	Y1    param.Opt[float64] `json:"y1,omitzero"`
	Y2    param.Opt[float64] `json:"y2,omitzero"`
	paramObj
}

func (r ComputerDragParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerDragParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerDragParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteActionParams struct {
	Action ComputerExecuteActionParamsAction `json:"action,omitzero"`
	paramObj
}

func (r ComputerExecuteActionParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerExecuteActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerExecuteActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteActionParamsAction struct {
	// For get_html_content
	AutoDetectEncoding param.Opt[bool] `json:"auto_detect_encoding,omitzero"`
	// For screenshot
	Base64 param.Opt[bool]   `json:"base64,omitzero"`
	Button param.Opt[string] `json:"button,omitzero"`
	// For scrolling
	Dx     param.Opt[float64] `json:"dx,omitzero"`
	Dy     param.Opt[float64] `json:"dy,omitzero"`
	Height param.Opt[int64]   `json:"height,omitzero"`
	// Include page context in response
	IncludeContext param.Opt[bool] `json:"include_context,omitzero"`
	// For key_down/key_up
	Key      param.Opt[string] `json:"key,omitzero"`
	Ms       param.Opt[int64]  `json:"ms,omitzero"`
	ProxyURL param.Opt[string] `json:"proxy_url,omitzero"`
	// RequestId is used for correlating streaming output to the originating request.
	// Set on ActionRequest, not individual action types.
	RequestID   param.Opt[string]  `json:"request_id,omitzero"`
	ScaleFactor param.Opt[float64] `json:"scale_factor,omitzero"`
	// For tab management (browser sessions only)
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	Text  param.Opt[string] `json:"text,omitzero"`
	// click|double_click|right_click|drag|type|keypress|scroll|wait|screenshot|go_to_url|debug|get_html_content|set_viewport|list_tabs|new_tab|switch_tab|close_tab|key_down|key_up|mouse_down|mouse_up
	Type param.Opt[string] `json:"type,omitzero"`
	URL  param.Opt[string] `json:"url,omitzero"`
	// For set_viewport
	Width param.Opt[int64]   `json:"width,omitzero"`
	X     param.Opt[float64] `json:"x,omitzero"`
	// For dragging/scrolling
	X1 param.Opt[float64] `json:"x1,omitzero"`
	// For dragging
	X2    param.Opt[float64]                     `json:"x2,omitzero"`
	Y     param.Opt[float64]                     `json:"y,omitzero"`
	Y1    param.Opt[float64]                     `json:"y1,omitzero"`
	Y2    param.Opt[float64]                     `json:"y2,omitzero"`
	Debug ComputerExecuteActionParamsActionDebug `json:"debug,omitzero"`
	Keys  []string                               `json:"keys,omitzero"`
	paramObj
}

func (r ComputerExecuteActionParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow ComputerExecuteActionParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerExecuteActionParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteActionParamsActionDebug struct {
	Command         param.Opt[string] `json:"command,omitzero"`
	Cwd             param.Opt[string] `json:"cwd,omitzero"`
	MaxOutputLength param.Opt[int64]  `json:"max_output_length,omitzero"`
	RequestID       param.Opt[string] `json:"request_id,omitzero"`
	Stream          param.Opt[bool]   `json:"stream,omitzero"`
	TimeoutSeconds  param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	Env             map[string]string `json:"env,omitzero"`
	paramObj
}

func (r ComputerExecuteActionParamsActionDebug) MarshalJSON() (data []byte, err error) {
	type shadow ComputerExecuteActionParamsActionDebug
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerExecuteActionParamsActionDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteBatchParams struct {
	Actions []ComputerExecuteBatchParamsAction `json:"actions,omitzero"`
	paramObj
}

func (r ComputerExecuteBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerExecuteBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerExecuteBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteBatchParamsAction struct {
	// For get_html_content
	AutoDetectEncoding param.Opt[bool] `json:"auto_detect_encoding,omitzero"`
	// For screenshot
	Base64 param.Opt[bool]   `json:"base64,omitzero"`
	Button param.Opt[string] `json:"button,omitzero"`
	// For scrolling
	Dx     param.Opt[float64] `json:"dx,omitzero"`
	Dy     param.Opt[float64] `json:"dy,omitzero"`
	Height param.Opt[int64]   `json:"height,omitzero"`
	// Include page context in response
	IncludeContext param.Opt[bool] `json:"include_context,omitzero"`
	// For key_down/key_up
	Key      param.Opt[string] `json:"key,omitzero"`
	Ms       param.Opt[int64]  `json:"ms,omitzero"`
	ProxyURL param.Opt[string] `json:"proxy_url,omitzero"`
	// RequestId is used for correlating streaming output to the originating request.
	// Set on ActionRequest, not individual action types.
	RequestID   param.Opt[string]  `json:"request_id,omitzero"`
	ScaleFactor param.Opt[float64] `json:"scale_factor,omitzero"`
	// For tab management (browser sessions only)
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	Text  param.Opt[string] `json:"text,omitzero"`
	// click|double_click|right_click|drag|type|keypress|scroll|wait|screenshot|go_to_url|debug|get_html_content|set_viewport|list_tabs|new_tab|switch_tab|close_tab|key_down|key_up|mouse_down|mouse_up
	Type param.Opt[string] `json:"type,omitzero"`
	URL  param.Opt[string] `json:"url,omitzero"`
	// For set_viewport
	Width param.Opt[int64]   `json:"width,omitzero"`
	X     param.Opt[float64] `json:"x,omitzero"`
	// For dragging/scrolling
	X1 param.Opt[float64] `json:"x1,omitzero"`
	// For dragging
	X2    param.Opt[float64]                    `json:"x2,omitzero"`
	Y     param.Opt[float64]                    `json:"y,omitzero"`
	Y1    param.Opt[float64]                    `json:"y1,omitzero"`
	Y2    param.Opt[float64]                    `json:"y2,omitzero"`
	Debug ComputerExecuteBatchParamsActionDebug `json:"debug,omitzero"`
	Keys  []string                              `json:"keys,omitzero"`
	paramObj
}

func (r ComputerExecuteBatchParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow ComputerExecuteBatchParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerExecuteBatchParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecuteBatchParamsActionDebug struct {
	Command         param.Opt[string] `json:"command,omitzero"`
	Cwd             param.Opt[string] `json:"cwd,omitzero"`
	MaxOutputLength param.Opt[int64]  `json:"max_output_length,omitzero"`
	RequestID       param.Opt[string] `json:"request_id,omitzero"`
	Stream          param.Opt[bool]   `json:"stream,omitzero"`
	TimeoutSeconds  param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	Env             map[string]string `json:"env,omitzero"`
	paramObj
}

func (r ComputerExecuteBatchParamsActionDebug) MarshalJSON() (data []byte, err error) {
	type shadow ComputerExecuteBatchParamsActionDebug
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerExecuteBatchParamsActionDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerGetHTMLParams struct {
	AutoDetectEncoding param.Opt[bool]   `json:"auto_detect_encoding,omitzero"`
	TabID              param.Opt[string] `json:"tab_id,omitzero"`
	paramObj
}

func (r ComputerGetHTMLParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerGetHTMLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerGetHTMLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerKeyDownParams struct {
	Key   param.Opt[string] `json:"key,omitzero"`
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	paramObj
}

func (r ComputerKeyDownParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerKeyDownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerKeyDownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerKeyUpParams struct {
	Key   param.Opt[string] `json:"key,omitzero"`
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	paramObj
}

func (r ComputerKeyUpParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerKeyUpParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerKeyUpParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerMouseDownParams struct {
	TabID param.Opt[string]  `json:"tab_id,omitzero"`
	X     param.Opt[float64] `json:"x,omitzero"`
	Y     param.Opt[float64] `json:"y,omitzero"`
	paramObj
}

func (r ComputerMouseDownParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerMouseDownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerMouseDownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerMouseUpParams struct {
	TabID param.Opt[string]  `json:"tab_id,omitzero"`
	X     param.Opt[float64] `json:"x,omitzero"`
	Y     param.Opt[float64] `json:"y,omitzero"`
	paramObj
}

func (r ComputerMouseUpParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerMouseUpParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerMouseUpParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerNavigateParams struct {
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	URL   param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r ComputerNavigateParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerNavigateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerNavigateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerPressHotkeyParams struct {
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	Keys  []string          `json:"keys,omitzero"`
	paramObj
}

func (r ComputerPressHotkeyParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerPressHotkeyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerPressHotkeyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerRightClickParams struct {
	TabID param.Opt[string]  `json:"tab_id,omitzero"`
	X     param.Opt[float64] `json:"x,omitzero"`
	Y     param.Opt[float64] `json:"y,omitzero"`
	paramObj
}

func (r ComputerRightClickParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerRightClickParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerRightClickParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerScrollViewportParams struct {
	Dx    param.Opt[float64] `json:"dx,omitzero"`
	Dy    param.Opt[float64] `json:"dy,omitzero"`
	TabID param.Opt[string]  `json:"tab_id,omitzero"`
	X     param.Opt[float64] `json:"x,omitzero"`
	Y     param.Opt[float64] `json:"y,omitzero"`
	paramObj
}

func (r ComputerScrollViewportParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerScrollViewportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerScrollViewportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerSetViewportParams struct {
	Height      param.Opt[int64]   `json:"height,omitzero"`
	ScaleFactor param.Opt[float64] `json:"scale_factor,omitzero"`
	TabID       param.Opt[string]  `json:"tab_id,omitzero"`
	Width       param.Opt[int64]   `json:"width,omitzero"`
	paramObj
}

func (r ComputerSetViewportParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerSetViewportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerSetViewportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerTypeTextParams struct {
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	Text  param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r ComputerTypeTextParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerTypeTextParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerTypeTextParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
