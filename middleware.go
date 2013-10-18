package mosquito


type Middleware interface {
	ServeHTTP(w Response, r *Request, next func())
}


