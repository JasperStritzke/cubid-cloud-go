package security

import (
	"encoding/base64"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"github.com/JasperStritzke/cubid-cloud/util/random"
)

const (
	keyLength = 2048
	pathName  = "executor.cubidkey"
)

var secretKey string

func InitControllerKey() {
	randomKey, _ := random.GenerateRandomString(keyLength)

	if fileutil.WriteIfNotExists(
		pathName,
		base64.StdEncoding.EncodeToString([]byte(randomKey)),
	) {
		logger.Info("Successfully created secret-key. Please copy the executor.cubidkey file to all executors.")
	}
}

func LoadControllerKey() {
	key := fileutil.ReadString(pathName)
	bytes, err := base64.StdEncoding.DecodeString(key)

	if len(key) == 0 || err != nil {
		panic("Can't load secret-key. Please make sure to copy the executor.cubidkey file into the working directory of this service.")
	}

	secretKey = string(bytes)
	logger.Info("Successfully loaded secret-key.")
}

func GetKey() string {
	return secretKey
}

func IsKeyValid(hash string) bool {
	return GetKey() == hash
}
