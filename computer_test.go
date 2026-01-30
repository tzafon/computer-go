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

func TestComputerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.New(context.TODO(), githubcomtzafoncomputergo.ComputerNewParams{
		AutoKill:  githubcomtzafoncomputergo.Bool(true),
		ContextID: githubcomtzafoncomputergo.String("context_id"),
		Display: githubcomtzafoncomputergo.ComputerNewParamsDisplay{
			Height: githubcomtzafoncomputergo.Int(0),
			Scale:  githubcomtzafoncomputergo.Float(0),
			Width:  githubcomtzafoncomputergo.Int(0),
		},
		EnvironmentID:            githubcomtzafoncomputergo.String("environment_id"),
		InactivityTimeoutSeconds: githubcomtzafoncomputergo.Int(0),
		Kind:                     githubcomtzafoncomputergo.String("kind"),
		Persistent:               githubcomtzafoncomputergo.Bool(true),
		Stealth:                  map[string]any{},
		TimeoutSeconds:           githubcomtzafoncomputergo.Int(0),
		UseAdvancedProxy:         githubcomtzafoncomputergo.Bool(true),
	})
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
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
	client := githubcomtzafoncomputergo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Computers.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.List(context.TODO(), githubcomtzafoncomputergo.ComputerListParams{
		Type: githubcomtzafoncomputergo.ComputerListParamsTypeLive,
	})
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerCaptureScreenshotWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.CaptureScreenshot(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerCaptureScreenshotParams{
			Base64: githubcomtzafoncomputergo.Bool(true),
			TabID:  githubcomtzafoncomputergo.String("tab_id"),
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

func TestComputerChangeProxyWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.ChangeProxy(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerChangeProxyParams{
			ProxyURL: githubcomtzafoncomputergo.String("proxy_url"),
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

func TestComputerClickWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.Click(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerClickParams{
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			X:     githubcomtzafoncomputergo.Float(0),
			Y:     githubcomtzafoncomputergo.Float(0),
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

func TestComputerConnectWebsocket(t *testing.T) {
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
	err := client.Computers.ConnectWebsocket(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerDebugWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.Debug(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerDebugParams{
			Command:         githubcomtzafoncomputergo.String("command"),
			MaxOutputLength: githubcomtzafoncomputergo.Int(0),
			TabID:           githubcomtzafoncomputergo.String("tab_id"),
			TimeoutSeconds:  githubcomtzafoncomputergo.Int(0),
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

func TestComputerDoubleClickWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.DoubleClick(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerDoubleClickParams{
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			X:     githubcomtzafoncomputergo.Float(0),
			Y:     githubcomtzafoncomputergo.Float(0),
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

func TestComputerDragWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.Drag(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerDragParams{
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			X1:    githubcomtzafoncomputergo.Float(0),
			X2:    githubcomtzafoncomputergo.Float(0),
			Y1:    githubcomtzafoncomputergo.Float(0),
			Y2:    githubcomtzafoncomputergo.Float(0),
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

func TestComputerExecuteActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.ExecuteAction(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerExecuteActionParams{
			Action: githubcomtzafoncomputergo.ComputerExecuteActionParamsAction{
				AutoDetectEncoding: githubcomtzafoncomputergo.Bool(true),
				Base64:             githubcomtzafoncomputergo.Bool(true),
				Button:             githubcomtzafoncomputergo.String("button"),
				Debug: githubcomtzafoncomputergo.ComputerExecuteActionParamsActionDebug{
					Command: githubcomtzafoncomputergo.String("command"),
					Cwd:     githubcomtzafoncomputergo.String("cwd"),
					Env: map[string]string{
						"foo": "string",
					},
					MaxOutputLength: githubcomtzafoncomputergo.Int(0),
					RequestID:       githubcomtzafoncomputergo.String("request_id"),
					Stream:          githubcomtzafoncomputergo.Bool(true),
					TimeoutSeconds:  githubcomtzafoncomputergo.Int(0),
				},
				Dx:             githubcomtzafoncomputergo.Float(0),
				Dy:             githubcomtzafoncomputergo.Float(0),
				Height:         githubcomtzafoncomputergo.Int(0),
				IncludeContext: githubcomtzafoncomputergo.Bool(true),
				Key:            githubcomtzafoncomputergo.String("key"),
				Keys:           []string{"string"},
				Ms:             githubcomtzafoncomputergo.Int(0),
				ProxyURL:       githubcomtzafoncomputergo.String("proxy_url"),
				RequestID:      githubcomtzafoncomputergo.String("request_id"),
				ScaleFactor:    githubcomtzafoncomputergo.Float(0),
				TabID:          githubcomtzafoncomputergo.String("tab_id"),
				Text:           githubcomtzafoncomputergo.String("text"),
				Type:           githubcomtzafoncomputergo.String("type"),
				URL:            githubcomtzafoncomputergo.String("url"),
				Width:          githubcomtzafoncomputergo.Int(0),
				X:              githubcomtzafoncomputergo.Float(0),
				X1:             githubcomtzafoncomputergo.Float(0),
				X2:             githubcomtzafoncomputergo.Float(0),
				Y:              githubcomtzafoncomputergo.Float(0),
				Y1:             githubcomtzafoncomputergo.Float(0),
				Y2:             githubcomtzafoncomputergo.Float(0),
			},
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

func TestComputerExecuteBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.ExecuteBatch(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerExecuteBatchParams{
			Actions: []githubcomtzafoncomputergo.ComputerExecuteBatchParamsAction{{
				AutoDetectEncoding: githubcomtzafoncomputergo.Bool(true),
				Base64:             githubcomtzafoncomputergo.Bool(true),
				Button:             githubcomtzafoncomputergo.String("button"),
				Debug: githubcomtzafoncomputergo.ComputerExecuteBatchParamsActionDebug{
					Command: githubcomtzafoncomputergo.String("command"),
					Cwd:     githubcomtzafoncomputergo.String("cwd"),
					Env: map[string]string{
						"foo": "string",
					},
					MaxOutputLength: githubcomtzafoncomputergo.Int(0),
					RequestID:       githubcomtzafoncomputergo.String("request_id"),
					Stream:          githubcomtzafoncomputergo.Bool(true),
					TimeoutSeconds:  githubcomtzafoncomputergo.Int(0),
				},
				Dx:             githubcomtzafoncomputergo.Float(0),
				Dy:             githubcomtzafoncomputergo.Float(0),
				Height:         githubcomtzafoncomputergo.Int(0),
				IncludeContext: githubcomtzafoncomputergo.Bool(true),
				Key:            githubcomtzafoncomputergo.String("key"),
				Keys:           []string{"string"},
				Ms:             githubcomtzafoncomputergo.Int(0),
				ProxyURL:       githubcomtzafoncomputergo.String("proxy_url"),
				RequestID:      githubcomtzafoncomputergo.String("request_id"),
				ScaleFactor:    githubcomtzafoncomputergo.Float(0),
				TabID:          githubcomtzafoncomputergo.String("tab_id"),
				Text:           githubcomtzafoncomputergo.String("text"),
				Type:           githubcomtzafoncomputergo.String("type"),
				URL:            githubcomtzafoncomputergo.String("url"),
				Width:          githubcomtzafoncomputergo.Int(0),
				X:              githubcomtzafoncomputergo.Float(0),
				X1:             githubcomtzafoncomputergo.Float(0),
				X2:             githubcomtzafoncomputergo.Float(0),
				Y:              githubcomtzafoncomputergo.Float(0),
				Y1:             githubcomtzafoncomputergo.Float(0),
				Y2:             githubcomtzafoncomputergo.Float(0),
			}},
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

func TestComputerGetHTMLWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.GetHTML(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerGetHTMLParams{
			AutoDetectEncoding: githubcomtzafoncomputergo.Bool(true),
			TabID:              githubcomtzafoncomputergo.String("tab_id"),
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

func TestComputerKeepAlive(t *testing.T) {
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
	_, err := client.Computers.KeepAlive(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerKeyDownWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.KeyDown(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerKeyDownParams{
			Key:   githubcomtzafoncomputergo.String("shift"),
			TabID: githubcomtzafoncomputergo.String("tab_id"),
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

func TestComputerKeyUpWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.KeyUp(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerKeyUpParams{
			Key:   githubcomtzafoncomputergo.String("shift"),
			TabID: githubcomtzafoncomputergo.String("tab_id"),
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

func TestComputerMouseDownWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.MouseDown(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerMouseDownParams{
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			X:     githubcomtzafoncomputergo.Float(0),
			Y:     githubcomtzafoncomputergo.Float(0),
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

func TestComputerMouseUpWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.MouseUp(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerMouseUpParams{
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			X:     githubcomtzafoncomputergo.Float(0),
			Y:     githubcomtzafoncomputergo.Float(0),
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

func TestComputerNavigateWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.Navigate(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerNavigateParams{
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			URL:   githubcomtzafoncomputergo.String("url"),
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

func TestComputerPressHotkeyWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.PressHotkey(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerPressHotkeyParams{
			Keys:  []string{"string"},
			TabID: githubcomtzafoncomputergo.String("tab_id"),
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

func TestComputerGetStatus(t *testing.T) {
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
	_, err := client.Computers.GetStatus(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerRightClickWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.RightClick(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerRightClickParams{
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			X:     githubcomtzafoncomputergo.Float(0),
			Y:     githubcomtzafoncomputergo.Float(0),
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

func TestComputerScrollViewportWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.ScrollViewport(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerScrollViewportParams{
			Dx:    githubcomtzafoncomputergo.Float(0),
			Dy:    githubcomtzafoncomputergo.Float(0),
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			X:     githubcomtzafoncomputergo.Float(0),
			Y:     githubcomtzafoncomputergo.Float(0),
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

func TestComputerSetViewportWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.SetViewport(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerSetViewportParams{
			Height:      githubcomtzafoncomputergo.Int(0),
			ScaleFactor: githubcomtzafoncomputergo.Float(0),
			TabID:       githubcomtzafoncomputergo.String("tab_id"),
			Width:       githubcomtzafoncomputergo.Int(0),
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

func TestComputerStreamEvents(t *testing.T) {
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
	err := client.Computers.StreamEvents(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerStreamScreencast(t *testing.T) {
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
	err := client.Computers.StreamScreencast(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
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
	client := githubcomtzafoncomputergo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Computers.Terminate(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomtzafoncomputergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputerTypeTextWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.TypeText(
		context.TODO(),
		"id",
		githubcomtzafoncomputergo.ComputerTypeTextParams{
			TabID: githubcomtzafoncomputergo.String("tab_id"),
			Text:  githubcomtzafoncomputergo.String("text"),
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
