package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

//SERVER ENTRY POINT
func StartServer() {
	// Non-Protected Endpoint(s)
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)

	// Protected Endpoints
	//To access the protected resource at_ /resource_, make sure to add the following header to the request:
	//Key: Authorization
	//Value:Bearer <space> token
	r.Handle("/resource", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))

	log.Println("Now listening...")
	http.ListenAndServe(":"+ config.hostPort, nil)
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
