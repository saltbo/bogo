package types

type ConfigInterface interface {
	Get(key string)
}

type RouterInterface interface {
}

type DatabaseInterface interface {
}

type CacheInterface interface {
}

type LoggerInterface interface {
}
