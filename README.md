# voca

**V**ersatile **O**ffline **C**ommunication **A**ssistant — real-time translation TUI powered by local Ollama models.

```bash
go run . --model llama3.2:3b
```

The binary manages the full Ollama lifecycle on its own: starts the server if offline, pulls the model if missing, and cleans up on exit.

## Benchmarks

See [BENCHMARKS.md](BENCHMARKS.md) for detailed multi-sentence translation benchmarks across 6 models and 14 languages.

## Model selection

```bash
go run .

# Fast, good for European languages
go run . --model llama3.2:3b

# Use a specific model
go run . --model <model-name>
```
