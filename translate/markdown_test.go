package translate

import (
	"context"
	"testing"
)

func TestTranslateMarkdown_CodeBlock(t *testing.T) {
	input := "# Hello\n\nSome text\n\n```go\nfunc main() {\n\tfmt.Println(\"hello\")\n}\n```\n\n## Getting Started\n\n### Installation\n\nMore text"
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	result, err := TranslateMarkdown(context.Background(), core, input, "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "# [en->it] Hello\n\n[en->it] Some text\n\n```go\nfunc main() {\n\tfmt.Println(\"hello\")\n}\n```\n\n## [en->it] Getting Started\n\n### [en->it] Installation\n\n[en->it] More text"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}

func TestTranslateMarkdown_Blockquote(t *testing.T) {
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	result, err := TranslateMarkdown(context.Background(), core, "> This is a quote\n\n> Nested?\n> Still quote", "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "> [en->it] This is a quote\n\n> [en->it] Nested?\n> [en->it] Still quote"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}

func TestTranslateMarkdown_List(t *testing.T) {
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	result, err := TranslateMarkdown(context.Background(), core, "- item one\n- item two\n* star item\n+ plus item\n1. first\n2. second", "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "- [en->it] item one\n- [en->it] item two\n* [en->it] star item\n+ [en->it] plus item\n1. [en->it] first\n2. [en->it] second"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}

func TestTranslateMarkdown_ThematicBreak(t *testing.T) {
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	result, err := TranslateMarkdown(context.Background(), core, "before\n\n---\n\nafter", "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "[en->it] before\n\n---\n\n[en->it] after"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}

func TestTranslateMarkdown_NoTrailingNewline(t *testing.T) {
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	result, err := TranslateMarkdown(context.Background(), core, "just one line", "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "[en->it] just one line"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}

func TestTranslateMarkdown_EmptyInput(t *testing.T) {
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	result, err := TranslateMarkdown(context.Background(), core, "", "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "" {
		t.Fatalf("expected empty, got %q", result)
	}
}

func TestTranslateMarkdown_InlineLinks(t *testing.T) {
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	result, err := TranslateMarkdown(context.Background(), core, "See [installation](#installation) for details", "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "[en->it] See [installation](#installation) for details"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}

func TestTranslateMarkdown_ImageLinks(t *testing.T) {
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	input := "![alt](img.png) and [text](link.com)"
	result, err := TranslateMarkdown(context.Background(), core, input, "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "[en->it] ![alt](img.png) and [text](link.com)"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}

func TestTranslateMarkdown_LinkWithImage(t *testing.T) {
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	input := "[![Go](badge.svg)](https://go.dev)"
	result, err := TranslateMarkdown(context.Background(), core, input, "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// only a protected link, no translatable content → returned as-is
	want := "[![Go](badge.svg)](https://go.dev)"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}

func TestTranslateMarkdown_NestedFences(t *testing.T) {
	input := "text\n\n```\nouter\n```inside\n```\nstill outer\n```\n\nend"
	core := NewTranslator(NewMockBackend(), NewStaticLanguages())
	result, err := TranslateMarkdown(context.Background(), core, input, "en", "it")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "[en->it] text\n\n```\nouter\n```inside\n```\nstill outer\n```\n\n[en->it] end"
	if result != want {
		t.Fatalf("got:\n%s\n\nwant:\n%s", result, want)
	}
}
