package swordtech

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	// Static folders
	assetFolder = "assets/"
	imgFolder   = assetFolder + "img/"
	jsFolder    = assetFolder + "js/"
	cssFolder   = assetFolder + "css/"
	fntFolder   = assetFolder + "fnt/"
)

// Serve a static file, ensuring that we do not display the folder
func assetServ(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func registerAssets(r *mux.Router) {
	// Set paths and routes for static assets
	imgfs := http.FileServer(http.Dir(imgFolder))
	r.PathPrefix("/img/").Handler(assetServ(http.StripPrefix("/img/", imgfs)))

	cssfs := http.FileServer(http.Dir(cssFolder))
	r.PathPrefix("/css/").Handler(assetServ(http.StripPrefix("/css/", cssfs)))

	jsfs := http.FileServer(http.Dir(jsFolder))
	r.PathPrefix("/js/").Handler(assetServ(http.StripPrefix("/js/", jsfs)))

	fntfs := http.FileServer(http.Dir(fntFolder))
	r.PathPrefix("/fnt/").Handler(assetServ(http.StripPrefix("/fnt/", fntfs)))

}
