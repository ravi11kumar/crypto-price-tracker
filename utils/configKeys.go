package utils

type ConfigKey string

const (
	ServerPort  ConfigKey = "server.port"
	CacheExpiry ConfigKey = "cache.expiry"
	BaseApiUrl  ConfigKey = "baseApiUrls.coindesk"
)
