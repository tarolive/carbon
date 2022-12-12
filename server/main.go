package main

import (
	"html/template"
	"net/http"
)

func main() {

	handleStaticFiles()
	handleRobotsFile()

	handleIndexPage()

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
		robotsFilePath    = "./web/static/configurations/robots.txt"
		robotsFilePattern = "/robots.txt"
	)

	http.HandleFunc(robotsFilePattern, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, robotsFilePath)
	})
}

func handleIndexPage() {

	var (
		baseLayoutPath   = "./web/templates/layouts/base.tmpl"
		indexPagePath    = "./web/templates/pages/index.tmpl"
		indexPagePattern = "/"
	)

	indexPageTemplate := template.Must(template.ParseFiles(baseLayoutPath, indexPagePath))

	http.HandleFunc(indexPagePattern, func(responseWriter http.ResponseWriter, request *http.Request) {

		err := indexPageTemplate.Execute(responseWriter, nil)

		if err != nil {
			panic(err)
		}
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
