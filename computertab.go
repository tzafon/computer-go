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

// Create a new tab, optionally navigating to a URL. The new tab becomes the main
// tab (browser sessions only).
func (r *ComputerTabService) New(ctx context.Context, id string, body ComputerTabNewParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("computers/%s/tabs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of open tabs with IDs, URLs, titles, and main tab status (browser
// sessions only). Includes external CDP pages (e.g., Playwright). Excludes
// devtools:// and chrome:// tabs. Results may be eventually consistent for newly
// created tabs.
func (r *ComputerTabService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *ActionResult, err error) {
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
// sessions only). Tab IDs come from ListTabs.
func (r *ComputerTabService) Delete(ctx context.Context, tabID string, body ComputerTabDeleteParams, opts ...option.RequestOption) (res *ActionResult, err error) {
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

// Switch the main/active tab to a different tab by ID (browser sessions only). Tab
// IDs come from ListTabs.
func (r *ComputerTabService) Switch(ctx context.Context, tabID string, body ComputerTabSwitchParams, opts ...option.RequestOption) (res *ActionResult, err error) {
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
