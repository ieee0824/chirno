package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthChecker(t *testing.T) {
	h := &HealthChecker{}
	req := httptest.NewRequest(http.MethodGet, "/health_check", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(body); got != "It works!" {
		t.Errorf("body = %q, want %q", got, "It works!")
	}
}

func TestGenProxy_NilConfig(t *testing.T) {
	proxy := genProxy(nil)
	if proxy != nil {
		t.Error("expected nil proxy for nil config")
	}
}

func TestGenProxy_EmptyBackendURL(t *testing.T) {
	cnf := &Config{BackendURL: ""}
	proxy := genProxy(cnf)
	if proxy != nil {
		t.Error("expected nil proxy for empty BackendURL")
	}
}

func TestGenProxy_ValidBackendURL(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("backend"))
	}))
	defer backend.Close()

	cnf := &Config{BackendURL: backend.URL}
	proxy := genProxy(cnf)
	if proxy == nil {
		t.Fatal("expected non-nil proxy")
	}

	// proxy 経由でバックエンドにリクエストが転送されることを確認
	frontend := httptest.NewServer(proxy)
	defer frontend.Close()

	resp, err := http.Get(frontend.URL + "/test")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(body); got != "backend" {
		t.Errorf("body = %q, want %q", got, "backend")
	}
}

func TestRunTargetProcess_NilConfig(t *testing.T) {
	err := runTargetProcess(nil)
	if err == nil {
		t.Error("expected error for nil config")
	}
}

func TestRunTargetProcess_EmptyCommand(t *testing.T) {
	cnf := &Config{Command: []string{}}
	err := runTargetProcess(cnf)
	if err == nil {
		t.Error("expected error for empty command")
	}
}

func TestRunTargetProcess_SingleCommand(t *testing.T) {
	cnf := &Config{Command: []string{"true"}}
	if err := runTargetProcess(cnf); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRunTargetProcess_MultipleArgs(t *testing.T) {
	cnf := &Config{Command: []string{"echo", "hello"}}
	if err := runTargetProcess(cnf); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRunTargetProcess_InvalidCommand(t *testing.T) {
	cnf := &Config{Command: []string{"nonexistent_command_xyz"}}
	err := runTargetProcess(cnf)
	if err == nil {
		t.Error("expected error for invalid command")
	}
}
