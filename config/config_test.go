package config

import (
	"os"
	"testing"
)

func TestDefault_BackendType(t *testing.T) {
	cfg := Default()
	if cfg.Backend.Type != "ollama" {
		t.Fatalf("expected ollama, got %s", cfg.Backend.Type)
	}
}

func TestDefault_BackendModel(t *testing.T) {
	cfg := Default()
	if cfg.Backend.Model != DefaultModel {
		t.Fatalf("expected %s, got %s", DefaultModel, cfg.Backend.Model)
	}
}

func TestDefault_BackendBaseURL(t *testing.T) {
	cfg := Default()
	if cfg.Backend.BaseURL != DefaultBaseURL {
		t.Fatalf("expected %s, got %s", DefaultBaseURL, cfg.Backend.BaseURL)
	}
}

func TestDefault_BackendOptions(t *testing.T) {
	cfg := Default()
	if cfg.Backend.Options == nil {
		t.Fatal("options should not be nil")
	}
	if cfg.Backend.Options["num_predict"] != float64(2048) {
		t.Fatalf("expected num_predict 2048, got %v", cfg.Backend.Options["num_predict"])
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	cfg, err := Load("/nonexistent/path/config.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Backend.Model != DefaultModel {
		t.Fatalf("expected defaults on missing file, got %s", cfg.Backend.Model)
	}
}

func TestLoad_CustomFile(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/test.yaml"
	data := []byte("backend:\n  model: custom-model\n")
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Backend.Model != "custom-model" {
		t.Fatalf("expected custom-model, got %s", cfg.Backend.Model)
	}
	if cfg.Backend.BaseURL != DefaultBaseURL {
		t.Fatalf("expected default base_url %s, got %s", DefaultBaseURL, cfg.Backend.BaseURL)
	}
}

func TestLoad_PartialOverride(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/partial.yaml"
	data := []byte("backend:\n  base_url: http://custom:11434\n")
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Backend.BaseURL != "http://custom:11434" {
		t.Fatalf("expected custom url, got %s", cfg.Backend.BaseURL)
	}
	if cfg.Backend.Model != DefaultModel {
		t.Fatalf("expected default model, got %s", cfg.Backend.Model)
	}
}
