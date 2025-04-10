package services

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	types "workout_tracker"
	db "workout_tracker/db/queries"
	"workout_tracker/global"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(email, password string) error {
	err := validateUserForm(email, password)
	if err != nil {
		return err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	err = db.InsertUser(email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func Login(email, password string) (string, error) {
	err := validateUserForm(email, password)
	if err != nil {
		return "", err
	}

	user, err := db.GetUser(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := generateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT(email string) (string, error) {
	claims := types.UserClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "workout tracker",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(global.CONFIG.JwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ParseJWT(tokenStr string) (*types.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &types.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.CONFIG.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*types.UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func validateUserForm(email, password string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
