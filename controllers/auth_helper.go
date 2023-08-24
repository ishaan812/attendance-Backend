package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"service/database"
)

func LoginUser(Faculty *database.Faculty) (*database.Faculty, error) {
	var faculty database.Faculty
	err := dbconn.Where("s_api_d = ?", Faculty.SAPID).First(&faculty).Error
	if err != nil {
		fmt.Println("ERROR: sapid does not exist")
		return nil, err
	}

	if faculty.Password != Faculty.Password {
		fmt.Println("ERROR: Wrong Password Entered")
		return nil, errors.New("wrong password")
	}

	fmt.Println("INFO: ", Faculty.SAPID, " logged in")
	Faculty = &faculty
	return &faculty, nil
}

func RegisterUser(Faculty *database.Faculty) error {
	err := dbconn.Where("s_api_d = ?", Faculty.SAPID).First(&Faculty).Error
	if err != nil {
		dbconn.Create(&Faculty)
		fmt.Println("INFO: New Faculty ", Faculty.SAPID, " has been registered")
		return nil
	}
	fmt.Println("ERROR: Faculty ", Faculty.SAPID, " already exists")
	return errors.New("user already exists")
}

func LogoutUser(c *http.Cookie) error {
	fmt.Println("INFO: Logged out")

	return nil
}

// func CreateJWT(Faculty *database.Faculty) (*http.Cookie, error) {
// 	expirationTime := time.Now().Add(tokenValidityDuration)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sapid":      Faculty.SAPID,
// 		"name":       Faculty.Name,
// 		"email":      Faculty.Email,
// 		"expiresat":  expirationTime,
// 		"department": Faculty.Department,
// 	})
// 	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
// 	if err != nil {
// 		fmt.Println("ERROR: Error in JWT Key")
// 		return nil, err
// 	}
// 	JWTCookie := &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	}
// 	fmt.Println("INFO: JWT of ", Faculty.SAPID, " generated with expiration time ", expirationTime)
// 	return JWTCookie, nil
// }

// func ValidateJWT(c *http.Cookie, jwtKey string) (*Claims, error) {
// 	var claims Claims
// 	tknStr := c.Value
// 	tkn, err := jwt.ParseWithClaims(tknStr, &claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(jwtKey), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !tkn.Valid {
// 		fmt.Println("ERROR: JWT of ", claims.SAPID, " is invalid")
// 		return nil, err
// 	}
// 	fmt.Println("INFO: JWT of ", claims.SAPID, " validated")
// 	return &claims, nil
// }

// func RefreshJWT(claims *Claims, jwtKey string) (*http.Cookie, error) {
// 	expirationTime := time.Now().Add(tokenValidityDuration)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sapid":      claims.SAPID,
// 		"name":       claims.Name,
// 		"email":      claims.Email,
// 		"expiresat":  expirationTime,
// 		"department": claims.Department,
// 	})
// 	tokenString, err := token.SignedString([]byte(jwtKey))
// 	if err != nil {
// 		fmt.Println("ERROR: Error while creating JWT.")
// 		return nil, err
// 	}
// 	JWTCookie := &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	}
// 	fmt.Println("INFO: JWT of ", claims.SAPID, " refreshed to ", expirationTime)
// 	return JWTCookie, nil
// }
