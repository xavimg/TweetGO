package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xavimg/TweetGO/models"
)

/*GeneroJWT geenra el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("MastersDelDesarrollo_grupoDeFacebook")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioWeb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err

	}
	return tokenStr, nil

}
