package net // import "github.com/thanks173/wc-api-go/net"

import (
	"github.com/thanks173/wc-api-go/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
