package logic

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Employee-s-file-cabinet/backend/internal/server/internal/api"
	"github.com/o1egl/paseto"
	"github.com/o1egl/paseto/timeutil"
	"golang.org/x/crypto/bcrypt"

	MyStorage "../../internal/database"
)

var secretKey = []byte(os.Getenv("PASETO_SECRET_KEY"))

type ResponceToken struct {
	PasetoToken    string
	ExpirationTime int64
}

func LogicAuthorization(r *http.Request, auth api.Auth) (int, ResponceToken) {

	var PASETOToken ResponceToken

	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PG_PASSWORD")
	dbHost := os.Getenv("APP_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbName := os.Getenv("PG_DB")

	DBDSN := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := MyStorage.OpenDB(DBDSN)
	if err != nil {
		log.Println(err)
		return 400, PASETOToken
	}
	defer db.Close()

	// Выполнение запроса к базе данных
	rows, err := db.Query("SELECT password_hash FROM authorizations WHERE username = $1", auth.Login)
	if err != nil {
		log.Println(err)
		return 400, PASETOToken
	}
	defer rows.Close()

	var userPasswordHash string

	for rows.Next() {
		err := rows.Scan(&userPasswordHash)
		if err != nil {
			log.Println(err)
			return 400, PASETOToken
		}
	}

	err = CheckPassword(auth.Password, userPasswordHash)

	if err != nil {
		log.Println(err)
		return 400, PASETOToken
	}

	// Создаем PASETO токен
	expirationTime := time.Now().Add(1 * time.Hour)

	v2 := paseto.NewV2()
	PASETOToken.PasetoToken, err = v2.Encrypt(secretKey, api.PasetoClaims{
		Username: auth.Username,
		Password: userPasswordHash,
		Expiration: &timeutil.Time{
			Time: expirationTime,
		},
	})
	if err != nil {
		log.Println(err)
		return 400, PASETOToken
	}
	PASETOToken.ExpirationTime = expirationTime.Unix()
	return 200, PASETOToken
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
