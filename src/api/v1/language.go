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
		currentPrefs := middleware.GetPreference(r)

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

		updatedRequest := middleware.SetPreference(r, newLanguage, currentPrefs.Theme)

		languageContent := utils.GetAppContext().Language.L(updatedRequest.Context(), newLanguage)
		themeMode := utils.GetAppContext().Theme.T(updatedRequest.Context(), currentPrefs.Theme)

		if updatedRequest.Header.Get("HX-Request") == "true" {
			html.Body(languageContent, themeMode, html.HomeContent(languageContent, themeMode)).Render(updatedRequest.Context(), w)
			return
		}

		html.Document(languageContent, themeMode, html.HomeContent(languageContent, themeMode)).Render(updatedRequest.Context(), w)
	}
}
