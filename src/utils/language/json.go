package language

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type JSONLanguage struct {
	locale          string
	languages       map[string]Content
	mu              sync.RWMutex
	defaultLanguage string
}

func NewJSONLanguage(locale, defaultLanguage string) *JSONLanguage {
	return &JSONLanguage{
		locale:          locale,
		languages:       make(map[string]Content),
		mu:              sync.RWMutex{},
		defaultLanguage: defaultLanguage,
	}
}

func (jl *JSONLanguage) LoadLanguages() error {
	jl.mu.Lock()
	defer jl.mu.Unlock()

	files, err := filepath.Glob(filepath.Join(jl.locale, "*.json"))
	if err != nil {
		return err
	}

	for _, file := range files {
		lang := filepath.Base(file[:len(file)-len(filepath.Ext(file))])

		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read locale file %s: %w", file, err)
		}

		var content Content
		if err := json.Unmarshal(data, &content); err != nil {
			return fmt.Errorf("failed to parse JSON in %s: %w", file, err)
		}

		jl.languages[lang] = content
	}

	return nil
}

func (jl *JSONLanguage) L(ctx context.Context, key string) Content {
	if key == "" {
		key = jl.defaultLanguage
	}

	jl.mu.RLock()
	defer jl.mu.RUnlock()

	if content, exists := jl.languages[key]; exists {
		return content
	}

	return jl.languages[jl.defaultLanguage]
}
