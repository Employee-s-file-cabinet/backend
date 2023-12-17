package logic

import (
	"net/http"
	"os"
	"time"

	"github.com/go-eden/slf4go"
	"github.com/o1egl/paseto"
)

var secretKey = []byte(os.Getenv("PASETO_SECRET_KEY"))

// PasetoClaims - пример пользовательской структуры для хранения данных в токене.
type PasetoClaims struct {
	Username     string    `json:"username"`
	PasswordHash string    `json:"password"`
	Expiration   time.Time `json:"exp"`
	// Другие поля, если необходимо
}

// verifyPasetoToken - функция проверки PASETO токена.
func verifyPasetoToken(secretKey []byte, token string) error {
	v2 := paseto.NewV2()
	var claims PasetoClaims
	err := v2.Decrypt(token, secretKey, &claims, nil)
	if err != nil {
		return err
	}

	return nil
}

func LogicAuthentication(r *http.Request) error {
	// Пример проверки токена
	authHeader := r.Header.Get("AUTHORIZATION")

	logger := slf4go.Get("example")

	err := verifyPasetoToken(secretKey, authHeader)
	if err != nil {
		loger.Error(err)
		return err
	}

	return nil
}
