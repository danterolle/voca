# Contributing

**PRs and any contributions aimed at improving this project are welcome**: open issues, suggest ideas, submit merge requests and provide any other help.

## Commit style

Use conventional commits:

    feat: add new feature
    fix: correct a bug
    refactor: restructure without changing behavior
    style: formatting, gofmt, lint
    docs: documentation only
    test: add or fix tests

Before pushing please use:

```sh
go build ./...
go test -race ./...
gofmt -l .
go vet ./...
golangci-lint run
```

CI runs the same checks.

## Adding a backend

1. Create `translate/<name>/backend.go`
2. Define request/response types
3. Implement `Translate(ctx, text, source, target) (string, error)` using `httpclient.PostJSON`
4. Add `case "<name>"` in `translate/factory.go`

See `translate/ollama/backend.go` or `translate/llamacpp/backend.go` as template.

## Project layout

```
cmd/loqi/           — entry point + CLI dispatch + flags
cmd/bench/          — standalone benchmark tool
translate/          — core domain (interfaces, languages, factory)
translate/http/     — shared HTTP client and backend config
translate/ollama/   — Ollama backend
translate/llamacpp/ — llama.cpp backend
translate/argos/    — argos-translate backend (embedded Python server)
translate/setup/    — backend lifecycle (start, health, cleanup)
tui/                — Bubble Tea TUI (model, update, view, styles)
config/             — YAML config loading
docs/               — technical docs, benchmarks, landing page
test_data/          — test fixtures
```

## Versioning

Version is injected at build time:

```sh
make build        # uses git describe --tags via ldflags
make release      # builds for linux/darwin/windows + checksums
```
