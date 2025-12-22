// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package computer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/computer-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/computer-go/internal/encoding/json"
	"github.com/stainless-sdks/computer-go/internal/requestconfig"
	"github.com/stainless-sdks/computer-go/option"
	"github.com/stainless-sdks/computer-go/packages/param"
	"github.com/stainless-sdks/computer-go/packages/ssestream"
)

// AgentTaskService contains methods and other services that help with interacting
// with the computer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentTaskService] method instead.
type AgentTaskService struct {
	Options []option.RequestOption
}

// NewAgentTaskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentTaskService(opts ...option.RequestOption) (r AgentTaskService) {
	r = AgentTaskService{}
	r.Options = opts
	return
}

// Get a list of all agent tasks for the authenticated user.
func (r *AgentTaskService) List(ctx context.Context, opts ...option.RequestOption) (res *[]AgentTaskListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the current status and metadata of an agent task.
func (r *AgentTaskService) GetStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *AgentTaskGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("agent/tasks/%s/status", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Inject a user message into a running agent task with debouncing and rate
// limiting. The message will interrupt the agent's current workflow and be
// processed after a 2-second debounce window.
func (r *AgentTaskService) SendMessage(ctx context.Context, taskID string, body AgentTaskSendMessageParams, opts ...option.RequestOption) (res *AgentTaskSendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("agent/tasks/%s/messages", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start a freeform agent task with custom instruction or structured messages. The
// task will execute browser automation actions based on the provided instruction
// or conversation history.
func (r *AgentTaskService) Start(ctx context.Context, params AgentTaskStartParams, opts ...option.RequestOption) (res *AgentTaskStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Start a task by predefined task ID. Currently unused - intended for OSWorld
// benchmark tasks with predefined IDs.
func (r *AgentTaskService) StartByID(ctx context.Context, taskID string, params AgentTaskStartByIDParams, opts ...option.RequestOption) (res *AgentTaskStartByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("agent/tasks/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Stream real-time updates for an agent task using Server-Sent Events (SSE).
//
// **Standardized Event Types:**
//
//   - `thinking`: Agent's internal thought process (content: thought text)
//   - `tool_call`: Agent action being executed (content: e.g. "click(x=100, y=200)",
//     "finished()")
//   - `message`: Final response to user (content: response text)
//   - `error`: Error occurred (content: error message)
//   - `image`: Generated image (content: image URL)
//
// **Agent-Specific Events (informational):**
//
// - `screenshot`: Browser screenshot available
// - `setup_start`, `setup_progress`, `setup_complete`: Task initialization
// - `computer_session_created`: Browser session ready
// - `stream_stopped`: Stream ended (reason: task_completed, no_action, etc.)
// - `stream_closed`: Stream closed by server (reason: inactivity_timeout, etc.)
// - `keep_alive`: Connection keep-alive ping
func (r *AgentTaskService) StreamUpdatesStreaming(ctx context.Context, taskID string, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("agent/tasks/%s/stream", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

type AgentTaskListResponse map[string]any

type AgentTaskGetStatusResponse map[string]any

type AgentTaskSendMessageResponse map[string]any

type AgentTaskStartResponse map[string]any

type AgentTaskStartByIDResponse map[string]any

type AgentTaskSendMessageParams struct {
	Body any
	paramObj
}

func (r AgentTaskSendMessageParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AgentTaskSendMessageParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type AgentTaskStartParams struct {
	Body any
	// Enable streaming via Server-Sent Events (SSE)
	Stream param.Opt[bool] `query:"stream,omitzero" json:"-"`
	paramObj
}

func (r AgentTaskStartParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AgentTaskStartParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [AgentTaskStartParams]'s query parameters as `url.Values`.
func (r AgentTaskStartParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentTaskStartByIDParams struct {
	Body any
	// Enable streaming via Server-Sent Events (SSE)
	Stream param.Opt[bool] `query:"stream,omitzero" json:"-"`
	paramObj
}

func (r AgentTaskStartByIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AgentTaskStartByIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [AgentTaskStartByIDParams]'s query parameters as
// `url.Values`.
func (r AgentTaskStartByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
