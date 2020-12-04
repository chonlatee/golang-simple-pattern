package main

import "fmt"

type server interface {
	handleRequest(string, string) (int, string)
}

func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	appStatusMethod := "GET"

	httpCode, resp := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("%s %s %s %d\n", appStatusMethod, appStatusURL, resp, httpCode)
	httpCode, resp = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("%s %s %s %d\n", appStatusMethod, appStatusURL, resp, httpCode)
	httpCode, resp = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("%s %s %s %d\n", appStatusMethod, appStatusURL, resp, httpCode)

	fmt.Println()
	appStatusURL = "/app/status"
	appStatusMethod = "POST"
	httpCode, resp = nginxServer.handleRequest(appStatusURL, appStatusMethod)
	fmt.Printf("%s %s %s %d\n", appStatusMethod, appStatusURL, resp, httpCode)
}