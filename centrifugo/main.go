package centrifugo

import (
	"github.com/centrifugal/centrifuge-go"
	"github.com/dgrijalva/jwt-go"
	"log"
)

const Channel = "test-channel"

func NewCentrifugoConnection(user string) *centrifuge.Client {
	client := centrifuge.NewJsonClient(
		"ws://localhost:8000/connection/websocket",
		centrifuge.Config{
			Token: connToken(user, 0),
		})
	return client
}

func connToken(user string, exp int64) string {
	claims := jwt.MapClaims{"sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("bbe7d157-a253-4094-9759-06a8236543f9"))
	if err != nil {
		log.Println(err)
	}
	return t
}
