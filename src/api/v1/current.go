package v1

import (
	"net/http"
	"net/url"

	"github.com/a-h/templ"
	"github.com/brottum-brass/forum/src/api/v1/html"
	"github.com/brottum-brass/forum/src/utils/language"
	"github.com/brottum-brass/forum/src/utils/theme"
)

func renderCurrentPageContent(r *http.Request, languageContent language.Content, theme theme.Mode) templ.Component {
	currentPath := "/"
	if hxURL := r.Header.Get("HX-Current-URL"); hxURL != "" {
		if parsed, err := url.Parse(hxURL); err == nil {
			currentPath = parsed.Path
		}
	} else if referer := r.Header.Get("Referer"); referer != "" {
		if parsed, err := url.Parse(referer); err == nil {
			currentPath = parsed.Path
		}
	}

	switch currentPath {
	case "/events":
		return html.EventsView(languageContent.Events.EventItems, languageContent, theme)
	case "/about":
		return html.AboutView(languageContent, theme)
	case "/members":
		return html.MembersView(languageContent.Members.MemberItems, languageContent, theme)
	case "/contact":
		return html.ContactView(languageContent, theme)
	default:
		return html.HomeContent(languageContent, theme)
	}
}
