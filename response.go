package mosquito


import (
	"net/http"
)


type Response struct {
	http.ResponseWriter
	Locals					map[string]interface{}
	//Render
}

type Handler interface {
	ServeHTTP(Request, Response)
}
