package main

import (
	"net/http"
)

func main() {

	handleStaticFiles()
	listenAndServe()
}

func handleStaticFiles() {

	var (
		staticFilesPath    = "./web/static/"
		staticFilesPattern = "/static/"
	)

	http.Handle(staticFilesPattern, http.StripPrefix(staticFilesPattern, http.FileServer(http.Dir(staticFilesPath))))
}

func listenAndServe() {

	var (
		serverAddress              = ":8080"
		serverHandler http.Handler = nil
	)

	err := http.ListenAndServe(serverAddress, serverHandler)

	if err != nil {
		panic(err)
	}
}
