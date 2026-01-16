// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package computer_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/tzafon/computer-go"
	"github.com/tzafon/computer-go/internal/testutil"
	"github.com/tzafon/computer-go/option"
)

func TestComputerExecExecuteSyncWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := computer.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Computers.Exec.ExecuteSync(
		context.TODO(),
		"id",
		computer.ComputerExecExecuteSyncParams{
			Command: computer.String("command"),
			Cwd:     computer.String("cwd"),
			Env: map[string]string{
				"foo": "string",
			},
			TimeoutSeconds: computer.Int(0),
		},
	)
	if err != nil {
		var apierr *computer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
