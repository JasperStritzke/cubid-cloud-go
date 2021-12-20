package console

import (
	"fmt"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
)

func AskFor(question, prompt string, v interface{}) {
	logger.Info(question)
	fmt.Print(prompt)

	_, err := fmt.Scan(v)

	if err != nil {
		return
	}
}
