package main

import (
	"net/http"
)

func main() {

	handleStaticFiles()
	handleRobotsFile()

	listenAndServe()
}

func handleStaticFiles() {

	var (
		staticFilesPath    = "./web/static/"
		staticFilesPattern = "/static/"
	)

	http.Handle(staticFilesPattern, http.StripPrefix(staticFilesPattern, http.FileServer(http.Dir(staticFilesPath))))
}

func handleRobotsFile() {

	var (
		robotsFilePath    = "./web/static/configuration/robots.txt"
		robotsFilePattern = "/robots.txt"
	)

	http.HandleFunc(robotsFilePattern, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, robotsFilePath)
	})
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
