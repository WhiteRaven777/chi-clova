package clova

import (
	"net/http"

	mw "github.com/WhiteRaven777/chi-clova/middleware"
	"github.com/line/clova-cek-sdk-go/cek"
)

func Clova(r *http.Request) (clova *cek.Extension) {
	return r.Context().Value(mw.CtxKey).(*cek.Extension)
}
