package secure

import (
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

//MyClaims constructs the Claims object for JWT
type MyClaims struct {
	jwt.StandardClaims
	SessionID string
}

var (
	key []byte
)

//getEnv retrieves key for JWT signing
func getEnv(k string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
	return (os.Getenv(k))
}

//GenerateJWT generates JWT token for user sessions
func GenerateJWT(c *MyClaims) (string, error) {
	// key = []byte(getEnv("JWT_KEY"))
	key = []byte("thisismyjwtkey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error in creating token")
	}
	return signedToken, nil
}

//ParseToken receives token and validates
func ParseToken(signedToken string) (*MyClaims, error) {
	claims := &MyClaims{}
	t, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() { //the token here is not yet verified
			return nil, fmt.Errorf("Invalid signing algorithm")
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error in parseToken while parsing token")
	}
	if !t.Valid {
		return nil, fmt.Errorf("Error in parseToken, token is no longer valid")
	}
	return t.Claims.(*MyClaims), err
}
