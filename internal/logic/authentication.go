package logic

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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
func verifyPasetoToken(secretKey []byte, token string) (*PasetoClaims, error) {
	v2 := paseto.NewV2()
	var claims PasetoClaims
	err := v2.Decrypt(token, secretKey, &claims, nil)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

type RequestBodyType struct {
	PasetoTokenString string `json:"posetotokenstring"`
	Username          string `json:"username"`
	Password          string `json:"password"`
}

func LogicAuthentication(r *http.Request) int {
	// Пример проверки токена
	rbstruct := RequestBodyType{}
	byteToken, err := io.ReadAll(r.Body)
	if err := json.Unmarshal(byteToken, &rbstruct); err != nil {
		log.Println(err)
	}
	claims, err := verifyPasetoToken(secretKey, rbstruct.PasetoTokenString)
	if err != nil {
		fmt.Println("Ошибка при проверке токена:", err)
		return 400
	}
	// Вывод информации из токена
	tmpHash, err := HashPassword(rbstruct.Password)
	if err != nil {
		fmt.Println("Ошибка при проверке токена:", err)
		return 400
	}
	if (claims.Username != rbstruct.Username) || (claims.PasswordHash != tmpHash) {
		fmt.Println("Ошибка при проверке токена:", err)
		return 400
	}
	return 200
}
