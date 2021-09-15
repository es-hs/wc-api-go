package net // import "github.com/Buccaneer69/wc-api-go/net"

import (
	"github.com/Buccaneer69/wc-api-go/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
