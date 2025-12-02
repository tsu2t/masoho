package handlers

import "net/http"

// UsersHandler はダミーのユーザ一覧を返すRESTエンドポイントの例です。
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		// 本来はDBなどから取得した結果を返す
		_, _ = w.Write([]byte(`[
  {"id":1,"name":"Alice"},
  {"id":2,"name":"Bob"}
]`))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(`{"error":"method not allowed"}`))
	}
}


