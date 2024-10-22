package utils

import (
	"os"
)

const DEFAULT_FC_HOST = "https://api.fnctlst.io"
const FUSIONCATALYST_HOST_ENV_VAR = "FUSIONCATALYST_HOST"
const FUSIONCATALYST_API_KEY_ENV_VAR = "FUSIONCATALYST_API_KEY"

func GetFCHost() string {
	if value, exists := os.LookupEnv(FUSIONCATALYST_HOST_ENV_VAR); exists {
		return value
	}
	return DEFAULT_FC_HOST
}

func GetFCAPIKey() (string, bool) {
	if value, exists := os.LookupEnv(FUSIONCATALYST_API_KEY_ENV_VAR); exists {
		return value, true
	}
	return "", false
}
