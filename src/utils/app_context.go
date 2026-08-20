package utils

import (
	"github.com/brottum-brass/forum/src/utils/language"
	"github.com/brottum-brass/forum/src/utils/theme"
)

type AppContext struct {
	Config   *Config
	Language language.Language
	Theme    theme.Theme
}

var appCtx *AppContext

func initAppContext() {
	if appCtx == nil {
		config := NewConfig()

		language := language.NewJSONLanguage(config.Locales, config.DefaultLanguage)
		if err := language.LoadLanguages(); err != nil {
			panic(err)
		}

		theme := theme.NewJSONTheme(config.Themes, config.DefaultTheme)
		if err := theme.LoadThemes(); err != nil {
			panic(err)
		}

		appCtx = &AppContext{
			Config:   config,
			Language: language,
			Theme:    theme,
		}
	}
}

func GetAppContext() *AppContext {
	if appCtx == nil {
		initAppContext()
	}

	return appCtx
}
