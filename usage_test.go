// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomtzafoncomputergo_test

import (
	"context"
	"os"
	"testing"

	"github.com/tzafon/computer-go"
	"github.com/tzafon/computer-go/internal/testutil"
	"github.com/tzafon/computer-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	computerResponses, err := client.Computers.List(context.TODO(), githubcomtzafoncomputergo.ComputerListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", computerResponses)
}
