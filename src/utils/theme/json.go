package theme

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type JSONTheme struct {
	locale       string
	themes       map[string]Mode
	mu           sync.RWMutex
	defaultTheme string
}

func NewJSONTheme(locale, defaultTheme string) *JSONTheme {
	return &JSONTheme{
		locale:       locale,
		themes:       make(map[string]Mode),
		mu:           sync.RWMutex{},
		defaultTheme: defaultTheme,
	}
}

func (jt *JSONTheme) LoadThemes() error {
	jt.mu.Lock()
	defer jt.mu.Unlock()

	files, err := filepath.Glob(filepath.Join(jt.locale, "*.json"))
	if err != nil {
		return err
	}

	for _, file := range files {
		t := filepath.Base(file[:len(file)-len(filepath.Ext(file))])

		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read locale file %s: %w", file, err)
		}

		var mode Mode
		if err := json.Unmarshal(data, &mode); err != nil {
			return fmt.Errorf("failed to parse JSON in %s: %w", file, err)
		}

		jt.themes[t] = mode
	}

	return nil
}

func (jt *JSONTheme) T(ctx context.Context, key string) Mode {
	if key == "" {
		key = jt.defaultTheme
	}

	jt.mu.RLock()
	defer jt.mu.RUnlock()

	if mode, exists := jt.themes[key]; exists {
		return mode
	}

	return jt.themes[jt.defaultTheme]
}
