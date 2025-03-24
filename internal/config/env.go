package config

import "strings"

func EnvVar(flag string) string {
	return "AUTHONCHAIN_" + strings.ToUpper(strings.ReplaceAll(flag, "-", "_"))
}
