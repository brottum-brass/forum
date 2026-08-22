package v1

import (
	"net/http"
	"time"

	"github.com/brottum-brass/forum/src/api/middleware"
	"github.com/brottum-brass/forum/src/api/v1/html"
	"github.com/brottum-brass/forum/src/utils"
)

func LanguageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preference := middleware.GetPreference(r)
		appCtx := utils.GetAppContext()

		newLanguage := r.URL.Query().Get("language")
		if newLanguage == "" {
			newLanguage = "en"
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "language",
			Value:    newLanguage,
			Path:     "/",
			Expires:  time.Now().Add(365 * 24 * time.Hour),
			SameSite: http.SameSiteLaxMode,
		})

		updatedRequest := middleware.SetPreference(r, newLanguage, preference.Theme)

		languageContent := appCtx.Language.L(updatedRequest.Context(), newLanguage)
		themeMode := appCtx.Theme.T(updatedRequest.Context(), preference.Theme)

		pageContent := renderCurrentPageContent(r, languageContent, themeMode)

		if updatedRequest.Header.Get("HX-Request") == "true" {
			html.Body(languageContent, themeMode, pageContent).Render(updatedRequest.Context(), w)
			return
		}

		html.Document(languageContent, themeMode, pageContent).Render(updatedRequest.Context(), w)
	}
}
