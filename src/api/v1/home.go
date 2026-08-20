package v1

import (
	"net/http"

	"github.com/brottum-brass/forum/src/api/middleware"
	"github.com/brottum-brass/forum/src/api/v1/html"
	"github.com/brottum-brass/forum/src/utils"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preference := middleware.GetPreference(r)

		languageContent := utils.GetAppContext().Language.L(r.Context(), preference.Language)
		themeMode := utils.GetAppContext().Theme.T(r.Context(), preference.Theme)

		html.Document(languageContent, themeMode, html.HomeContent(languageContent, themeMode)).Render(r.Context(), w)
	}
}
