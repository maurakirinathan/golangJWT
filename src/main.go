package src

//  Generate RSA signing files via shell (adjust as needed):
//
//  $ openssl genrsa -out app..keys 1024
//  $ openssl .keys -in app..keys -pubout > app..keys.pub
//
// Code borrowed and modified from the following sources:
// https://www.youtube.com/watch?v=dgJFeqeXVKw
// https://goo.gl/ofVjK4
// https://github.com/dgrijalva/jwt-go
//

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"crypto/rsa"

	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	// For simplicity these files are in the same folder as the app binary.
	privKeyPath = "/home/mauran/Desktop/rsakeys/app..keys"
	pubKeyPath  = "/home/mauran/Desktop/rsakeys/app..keys.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initKeys() {
   signBytes, err := ioutil.ReadFile(config.idRsa)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(config.idRsaPub)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

//STRUCT DEFINITIONS
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

//SERVER ENTRY POINT
func StartServer() {
	// Non-Protected Endpoint(s)
	http.HandleFunc("/login", LoginHandler)

	// Protected Endpoints
	//To access the protected resource at_ /resource_, make sure to add the following header to the request:
	//Key: Authorization
	//Value:Bearer <space> token
	http.Handle("/resource", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))

	log.Println("Now listening...")
	http.ListenAndServe(":"+ config.hostPort, nil)
}

func main() {
	initKeys()

	StartServer()
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user UserCredentials

	//decode request into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	//validate user credentials
	if (strings.ToLower(user.Username) != "mauran") || (user.Password != "p@ssword") {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	//create a .keys 256 token(signer)
	token := jwt.New(jwt.SigningMethodRS256)

	//set claims
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{user.Username, "Member"}
	token.Claims = claims

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)
	}

	//create a token instance using the token string
	response := Token{tokenString}
	JsonResponse(response, w)
}

//AUTH TOKEN VALIDATION
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//validate token
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}

//HELPER FUNCTIONS
func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
