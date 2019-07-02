package config

// DBCfg database general config
type DBCfg struct {
	ConnectionURI string
}

// ServerCfg Server general config
type ServerCfg struct {
	Port int
}

// RedisCfg Redis general config
type RedisCfg struct {
	ConnectionURI string
}
