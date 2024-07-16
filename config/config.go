// (c) Jisin0
//
// config/numerals.go creates user configured numbers.

package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	DBChannel   = int64Environ("DB_CHANNEL")                             // id of database channel
	Admins      = int64ListEnviron("ADMINS")                             // admins list
	AllowPublic = os.Getenv("ALLOW_PUBLIC") == "true" || len(Admins) < 1 // indicates wether anyone can use the bot
)

// int64Environ gets a environment variable and attempts to cast it into an int64.
func int64Environ(envVar string) int64 {
	s := os.Getenv(envVar)
	if s == "" {
		return 0
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("config.int64environ: failed to parse %s: %v\n", envVar, err)
	}

	return i
}

// int64ListEnviron returns a slice of int64 for an environment variable.
func int64ListEnviron(envVar string) (result []int64) {
	s := os.Getenv(envVar)
	if s == "" {
		return result
	}

	for _, str := range strings.Split(s, " ") {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Printf("config.int64listenviron: failed to parse %s: %v\n", envVar, err)
		}

		result = append(result, i)
	}

	return result
}