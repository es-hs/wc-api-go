package client // import "github.com/gempages/wc-api-go/client"

import (
	"github.com/gempages/wc-api-go/request"
	"net/http"
)

// Sender interface
type Sender interface {
	Send(req request.Request) (resp *http.Response, err error)
}
