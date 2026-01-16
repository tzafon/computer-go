// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package computer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/tzafon/computer-go/internal/apijson"
	"github.com/tzafon/computer-go/internal/requestconfig"
	"github.com/tzafon/computer-go/option"
	"github.com/tzafon/computer-go/packages/jsonl"
	"github.com/tzafon/computer-go/packages/param"
	"github.com/tzafon/computer-go/packages/respjson"
)

// ComputerExecService contains methods and other services that help with
// interacting with the computer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputerExecService] method instead.
type ComputerExecService struct {
	Options []option.RequestOption
}

// NewComputerExecService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputerExecService(opts ...option.RequestOption) (r ComputerExecService) {
	r = ComputerExecService{}
	r.Options = opts
	return
}

// Execute a shell command with real-time streaming output as NDJSON. Each line is
// a JSON object with type (stdout/stderr/exit/error).
func (r *ComputerExecService) ExecuteStreaming(ctx context.Context, id string, body ComputerExecExecuteParams, opts ...option.RequestOption) (stream *jsonl.Stream[ComputerExecExecuteResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/x-ndjson")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/exec", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return jsonl.NewStream[ComputerExecExecuteResponse](raw, err)
}

// Execute a shell command and wait for completion, returning buffered
// stdout/stderr.
func (r *ComputerExecService) ExecuteSync(ctx context.Context, id string, body ComputerExecExecuteSyncParams, opts ...option.RequestOption) (res *ComputerExecExecuteSyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/exec/sync", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ComputerExecExecuteResponse struct {
	// for exit
	Code int64 `json:"code"`
	// for stdout/stderr
	Data string `json:"data"`
	// for error
	Message string `json:"message"`
	// "stdout", "stderr", "exit", "error"
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Data        respjson.Field
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerExecExecuteResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerExecExecuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecExecuteSyncResponse struct {
	ExitCode int64  `json:"exit_code"`
	Stderr   string `json:"stderr"`
	Stdout   string `json:"stdout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExitCode    respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputerExecExecuteSyncResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputerExecExecuteSyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecExecuteParams struct {
	Command        param.Opt[string] `json:"command,omitzero"`
	Cwd            param.Opt[string] `json:"cwd,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	Env            map[string]string `json:"env,omitzero"`
	paramObj
}

func (r ComputerExecExecuteParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerExecExecuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerExecExecuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputerExecExecuteSyncParams struct {
	Command        param.Opt[string] `json:"command,omitzero"`
	Cwd            param.Opt[string] `json:"cwd,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	Env            map[string]string `json:"env,omitzero"`
	paramObj
}

func (r ComputerExecExecuteSyncParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputerExecExecuteSyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputerExecExecuteSyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
