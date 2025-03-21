package utilities

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "secretKey"

func GenerateToken(userID int64, emailAddress string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":       userID,
			"emailAddress": emailAddress,
			"exp":          time.Now().Add(time.Hour * 2).Unix(),
		},
	)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected-signing-method")
			}
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return 0, errors.New("could-not-parse-token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid-token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid-token-claims")
	}

	userID := int64(claims["userID"].(float64))
	// emailAddress := claims["emailAddress"].(string)

	return userID, nil
}
