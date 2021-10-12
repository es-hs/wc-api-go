package client // import "github.com/es-hs/wc-api-go/client"

import (
	"github.com/es-hs/wc-api-go/request"
	"net/http"
)

// Sender interface
type Sender interface {
	Send(req request.Request) (resp *http.Response, err error)
}
