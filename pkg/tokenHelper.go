package pkg

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/TD17/jade-palace/config"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

// type TempSignedDetails struct {
// 	Email string
// 	jwt.StandardClaims
// }

// user is the name of the collection
// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// Set the secret key in env section. This is for temporary purpose

// func TempGenerateAllTokens(email string) (signedToken string, signedRefreshToken string, err error) {
// 	claims := &SignedDetails{
// 		Email: email,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
// 		},
// 	}

// 	refreshClaims := &SignedDetails{
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
// 		},
// 	}

// 	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(config.GetKey("private_key"))
// 	if err != nil {
// 		log.Panic("Error parsing private key:", err)
// 		return "", "", err
// 	}

// 	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
// 	if err != nil {
// 		log.Panic(err)
// 		return
// 	}

// 	refreshedtoken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims).SignedString(privateKey)
// 	if err != nil {
// 		log.Panic(err)
// 		return
// 	}

// 	return token, refreshedtoken, err
// }

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	// Refresh tokens are the new tokens we get if our regular token has expired
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(config.GetKey("private_key"))
	if err != nil {
		log.Panic("Error parsing private key:", err)
		return "", "", err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		log.Panic(err)
		return
	}

	refreshedtoken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims).SignedString(privateKey)
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshedtoken, err
}
