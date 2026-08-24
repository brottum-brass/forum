package router

import (
	"net/http"

	v1 "github.com/brottum-brass/forum/src/api/v1"
	"github.com/brottum-brass/forum/src/utils"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	appCtx := utils.GetAppContext()
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", v1.HomeHandler())
	mux.Handle("GET /events", v1.HandleEvents())
	mux.Handle("GET /about", v1.HandleAbout())
	mux.Handle("GET /members", v1.HandleMembers())
	mux.Handle("GET /contact", v1.HandleContact())

	mux.Handle("GET /static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir(appCtx.Config.CSS))))
	mux.Handle("GET /static/icons/", http.StripPrefix("/static/icons/", http.FileServer(http.Dir(appCtx.Config.Icons))))
	mux.Handle("GET /static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir(appCtx.Config.JS))))
	mux.Handle("GET /static/locales/", http.StripPrefix("/static/locales/", http.FileServer(http.Dir(appCtx.Config.Locales))))
	mux.Handle("GET /static/themes/", http.StripPrefix("/static/themes/", http.FileServer(http.Dir(appCtx.Config.Themes))))

	mux.Handle("GET /toggle-theme", v1.ThemeHandler())
	mux.Handle("GET /change-language", v1.LanguageHandler())

	mux.Handle("GET /", v1.HandleNotFound())

	return &Router{mux: mux}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}
