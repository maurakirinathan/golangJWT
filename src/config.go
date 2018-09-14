package main

import "os"

type Config struct {
	idRsa      string
	idRsaPub   string
	hostPort   string
}

var config = Config{
	idRsa:      getEnv("ID_RSA", ".keys/app.rsa"),
	idRsaPub:   getEnv("ID_RSA_PUB", ".keys/app.rsa.pub"),
	hostPort:   getEnv("HPORT", "8002"),
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
