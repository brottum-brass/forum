package middleware

import (
	"context"
	"net/http"

	"github.com/brottum-brass/forum/src/utils"
)

type contextKey string

const (
	LanguageKey contextKey = "language"
	ThemeKey    contextKey = "theme"
)

type Preference struct {
	Language string
	Theme    string
}

func NewPreference() *Preference {
	return &Preference{
		Language: utils.GetAppContext().Config.DefaultLanguage,
		Theme:    utils.GetAppContext().Config.DefaultTheme,
	}
}

func GetPreference(r *http.Request) *Preference {
	language, _ := r.Context().Value(LanguageKey).(string)
	if language == "" {
		language = utils.GetAppContext().Config.DefaultLanguage
	}

	theme, _ := r.Context().Value(ThemeKey).(string)
	if theme == "" {
		theme = utils.GetAppContext().Config.DefaultTheme
	}

	return &Preference{
		Language: language,
		Theme:    theme,
	}
}

func SetPreference(r *http.Request, language, theme string) *http.Request {
	ctx := context.WithValue(r.Context(), LanguageKey, language)
	ctx = context.WithValue(ctx, ThemeKey, theme)
	return r.WithContext(ctx)
}

func (p *Preference) Next(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		language := utils.GetAppContext().Config.DefaultLanguage
		if cookie, err := r.Cookie("language"); err == nil && cookie.Value != "" {
			language = cookie.Value
		}

		theme := utils.GetAppContext().Config.DefaultTheme
		if cookie, err := r.Cookie("theme"); err == nil && cookie.Value != "" {
			theme = cookie.Value
		}

		updatedRequest := SetPreference(r, language, theme)
		next.ServeHTTP(w, updatedRequest)
	})
}
