package v1

import (
	"net/http"

	"github.com/brottum-brass/forum/src/api/middleware"
	"github.com/brottum-brass/forum/src/api/v1/html"
	"github.com/brottum-brass/forum/src/utils"
)

func HandleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preference := middleware.GetPreference(r)
		appCtx := utils.GetAppContext()

		languageContent := appCtx.Language.L(r.Context(), preference.Language)
		theme := appCtx.Theme.T(r.Context(), preference.Theme)

		w.WriteHeader(http.StatusNotFound)

		if r.Header.Get("HX-Request") == "true" {
			html.NotFoundView(languageContent, theme).Render(r.Context(), w)
			return
		}

		html.Document(languageContent, theme, html.NotFoundView(languageContent, theme)).Render(r.Context(), w)
	}
}
