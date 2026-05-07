// Package main_test contains a single black-box smoke test that drives the
// API over real HTTP. It is intentionally minimal — it does not constrain
// your internal package layout or your choice of router.
//
// The test is expected to FAIL on a fresh checkout. Make it pass as the
// first step of the task.
package main_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// baseURL boots the API binary in a child process unless GLOFOX_API_URL is
// set, in which case the test hits that URL instead (useful with `make run`).
func baseURL(t *testing.T) (string, func()) {
	t.Helper()

	if url := os.Getenv("GLOFOX_API_URL"); url != "" {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		return strings.TrimRight(url, "/"), func() {}
	}

	const addr = "127.0.0.1:18080"
	cmd := exec.Command("go", "run", "./cmd/api")
	cmd.Env = append(os.Environ(), "ADDR="+addr)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	if err := cmd.Start(); err != nil {
		t.Fatalf("start api: %v", err)
	}

	url := "http://" + addr
	if !waitForReady(url, 5*time.Second) {
		_ = cmd.Process.Kill()
		t.Fatalf("api did not become ready at %s", url)
	}

	return url, func() { _ = cmd.Process.Kill() }
}

func TestCreateClass_HappyPath(t *testing.T) {
	url, stop := baseURL(t)
	defer stop()

	body := map[string]any{
		"name":       "Pilates",
		"start_date": "2026-12-01",
		"end_date":   "2026-12-20",
		"capacity":   10,
	}
	raw, _ := json.Marshal(body)

	resp, err := http.Post(url+"/classes", "application/json", bytes.NewReader(raw))
	if err != nil {
		t.Fatalf("POST /classes: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		got, _ := io.ReadAll(resp.Body)
		t.Fatalf("want 201 Created, got %d. body=%s", resp.StatusCode, strings.TrimSpace(string(got)))
	}
}

func waitForReady(url string, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	client := &http.Client{Timeout: 200 * time.Millisecond}
	for time.Now().Before(deadline) {
		resp, err := client.Get(url + "/")
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(50 * time.Millisecond)
	}
	return false
}
