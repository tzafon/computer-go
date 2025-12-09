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
		Stealth:        map[string]any{},
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

func TestComputerCaptureScreenshotWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.CaptureScreenshot(
		context.TODO(),
		"id",
		computer.ComputerCaptureScreenshotParams{
			Base64: computer.Bool(true),
			TabID:  computer.String("tab_id"),
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

func TestComputerClickWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.Click(
		context.TODO(),
		"id",
		computer.ComputerClickParams{
			TabID: computer.String("tab_id"),
			X:     computer.Float(0),
			Y:     computer.Float(0),
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

func TestComputerConnectWebsocket(t *testing.T) {
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
	err := client.Computers.ConnectWebsocket(context.TODO(), "id")
	if err != nil {
		var apierr *computer.Error
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
	client := computer.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Computers.Debug(
		context.TODO(),
		"id",
		computer.ComputerDebugParams{
			Command:         computer.String("command"),
			MaxOutputLength: computer.Int(0),
			TabID:           computer.String("tab_id"),
			TimeoutSeconds:  computer.Int(0),
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

func TestComputerDoubleClickWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.DoubleClick(
		context.TODO(),
		"id",
		computer.ComputerDoubleClickParams{
			TabID: computer.String("tab_id"),
			X:     computer.Float(0),
			Y:     computer.Float(0),
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

func TestComputerDragWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.Drag(
		context.TODO(),
		"id",
		computer.ComputerDragParams{
			TabID: computer.String("tab_id"),
			X1:    computer.Float(0),
			X2:    computer.Float(0),
			Y1:    computer.Float(0),
			Y2:    computer.Float(0),
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
				Dx:             computer.Float(0),
				Dy:             computer.Float(0),
				Height:         computer.Int(0),
				IncludeContext: computer.Bool(true),
				Key:            computer.String("key"),
				Keys:           []string{"string"},
				Ms:             computer.Int(0),
				ScaleFactor:    computer.Float(0),
				TabID:          computer.String("tab_id"),
				Text:           computer.String("text"),
				Type:           computer.String("type"),
				URL:            computer.String("url"),
				Width:          computer.Int(0),
				X:              computer.Float(0),
				X1:             computer.Float(0),
				X2:             computer.Float(0),
				Y:              computer.Float(0),
				Y1:             computer.Float(0),
				Y2:             computer.Float(0),
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
				Dx:             computer.Float(0),
				Dy:             computer.Float(0),
				Height:         computer.Int(0),
				IncludeContext: computer.Bool(true),
				Key:            computer.String("key"),
				Keys:           []string{"string"},
				Ms:             computer.Int(0),
				ScaleFactor:    computer.Float(0),
				TabID:          computer.String("tab_id"),
				Text:           computer.String("text"),
				Type:           computer.String("type"),
				URL:            computer.String("url"),
				Width:          computer.Int(0),
				X:              computer.Float(0),
				X1:             computer.Float(0),
				X2:             computer.Float(0),
				Y:              computer.Float(0),
				Y1:             computer.Float(0),
				Y2:             computer.Float(0),
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

func TestComputerGetHTMLWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.GetHTML(
		context.TODO(),
		"id",
		computer.ComputerGetHTMLParams{
			AutoDetectEncoding: computer.Bool(true),
			TabID:              computer.String("tab_id"),
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

func TestComputerKeyDownWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.KeyDown(
		context.TODO(),
		"id",
		computer.ComputerKeyDownParams{
			Key:   computer.String("key"),
			TabID: computer.String("tab_id"),
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

func TestComputerKeyUpWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.KeyUp(
		context.TODO(),
		"id",
		computer.ComputerKeyUpParams{
			Key:   computer.String("key"),
			TabID: computer.String("tab_id"),
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

func TestComputerMouseDownWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.MouseDown(
		context.TODO(),
		"id",
		computer.ComputerMouseDownParams{
			TabID: computer.String("tab_id"),
			X:     computer.Float(0),
			Y:     computer.Float(0),
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

func TestComputerMouseUpWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.MouseUp(
		context.TODO(),
		"id",
		computer.ComputerMouseUpParams{
			TabID: computer.String("tab_id"),
			X:     computer.Float(0),
			Y:     computer.Float(0),
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
			TabID: computer.String("tab_id"),
			URL:   computer.String("url"),
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

func TestComputerPressHotkeyWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.PressHotkey(
		context.TODO(),
		"id",
		computer.ComputerPressHotkeyParams{
			Keys:  []string{"string"},
			TabID: computer.String("tab_id"),
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

func TestComputerRightClickWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.RightClick(
		context.TODO(),
		"id",
		computer.ComputerRightClickParams{
			TabID: computer.String("tab_id"),
			X:     computer.Float(0),
			Y:     computer.Float(0),
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

func TestComputerScrollViewportWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.ScrollViewport(
		context.TODO(),
		"id",
		computer.ComputerScrollViewportParams{
			Dx:    computer.Float(0),
			Dy:    computer.Float(0),
			TabID: computer.String("tab_id"),
			X:     computer.Float(0),
			Y:     computer.Float(0),
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

func TestComputerSetViewportWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.SetViewport(
		context.TODO(),
		"id",
		computer.ComputerSetViewportParams{
			Height:      computer.Int(0),
			ScaleFactor: computer.Float(0),
			TabID:       computer.String("tab_id"),
			Width:       computer.Int(0),
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

func TestComputerStreamScreencast(t *testing.T) {
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
	err := client.Computers.StreamScreencast(context.TODO(), "id")
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

func TestComputerTypeTextWithOptionalParams(t *testing.T) {
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
	_, err := client.Computers.TypeText(
		context.TODO(),
		"id",
		computer.ComputerTypeTextParams{
			TabID: computer.String("tab_id"),
			Text:  computer.String("text"),
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
