package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
)

var mySigningKey []byte = []byte("qwerty")

const apiKey string = "key"

// TokenJWTHandler Genertate JWT Token
func TokenJWTHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "apllication/json")

	authenAPI := false
	for name, values := range r.Header {
		//fmt.Printf("%v : %#v\n", name, values)
		if strings.ToLower(name) == "api-key" {
			if strings.ToLower(values[0]) == apiKey {
				authenAPI = true
				break
			}
		}
	}

	if !authenAPI {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := Token()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

// Token generate
func Token() (string, error) {
	unix := time.Now().Add(5 * time.Minute).Unix()

	claims := &jwt.StandardClaims{
		ExpiresAt: unix,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

// AuthKey api key header
func AuthKey(reqToken string) (bool, error) {

	if reqToken == "" {
		return false, nil
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: 0,
	}

	_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return false, fmt.Errorf("token expired")
	}

	return true, nil
}

func BearerAuthKey(c echo.Context) (bool, error) {
	reqToken := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = ""

	if len(splitToken) > 1 {
		reqToken = splitToken[1]
	}

	if reqToken == "" {
		return false, c.JSON(http.StatusUnauthorized, map[string]string{
			"err": "JWT is required",
		})
	}

	if ok, err := AuthKey(reqToken); !ok {
		if err != nil {
			return false, c.JSON(http.StatusUnauthorized, map[string]string{
				"err": "invalid token",
			})
		}

		return false, c.JSON(http.StatusUnauthorized, map[string]string{
			"err": "token expired",
		})
	}

	return true, nil
}
