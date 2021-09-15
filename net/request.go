package net // import "github.com/Buccaneer69/wc-api-go/net"

import "net/http"

// RequestEnricher adds Basic Authentication settings in Request in case of Basic Authentication
type RequestEnricher interface {
	EnrichRequest(r *http.Request, URL string)
}
