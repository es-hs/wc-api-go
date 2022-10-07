package net // import "github.com/gempages/wc-api-go/net"

import (
	"github.com/gempages/wc-api-go/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
