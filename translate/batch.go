package translate

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func Batch(ctx context.Context, core *Core, input []byte, from, to string) ([]byte, error) {
	if json.Valid(input) {
		var data any
		if err := json.Unmarshal(input, &data); err != nil {
			return nil, fmt.Errorf("invalid JSON: %w", err)
		}
		if err := translateValue(ctx, core, &data, from, to); err != nil {
			return nil, err
		}
		return json.MarshalIndent(data, "", "  ")
	}

	text := strings.TrimSpace(string(input))
	if text == "" {
		return nil, fmt.Errorf("empty input")
	}

	result, err := core.Backend.Translate(ctx, text, from, to)
	if err != nil {
		return nil, err
	}
	return []byte(result), nil
}

func translateValue(ctx context.Context, core *Core, val *any, from, to string) error {
	switch v := (*val).(type) {
	case string:
		if v == "" {
			return nil
		}
		result, err := core.Backend.Translate(ctx, v, from, to)
		if err != nil {
			return err
		}
		*val = result
		return nil
	case map[string]any:
		for k, child := range v {
			childCopy := child
			if err := translateValue(ctx, core, &childCopy, from, to); err != nil {
				return err
			}
			v[k] = childCopy
		}
	case []any:
		for i, child := range v {
			childCopy := child
			if err := translateValue(ctx, core, &childCopy, from, to); err != nil {
				return err
			}
			v[i] = childCopy
		}
	}
	return nil
}
