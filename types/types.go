package types

type ConfigInterface interface {
	Get(key string)
}

type DatabaseInterface interface {
}

type CacheInterface interface {
}

type LoggerInterface interface {
}
