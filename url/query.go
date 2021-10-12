package url // import "github.com/es-hs/wc-api-go/url"

import (
	"github.com/es-hs/wc-api-go/request"
	"net/url"
)

// QueryEnricher uses package auth to enrich existing query parameters with Authentication Based ones
type QueryEnricher interface {
	GetEnrichedQuery(url string, query url.Values, req request.Request) url.Values
}
