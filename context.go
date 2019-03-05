package httprouter

import (
	"net/http"
	"context"
)

type paramsKey struct{}

// ParamsKey is the request context key under which URL params are stored.
//
// This is only present from go 1.7.
var ParamsKey = paramsKey{}

// ParamsFromContext pulls the URL parameters from a request context,
// or returns nil if none are present.
//
// This is only present from go 1.7.
func ParamsFromContext(ctx context.Context) Params {
	p, _ := ctx.Value(ParamsKey).(Params)
	return p
}


// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

func wrapWithContext(handler http.Handler, ps Params) http.Handler {
	return http.HandlerFunc(func (res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		ctx = context.WithValue(ctx, ParamsKey, ps)
		req = req.WithContext(ctx)
		handler.ServeHTTP(res, req)
	})
}
