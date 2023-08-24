package common

import (
	"os"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	DefaultAuthPort = getEnvInt("DEFAULT_AUTH_PORT", 121)
	DefaultHoarderPort = getEnvInt("DEFAULT_HOARDER_PORT", 142)
	// ... continue with other variables
	DefaultHost = getEnv("DEFAULT_HOST", "127.0.0.1")
	DefaultP2PListenFormat = "/ip4/" + DefaultHost + "/tcp/%d"
	DefaultHTTPListenFormat = "%s://" + DefaultHost + ":%d"
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		// Convert the string to an integer
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

var (
	DefaultAuthPort      = 121
	DefaultHoarderPort   = 142
	DefaultMonkeyPort    = 163
	DefaultPatrickPort   = 184
	DefaultSeerPort      = 205
	DefaultTNSPort       = 226
	DefaultSubstratePort = 282
	DefaultDnsPort       = 304

	DefaultSeerHttpPort      = 403
	DefaultPatrickHttpPort   = 424
	DefaultAuthHttpPort      = 445
	DefaultTNSHttpPort       = 466
	DefaultSubstrateHttpPort = 529

	DreamlandApiListen = DefaultHost + ":1421"
)

var (
	DefaultHost             string = "127.0.0.1"
	DefaultP2PListenFormat  string = "/ip4/" + DefaultHost + "/tcp/%d"
	DefaultHTTPListenFormat string = "%s://" + DefaultHost + ":%d"
)
