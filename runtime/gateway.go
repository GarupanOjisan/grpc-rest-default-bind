package runtime

import (
	"net/http"

	"google.golang.org/grpc"
)

type Gateway interface {
	Methods() []Method
}

type Method interface {
	SetConn(conn *grpc.ClientConn)
	Path() string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
