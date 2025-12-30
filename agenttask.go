// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package computer

import (
	"github.com/stainless-sdks/computer-go/option"
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
