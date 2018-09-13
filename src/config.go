package src

import "os"

type Config struct {
	switchName string
	switchHost string
	switchPort string
	senzieMode string
	senzieName string
	dotKeys    string
	idRsa      string
	idRsaPub   string
}

var config = Config{
	idRsa:      getEnv("ID_RSA", ".keys/app..keys"),
	idRsaPub:   getEnv("ID_RSA_PUB", "..keys/app..keys.pub"),
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
