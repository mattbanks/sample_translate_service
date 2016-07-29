package main

import (
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
	ws.Route(ws.POST("/translate").Consumes("application/json").To(translate))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func translate(req *restful.Request, resp *restful.Response,) {
	message := new(Message)

	if err := req.ReadEntity(message); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	message.Translated = time.Now()

	resp.WriteAsJson(message)
}
