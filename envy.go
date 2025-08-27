package envy

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func String(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

func MustString(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("missing required environment variable: %q", key))
	}

	return value
}

func parseBool(key, value string) bool {
	var truthyValues = []string{"true", "1", "on", "yes"}
	if slices.Contains(truthyValues, value) {
		return true
	}

	var falsyValues = []string{"false", "0", "off"}
	if slices.Contains(falsyValues, value) {
		return false
	}

	panic(
		fmt.Errorf(
			"invalid bool environment variable %q: %q (must be one of: %v)",
			key,
			value,
			append(truthyValues, falsyValues...),
		),
	)
}

func Bool(key string) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return false
	}

	return parseBool(key, value)
}

func MustBool(key string) bool {
	return parseBool(key, MustString(key))
}

func Int64(key string, fallback int) int64 {
	fallbackString := strconv.FormatInt(int64(fallback), 10)
	i, err := strconv.ParseInt(String(key, fallbackString), 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid int environment variable %q: %w", key, err))
	}

	return i
}

func MustInt64(key string) int64 {
	i, err := strconv.ParseInt(MustString(key), 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid int environment variable %q: %w", key, err))
	}

	return i
}

func Int32(key string, fallback int32) int32 {
	fallbackString := strconv.FormatInt(int64(fallback), 10)
	i, err := strconv.ParseInt(String(key, fallbackString), 10, 32)
	if err != nil {
		panic(fmt.Errorf("invalid int environment variable %q: %w", key, err))
	}

	return int32(i)
}

func MustInt32(key string) int32 {
	i, err := strconv.ParseInt(MustString(key), 10, 32)
	if err != nil {
		panic(fmt.Errorf("invalid int environment variable %q: %w", key, err))
	}

	return int32(i)
}

type PortNumber int32

func parsePort(key string, number int32) PortNumber {
	if number < 0 {
		panic(fmt.Errorf("invalid port number %q %d < 0", key, number))
	}

	if maxPortNumber := int32(65535); number > maxPortNumber {
		panic(fmt.Errorf("invalid port number %q %d > %d", key, number, maxPortNumber))
	}

	return PortNumber(number)
}

func Port(key string, fallback int32) PortNumber {
	return parsePort(key, Int32(key, fallback))
}

func MustPort(key string) PortNumber {
	return parsePort(key, MustInt32(key))
}
