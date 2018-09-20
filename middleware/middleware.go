package middleware

import (
	"context"
	"net/http"

	"github.com/line/clova-cek-sdk-go/cek"
)

const (
	CtxKey = "clova"
)

type Clova struct {
	ce *cek.Extension
}

func New(extensionID string, options ...cek.ExtensionOption) (clova *Clova) {
	return &Clova{ce: cek.NewExtension(extensionID, options...)}
}

func (clova *Clova) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxKey, clova.ce)))
	})
}
