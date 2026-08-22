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
		preference := middleware.GetPreference(r)
		appCtx := utils.GetAppContext()

		newTheme := "darkmode"
		if preference.Theme == "darkmode" {
			newTheme = "lightmode"
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "theme",
			Value:    newTheme,
			Path:     "/",
			Expires:  time.Now().Add(365 * 24 * time.Hour),
			SameSite: http.SameSiteLaxMode,
		})

		updatedRequest := middleware.SetPreference(r, preference.Language, newTheme)

		languageContent := appCtx.Language.L(updatedRequest.Context(), preference.Language)
		themeMode := appCtx.Theme.T(updatedRequest.Context(), newTheme)

		pageContent := renderCurrentPageContent(r, languageContent, themeMode)

		if updatedRequest.Header.Get("HX-Request") == "true" {
			html.Body(languageContent, themeMode, pageContent).Render(updatedRequest.Context(), w)
			return
		}

		html.Document(languageContent, themeMode, pageContent).Render(updatedRequest.Context(), w)
	}
}
