package verify

import (
	"github.com/Vozhlak/golang-advanced/3-validation-api/pkg/utils"
)

func generateRandomHash() (string, error) {
	hash, err := utils.GenerateRandomString(32)
	if err != nil {
		return "", err
	}
	return hash, nil
}
