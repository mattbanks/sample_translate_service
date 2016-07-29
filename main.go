package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"log"
	"time"
	"github.com/emicklei/go-restful"
)

type Message struct {
	Id 			int 		`json:"id"`
	Language 	string 		`json:"language"`
    Content 	string 		`json:"content"`
    Translated 	time.Time 	`json:"translated_on"`
}

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(home))
	ws.Route(ws.POST("/translate").Consumes("application/json").To(translate))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

func home(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "simple translation service")
}

func translate(req *restful.Request, resp *restful.Response) {
	message := new(Message)

	if err := req.ReadEntity(message); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	message.Translated = time.Now()

	resp.WriteAsJson(message)
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
