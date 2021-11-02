package services

import (
	"BackLight/models"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)
var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(user *models.User) (string, error) {

	claims := models.Claims{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: *user.Email,
		Mobile: user.Mobile,
		StandardClaims: jwt.StandardClaims{},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err)
	}

	return token, err
}

func CheckToken(stringToken string) (models.Claims, error) {

	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	fmt.Println(ok, claims, token.Valid)

	if ok && token.Valid {

		id, _ := ParseUint(fmt.Sprintf("%v", claims["ID"]))
		firstName := fmt.Sprintf("%v", claims["FirstName"])
		lastName := fmt.Sprintf("%v", claims["LastName"])
		email := fmt.Sprintf("%v", claims["Email"])
		Mobile := fmt.Sprintf("%v", claims["Mobile"])

		return models.Claims{
			ID: id,
			FirstName: firstName,
			LastName: lastName,
			Email: email,
			Mobile: Mobile,
		}, nil
	}

	return models.Claims{}, err
}
