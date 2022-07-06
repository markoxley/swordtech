package swordtech

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markoxley/swordtech/views"
)

func AddFunc(n string, f interface{}) {
	views.AddFunc(n, f)
}

// RunApp starts the server
func RunApp(r func(*mux.Router)) {
	configuration, err := loadConfiguration()
	if err != nil {
		panic(err)
	}
	Log(nil, nil, fmt.Sprintf("%s configured", configuration.ApplicationName), 0)

	configureSession(configuration.ApplicationName, configuration.Secret)
	Log(nil, nil, "Session store created", 0)

	router := mux.NewRouter()
	router.Use(middlewareHandler)

	views.LoadTemplates()
	Log(nil, nil, "Templates loaded", 0)

	r(router)
	Log(nil, nil, "Page routes configured", 0)

	registerAssets(router)
	Log(nil, nil, "Asset routes configured", 0)

	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()
	Log(nil, nil, "Error routes configured", 0)

	// Create server
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%v", configuration.Server.Port),
	}
	Log(nil, nil, "Server created", 0)

	executeStartup()

	go runProcesses()
	Log(nil, nil, "Background runner started", 0)

	Log(nil, nil, fmt.Sprintf("%s listening on port %d", configuration.ApplicationName, configuration.Server.Port), 0)
	if configuration.Server.UseSSL {
		if err := server.ListenAndServeTLS(configuration.Server.SSLCert, configuration.Server.SSLKey); err != nil {
			Log(nil, nil, fmt.Sprintf("Server failed to start : %v", err), -1)
		}
	} else {
		if err := server.ListenAndServe(); err != nil {
			Log(nil, nil, fmt.Sprintf("Server failed to start : %v", err), -1)
		}
	}
}
