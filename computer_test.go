// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package computer_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/computer-go"
	"github.com/stainless-sdks/computer-go/internal/testutil"
	"github.com/stainless-sdks/computer-go/option"
)

func TestComputerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.New(context.TODO(), computer.ComputerNewParams{
		AutoKill:  computer.Bool(true),
		ContextID: computer.String("context_id"),
		Display: computer.ComputerNewParamsDisplay{
			Height: computer.Int(0),
			Scale:  computer.Float(0),
			Width:  computer.Int(0),
		},
		Kind:           computer.String("kind"),
		Stealth:        map[string]interface{}{},
		TimeoutSeconds: computer.Int(0),
	})
	if err != nil {
		var apierr *computer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerGet(t *testing.T) {
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
	_, err := client.Computers.Get(context.TODO(), "id")
	if err != nil {
		var apierr *computer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerList(t *testing.T) {
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
	_, err := client.Computers.List(context.TODO())
	if err != nil {
		var apierr *computer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerExecuteActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.ExecuteAction(
		context.TODO(),
		"id",
		computer.ComputerExecuteActionParams{
			Action: computer.ComputerExecuteActionParamsAction{
				AutoDetectEncoding: computer.Bool(true),
				Base64:             computer.Bool(true),
				Button:             computer.String("button"),
				Debug: computer.ComputerExecuteActionParamsActionDebug{
					Command:         computer.String("command"),
					MaxOutputLength: computer.Int(0),
					TimeoutSeconds:  computer.Int(0),
				},
				Dx:          computer.Float(0),
				Dy:          computer.Float(0),
				Height:      computer.Int(0),
				Keys:        []string{"string"},
				Ms:          computer.Int(0),
				ScaleFactor: computer.Float(0),
				Text:        computer.String("text"),
				Type:        computer.String("type"),
				URL:         computer.String("url"),
				Width:       computer.Int(0),
				X:           computer.Float(0),
				X1:          computer.Float(0),
				X2:          computer.Float(0),
				Y:           computer.Float(0),
				Y1:          computer.Float(0),
				Y2:          computer.Float(0),
			},
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

func TestComputerExecuteBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.ExecuteBatch(
		context.TODO(),
		"id",
		computer.ComputerExecuteBatchParams{
			Actions: []computer.ComputerExecuteBatchParamsAction{{
				AutoDetectEncoding: computer.Bool(true),
				Base64:             computer.Bool(true),
				Button:             computer.String("button"),
				Debug: computer.ComputerExecuteBatchParamsActionDebug{
					Command:         computer.String("command"),
					MaxOutputLength: computer.Int(0),
					TimeoutSeconds:  computer.Int(0),
				},
				Dx:          computer.Float(0),
				Dy:          computer.Float(0),
				Height:      computer.Int(0),
				Keys:        []string{"string"},
				Ms:          computer.Int(0),
				ScaleFactor: computer.Float(0),
				Text:        computer.String("text"),
				Type:        computer.String("type"),
				URL:         computer.String("url"),
				Width:       computer.Int(0),
				X:           computer.Float(0),
				X1:          computer.Float(0),
				X2:          computer.Float(0),
				Y:           computer.Float(0),
				Y1:          computer.Float(0),
				Y2:          computer.Float(0),
			}},
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

func TestComputerKeepAlive(t *testing.T) {
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
	_, err := client.Computers.KeepAlive(context.TODO(), "id")
	if err != nil {
		var apierr *computer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerNavigateWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.Navigate(
		context.TODO(),
		"id",
		computer.ComputerNavigateParams{
			URL: computer.String("url"),
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

func TestComputerStreamEvents(t *testing.T) {
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
	err := client.Computers.StreamEvents(context.TODO(), "id")
	if err != nil {
		var apierr *computer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerTerminate(t *testing.T) {
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
	err := client.Computers.Terminate(context.TODO(), "id")
	if err != nil {
		var apierr *computer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
