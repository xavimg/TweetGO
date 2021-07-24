package routers

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xavimg/TweetGO/bd"
	"github.com/xavimg/TweetGO/models"
)

/* Email = Email */
var Email string

/*IDUsuario = IDUsuario */
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("MastersDelDesarrollo_grupoDeFacebook")
	claims := &models.Claim{}

	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
