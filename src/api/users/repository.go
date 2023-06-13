package users

import (
	"context"
	"fmt"
	"time"

	db "onlineshopgo/database"
	res "onlineshopgo/src/api/users/schemas"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Tokens struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func create_(customer *res.Create) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.PASSWORD), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err.Error())
	}

	db.DB.Exec(
		context.Background(),
		`INSERT INTO customers ("name", "password", "email", "phone_num") VALUES ($1,$2,$3,$4)`,
		customer.NAME, hashedPassword, customer.EMAIL, customer.PHONE_NUM,
	)
}

func login_(customer *res.Update) (Tokens, map[string]any) {

	err := db.DB.QueryRow(
		context.Background(),
		`
		SELECT
		id,
		name,
		password,
		email,
		phone_num
		FROM customers WHERE name =$1
		`, customer.NAME,
	).
		Scan(
			&customer.ID,
			&customer.NAME,
			&customer.PASSWORD,
			&customer.EMAIL,
			&customer.PHONE_NUM,
		)

	tokenString, err := generateToken(customer.ID, customer.NAME, customer.EMAIL)
	if err != nil {
		return Tokens{}, gin.H{"error": "Error generating token"}
	}

	return tokenString, nil
}

func generateToken(userID int, userName string, userEmail string) (Tokens, error) {
	var tokens Tokens
	var err error
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"username":  userName,
		"user_role": userEmail,
		"exp":       time.Now().Add(time.Hour * 24 * 1).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"username":  userName,
		"user_role": userEmail,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokens.AccessToken, err = accessToken.SignedString([]byte("secret-key"))
	if err != nil {
		return Tokens{}, err
	}

	tokens.RefreshToken, err = refreshToken.SignedString([]byte("secret-key"))
	if err != nil {
		return Tokens{}, err
	}
	fmt.Println(tokens)
	return tokens, nil
}

func RefreshToken(tokenString string) (Tokens, map[string]any) {
	var tokens Tokens

	if tokenString == "" {
		return Tokens{}, gin.H{"error": "No authorization header provided"}
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return Tokens{}, gin.H{"error": "Invalid token signature"}
		} else {
			return Tokens{}, gin.H{"error": "Invalid token"}
		}

	}

	if !token.Valid {
		return Tokens{}, gin.H{"error": "Invalid token"}
	}

	claims["exp"] = time.Now().Add(time.Hour * 24 * 1).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokens.AccessToken, err = accessToken.SignedString([]byte("secret-key"))

	if err != nil {
		return Tokens{}, gin.H{"error": err}
	}

	return tokens, nil
}
