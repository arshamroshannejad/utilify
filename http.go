package utilify

import "net/http"

func GetByCtx(r *http.Request, key string) string {
	return r.Context().Value(key).(string)
}
