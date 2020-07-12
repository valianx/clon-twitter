package middlew

import (
	"github.com/valianx/clon-twitter/bd"
	"net/http"
)

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Coneccion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
