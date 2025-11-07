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

// Take a screenshot of the current browser viewport, optionally as base64
func (r *ComputerService) CaptureScreenshot(ctx context.Context, id string, body ComputerCaptureScreenshotParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/screenshot", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a left mouse click at the specified x,y coordinates
func (r *ComputerService) Click(ctx context.Context, id string, body ComputerClickParams, opts ...option.RequestOption) (res *ActionResult, err error) {
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/ws", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Execute a shell command with optional timeout and output length limits
func (r *ComputerService) Debug(ctx context.Context, id string, body ComputerDebugParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/debug", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a double mouse click at the specified x,y coordinates
func (r *ComputerService) DoubleClick(ctx context.Context, id string, body ComputerDoubleClickParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/double-click", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a click-and-drag action from (x1,y1) to (x2,y2)
func (r *ComputerService) Drag(ctx context.Context, id string, body ComputerDragParams, opts ...option.RequestOption) (res *ActionResult, err error) {
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

// Get the HTML content of the current browser page
func (r *ComputerService) GetHTML(ctx context.Context, id string, body ComputerGetHTMLParams, opts ...option.RequestOption) (res *ActionResult, err error) {
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

// Press a combination of keys (e.g., ["Control", "c"] for copy)
func (r *ComputerService) PressHotkey(ctx context.Context, id string, body ComputerPressHotkeyParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/hotkey", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a right mouse click at the specified x,y coordinates
func (r *ComputerService) RightClick(ctx context.Context, id string, body ComputerRightClickParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/right-click", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scroll the browser viewport by the specified delta
func (r *ComputerService) ScrollViewport(ctx context.Context, id string, body ComputerScrollViewportParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/scroll", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change the browser viewport dimensions and scale factor
func (r *ComputerService) SetViewport(ctx context.Context, id string, body ComputerSetViewportParams, opts ...option.RequestOption) (res *ActionResult, err error) {
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Type text into the currently focused element in the browser
func (r *ComputerService) TypeText(ctx context.Context, id string, body ComputerTypeTextParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/type", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	// If true (default), kill session after inactivity. If false, only kill on
	// explicit stop or max_lifetime
	AutoKill  param.Opt[bool]   `json:"auto_kill,omitzero"`
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

type ComputerCaptureScreenshotParams struct {
	Base64 param.Opt[bool] `json:"base64,omitzero"`
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
	X param.Opt[float64] `json:"x,omitzero"`
	Y param.Opt[float64] `json:"y,omitzero"`
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
	X param.Opt[float64] `json:"x,omitzero"`
	Y param.Opt[float64] `json:"y,omitzero"`
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
	X1 param.Opt[float64] `json:"x1,omitzero"`
	X2 param.Opt[float64] `json:"x2,omitzero"`
	Y1 param.Opt[float64] `json:"y1,omitzero"`
	Y2 param.Opt[float64] `json:"y2,omitzero"`
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
	Dx          param.Opt[float64] `json:"dx,omitzero"`
	Dy          param.Opt[float64] `json:"dy,omitzero"`
	Height      param.Opt[int64]   `json:"height,omitzero"`
	Ms          param.Opt[int64]   `json:"ms,omitzero"`
	ScaleFactor param.Opt[float64] `json:"scale_factor,omitzero"`
	// For tab management (browser sessions only)
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	Text  param.Opt[string] `json:"text,omitzero"`
	// click|double_click|right_click|drag|type|keypress|scroll|wait|screenshot|go_to_url|debug|get_html_content|set_viewport|list_tabs|new_tab|switch_tab|close_tab
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
	MaxOutputLength param.Opt[int64]  `json:"max_output_length,omitzero"`
	TimeoutSeconds  param.Opt[int64]  `json:"timeout_seconds,omitzero"`
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
	Dx          param.Opt[float64] `json:"dx,omitzero"`
	Dy          param.Opt[float64] `json:"dy,omitzero"`
	Height      param.Opt[int64]   `json:"height,omitzero"`
	Ms          param.Opt[int64]   `json:"ms,omitzero"`
	ScaleFactor param.Opt[float64] `json:"scale_factor,omitzero"`
	// For tab management (browser sessions only)
	TabID param.Opt[string] `json:"tab_id,omitzero"`
	Text  param.Opt[string] `json:"text,omitzero"`
	// click|double_click|right_click|drag|type|keypress|scroll|wait|screenshot|go_to_url|debug|get_html_content|set_viewport|list_tabs|new_tab|switch_tab|close_tab
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
	MaxOutputLength param.Opt[int64]  `json:"max_output_length,omitzero"`
	TimeoutSeconds  param.Opt[int64]  `json:"timeout_seconds,omitzero"`
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
	AutoDetectEncoding param.Opt[bool] `json:"auto_detect_encoding,omitzero"`
	paramObj
}

func (r ComputerGetHTMLParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerGetHTMLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerGetHTMLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerNavigateParams struct {
	URL param.Opt[string] `json:"url,omitzero"`
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
	Keys []string `json:"keys,omitzero"`
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
	X param.Opt[float64] `json:"x,omitzero"`
	Y param.Opt[float64] `json:"y,omitzero"`
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
	Dx param.Opt[float64] `json:"dx,omitzero"`
	Dy param.Opt[float64] `json:"dy,omitzero"`
	X  param.Opt[float64] `json:"x,omitzero"`
	Y  param.Opt[float64] `json:"y,omitzero"`
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
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r ComputerTypeTextParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerTypeTextParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerTypeTextParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
