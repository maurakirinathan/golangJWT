package main

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


func main() {

	initKeys()

	StartServer()
}




