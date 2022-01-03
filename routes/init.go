package routes

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

func Index(res http.ResponseWriter, req *http.Request) {
	LoggerMiddleware(req)
	if req.Method == "GET" {
		res.Header().Set("Content-Type", "text/html")
		pageContents, _ := ioutil.ReadFile("pages/index.html")
		res.Write(pageContents)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("{\"error\":\"Method not allowed/supported\"}"))
	}
}

func Docs(res http.ResponseWriter, req *http.Request) {
	LoggerMiddleware(req)
	if req.Method == "GET" {
		res.Header().Set("Content-Type", "text/html")
		pageContents, _ := ioutil.ReadFile("pages/docs.html")
		res.Write(pageContents)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("{\"error\":\"Method not allowed/supported\"}"))
	}
}

func API(res http.ResponseWriter, req *http.Request) {
	LoggerMiddleware(req)
	if req.Method == "GET" {
		res.Header().Set("Content-Type", "application/json")
		facts_unsliced, _ := ioutil.ReadFile("facts")
		facts := strings.Split(string(facts_unsliced), "TERMINATOR")
		res.Write([]byte("{\"fact\":\"" + facts[rand.Intn(len(facts))] + "\"}"))
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("{\"error\":\"Method not allowed/supported\"}"))
	}
}
