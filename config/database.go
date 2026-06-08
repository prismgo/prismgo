package config

func init() {
	Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			/*
				|--------------------------------------------------------------------------
				| Default Database Connection Name
				|--------------------------------------------------------------------------
				|
				| Here you may specify which of the database connections below should be
				| used as your default connection for database operations.
				|
			*/
			"default": Env("DATABASE_CONNECTION", Env("DB_CONNECTION", "sqlite")),

			/*
				|--------------------------------------------------------------------------
				| Database Connections
				|--------------------------------------------------------------------------
				|
				| Each connection listed here may be selected by name. Prismgo ships with
				| common local defaults so new applications can run quickly.
				|
			*/
			"connections": map[string]interface{}{
				"sqlite": map[string]interface{}{
					"driver":             "sqlite",
					"database":           Env("DATABASE_NAME", Env("DB_DATABASE", "storage/database.sqlite")),
					"max_open_conns":     Env("DATABASE_MAX_OPEN_CONNS", Env("DB_MAX_OPEN_CONNS", 1)),
					"max_idle_conns":     Env("DATABASE_MAX_IDLE_CONNS", Env("DB_MAX_IDLE_CONNS", 1)),
					"conn_max_lifetime":  Env("DATABASE_CONN_MAX_LIFETIME", Env("DB_CONN_MAX_LIFETIME", "1h")),
					"conn_max_idle_time": Env("DATABASE_CONN_MAX_IDLE_TIME", Env("DB_CONN_MAX_IDLE_TIME", "10m")),
				},
				"mysql": map[string]interface{}{
					"driver":             Env("DATABASE_DRIVER", "mysql"),
					"host":               Env("DATABASE_HOST", Env("DB_HOST", "127.0.0.1")),
					"port":               Env("DATABASE_PORT", Env("DB_PORT", 3306)),
					"database":           Env("DATABASE_NAME", Env("DB_DATABASE", "prismgo")),
					"username":           Env("DATABASE_USER", Env("DB_USERNAME", "root")),
					"password":           Env("DATABASE_PASSWORD", Env("DB_PASSWORD", "")),
					"charset":            Env("DATABASE_CHARSET", Env("DB_CHARSET", "utf8mb4")),
					"parse_time":         Env("DATABASE_PARSE_TIME", Env("DB_PARSE_TIME", true)),
					"loc":                Env("DATABASE_LOC", Env("DB_LOC", "Local")),
					"dsn":                Env("DATABASE_DSN", Env("DB_DSN", "")),
					"max_open_conns":     Env("DATABASE_MAX_OPEN_CONNS", Env("DB_MAX_OPEN_CONNS", 30)),
					"max_idle_conns":     Env("DATABASE_MAX_IDLE_CONNS", Env("DB_MAX_IDLE_CONNS", 10)),
					"conn_max_lifetime":  Env("DATABASE_CONN_MAX_LIFETIME", Env("DB_CONN_MAX_LIFETIME", "1h")),
					"conn_max_idle_time": Env("DATABASE_CONN_MAX_IDLE_TIME", Env("DB_CONN_MAX_IDLE_TIME", "10m")),
				},
			},

			/*
				|--------------------------------------------------------------------------
				| Redis Databases
				|--------------------------------------------------------------------------
				|
				| Redis connections are configured here so cache, queue, session, and
				| Horizon may refer to a shared connection by name.
				|
			*/
			"redis": map[string]interface{}{
				"client":  Env("REDIS_CLIENT", "go"),
				"options": map[string]interface{}{},
				"default": map[string]interface{}{
					"url":           Env("REDIS_URL", ""),
					"scheme":        Env("REDIS_SCHEME", ""),
					"host":          Env("REDIS_HOST", "127.0.0.1"),
					"port":          Env("REDIS_PORT", "6379"),
					"username":      Env("REDIS_USERNAME", ""),
					"password":      Env("REDIS_PASSWORD", ""),
					"database":      Env("REDIS_MAIN_DB", 1),
					"name":          Env("REDIS_NAME", ""),
					"timeout":       Env("REDIS_TIMEOUT", ""),
					"read_timeout":  Env("REDIS_READ_TIMEOUT", ""),
					"write_timeout": Env("REDIS_WRITE_TIMEOUT", ""),
					"max_retries":   Env("REDIS_MAX_RETRIES", ""),
				},
				"cache": map[string]interface{}{
					"url":           Env("REDIS_CACHE_URL", ""),
					"scheme":        Env("REDIS_CACHE_SCHEME", ""),
					"host":          Env("REDIS_HOST", "127.0.0.1"),
					"port":          Env("REDIS_PORT", "6379"),
					"username":      Env("REDIS_USERNAME", ""),
					"password":      Env("REDIS_PASSWORD", ""),
					"database":      Env("REDIS_CACHE_DB", 0),
					"name":          Env("REDIS_CACHE_NAME", ""),
					"timeout":       Env("REDIS_CACHE_TIMEOUT", ""),
					"read_timeout":  Env("REDIS_CACHE_READ_TIMEOUT", ""),
					"write_timeout": Env("REDIS_CACHE_WRITE_TIMEOUT", ""),
					"max_retries":   Env("REDIS_CACHE_MAX_RETRIES", ""),
				},
			},
		}
	})
}
