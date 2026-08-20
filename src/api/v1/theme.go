package v1

import (
	"net/http"
	"time"

	"github.com/brottum-brass/forum/src/api/middleware"
	"github.com/brottum-brass/forum/src/api/v1/html"
	"github.com/brottum-brass/forum/src/utils"
)

func ThemeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentPrefs := middleware.GetPreference(r)

		newTheme := "darkmode"
		if currentPrefs.Theme == "darkmode" {
			newTheme = "lightmode"
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "theme",
			Value:    newTheme,
			Path:     "/",
			Expires:  time.Now().Add(365 * 24 * time.Hour),
			SameSite: http.SameSiteLaxMode,
		})

		updatedRequest := middleware.SetPreference(r, currentPrefs.Language, newTheme)

		languageContent := utils.GetAppContext().Language.L(updatedRequest.Context(), currentPrefs.Language)
		themeMode := utils.GetAppContext().Theme.T(updatedRequest.Context(), newTheme)

		if r.Header.Get("HX-Request") == "true" {
			html.Body(languageContent, themeMode, html.HomeContent(languageContent, themeMode)).Render(updatedRequest.Context(), w)
			return
		}

		html.Document(languageContent, themeMode, html.HomeContent(languageContent, themeMode)).Render(updatedRequest.Context(), w)
	}
}
