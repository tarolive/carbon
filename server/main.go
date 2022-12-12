package main

import (
	"net/http"
)

func main() {

	handleStaticFiles()
}

func handleStaticFiles() {

	var (
		staticFilesPath    = "./web/static/"
		staticFilesPattern = "/static/"
	)

	http.Handle(staticFilesPattern, http.StripPrefix(staticFilesPattern, http.FileServer(http.Dir(staticFilesPath))))
}
