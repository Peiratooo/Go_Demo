package jwtToken

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("Peirato.233")

func GenToken(data map[string]interface{}) (string, error) {

	claim := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iss": "Peirato",
		"uid": data["uid"],
	}
	openToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//fmt.Println(openToken, "t")
	token, err := openToken.SignedString(jwtKey)
	//fmt.Println(token)
	return token, err
}

func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		//fmt.Println("OK", claims)
		return claims, err
	}
	return nil, err
}
