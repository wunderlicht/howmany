package main

import (
	"os"
	"strconv"
)

// returns env's value when set otherwise fallback.
func getEnvOrDefaultString(key string, fallback string) string {
	val, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	return val
}

// returns env's value when set and well-formed otherwise fallback
func getEnvOrDefaultInt(key string, fallback int) int {
	val, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return i
}

// returns env's value when set and well-formed otherwise fallback
func getEnvOrDefaultFloat(key string, fallback float64) float64 {
	val, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fallback
	}
	return f
}

// returns env's value when set and well-formed otherwise fallback
func getEnvOrDefaultBool(key string, fallback bool) bool {
	val, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}
	return b
}
