package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/tiwtterGo/models"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		// rutnina que checa contra la BD
	}
	if !tkn.Valid {
		return &claims, false, string(""), errors.New("token invalido")
	}

	return &claims, false, string(""), err
}
