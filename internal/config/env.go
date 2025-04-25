package config

import "strings"

func EnvVar(flag string) string {
	return "AUTHONCHAIN_" + strings.ToUpper(strings.ReplaceAll(flag, "-", "_"))
}

func EnvVars(flag string) []string {
	return []string{EnvVar(flag)}
}
