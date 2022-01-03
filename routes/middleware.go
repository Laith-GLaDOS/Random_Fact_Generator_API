package routes

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(req *http.Request) {
	fmt.Println("Received ", req.Method, " request at ", req.URL.Path)
}
