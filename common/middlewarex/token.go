package middlewarex

import (
	"net/http"
)

type SetTokenMiddleware struct{}

func NewSetTokenMiddleware() *SetTokenMiddleware {
	return &SetTokenMiddleware{}
}
func (m *SetTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//token := r.Header.Get("Authorization")
		ctx := r.Context()
		//ctx = context.WithValue(ctx, "token", token)
		next(w, r.WithContext(ctx))
	}
}
