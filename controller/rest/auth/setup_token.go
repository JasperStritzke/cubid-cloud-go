package auth

import (
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"github.com/JasperStritzke/cubid-cloud/util/random"
	"io/ioutil"
	"os"
)

const setupTokenPath = "setupkey"

var cachedSetupToken string

func GenerateSetupToken() {
	if fileutil.ExistsFile(setupTokenPath) {
		DeleteSetupToken()
	}

	_ = fileutil.CreateIfNotExists(setupTokenPath)

	cachedSetupToken, _ = random.GenerateRandomString(64)
	err := ioutil.WriteFile(setupTokenPath, []byte(cachedSetupToken), os.ModePerm)

	if err != nil {
		logger.Error("Unable to write setup-token")
		os.Exit(0)
	}

}

func VerifySetupToken(providedToken string) bool {
	return cachedSetupToken == providedToken
}

func DeleteSetupToken() {
	_ = os.Remove(setupTokenPath)

	cachedSetupToken = ""
}
