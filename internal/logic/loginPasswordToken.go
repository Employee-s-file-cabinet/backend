package logic

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Employee-s-file-cabinet/backend/internal/server/internal/api"
	"github.com/go-eden/slf4go"
	"github.com/o1egl/paseto"
	"github.com/o1egl/paseto/timeutil"
	"golang.org/x/crypto/bcrypt"

	MyStorage "../../internal/db"
)

func LogicLoginPasswordToken(r *http.Request, auth api.Auth) (api.Token, error) {

	var PASETOToken api.Token

	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PG_PASSWORD")
	dbHost := os.Getenv("APP_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbName := os.Getenv("PG_DB")

	DBDSN := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	logger := slf4go.Get("example")

	db, err := MyStorage.OpenDB(DBDSN)
	if err != nil {
		logger.Error(err)
		return http.StatusBadRequest, PASETOToken
	}
	defer db.Close()

	// Выполнение запроса к базе данных
	ctx := r.Context()

	row := db.QueryRowContext(ctx, "SELECT password_hash FROM authorizations WHERE username = $1", auth.Login)

	var userPasswordHash string

	err = row.Scan(&userPasswordHash)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	err = CheckPassword(auth.Password, userPasswordHash)

	if err != nil {
		logger.Error(err)
		return "", err
	}

	// Создаем PASETO токен
	expirationTime := time.Now().Add(8 * time.Hour)

	v2 := paseto.NewV2()
	PASETOToken.PasetoToken, err = v2.Encrypt(secretKey, api.PasetoClaims{
		Username: auth.Username,
		Password: userPasswordHash,
		Expiration: &timeutil.Time{
			Time: expirationTime,
		},
	})
	if err != nil {
		logger.Error(err)
		return "", err
	}
	PASETOToken.ExpirationTime = int(8 * time.Hour)
	return PASETOToken, nil
}

// Функция сверки паролей.
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// HashPassword хеширует пароль.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
