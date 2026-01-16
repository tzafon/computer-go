// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomtzafoncomputergo_test

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
	client := githubcomtzafoncomputergo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Computers.Exec.ExecuteSync(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerExecExecuteSyncParams{
			Command: githubcomtzafoncomputergo.String("command"),
			Cwd:     githubcomtzafoncomputergo.String("cwd"),
			Env: map[string]string{
				"foo": "string",
			},
			TimeoutSeconds: githubcomtzafoncomputergo.Int(0),
		},
	)
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
