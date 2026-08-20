package utils

// Miscellaneous constants used throughout the application.
const (
	ROOT             = "."
	SLASH            = "/"
	DEFAULT_PORT     = 8080
	DEFAULT_LANGUAGE = "en"
)

// Static file paths for serving HTML content.
const (
	STATIC            = "static"
	HTML              = "html"
	HOME_HTML_FILE    = "index.html"
	MEMBERS_HTML_FILE = "members.html"
	ABOUT_HTML_FILE   = "about.html"
	CONTACT_HTML_FILE = "contact.html"
	EVENTS_HTML_FILE  = "events.html"

	HOME_PATH    = ROOT + SLASH + STATIC + SLASH + HTML + SLASH + HOME_HTML_FILE
	MEMBERS_PATH = ROOT + SLASH + STATIC + SLASH + HTML + SLASH + MEMBERS_HTML_FILE
	ABOUT_PATH   = ROOT + SLASH + STATIC + SLASH + HTML + SLASH + ABOUT_HTML_FILE
	CONTACT_PATH = ROOT + SLASH + STATIC + SLASH + HTML + SLASH + CONTACT_HTML_FILE
	EVENTS_PATH  = ROOT + SLASH + STATIC + SLASH + HTML + SLASH + EVENTS_HTML_FILE
)
