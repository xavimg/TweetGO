package middlewares

import (
	"net/http"

	"github.com/xavimg/TweetGO/bd"
)

/*ChequeoBD es el middleware que me permite conocer el estado de la BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConn() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)

	}
}
