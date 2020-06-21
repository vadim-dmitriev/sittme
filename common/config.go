package common

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config структура, которая описывает настройки приложения
type Config struct {

	// Server описывает настройки, относящиеся к серверу
	Server struct {
		Port         int
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	// Timeout это время работы таймера при переходе
	// из состояния Interrupted в Finished
	Timeout time.Duration
}

// NewConfig создает заполненный объект Config
// Испочником настроек являются ключи запуска и переменные
// окружения
func NewConfig() (*Config, error) {
	port := flag.Int("port", 8080, "Server port")
	readTimeout := flag.Int("rtimeout", 5, "Server read timeout in seconds")
	writeTimeout := flag.Int("wtimeout", 5, "Server write timeout in seconds")
	timeout := flag.Int("timeout", 5, "Transition time from state Interrupted to Finished")
	flag.Parse()

	env, ok := os.LookupEnv("PORT")
	if ok {
		pTmp, err := strconv.Atoi(env)
		if err != nil {
			return nil, fmt.Errorf("$PORT env value has invalid syntax '%s'", env)
		}
		*port = pTmp
	}

	env, ok = os.LookupEnv("RTIMEOUT")
	if ok {
		rtimeout, err := strconv.Atoi(env)
		if err != nil {
			return nil, fmt.Errorf("$RTIMEOUT env value has invalid syntax '%s'", env)
		}
		*readTimeout = rtimeout
	}

	env, ok = os.LookupEnv("WTIMEOUT")
	if ok {
		wtimeout, err := strconv.Atoi(env)
		if err != nil {
			return nil, fmt.Errorf("$WTIMEOUT env value has invalid syntax '%s'", env)
		}
		*writeTimeout = wtimeout
	}

	env, ok = os.LookupEnv("TIMEOUT")
	if ok {
		t, err := strconv.Atoi(env)
		if err != nil {
			return nil, fmt.Errorf("$TIMEOUT env value has invalid syntax '%s'", env)
		}
		*timeout = t
	}

	cfg := &Config{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	cfg.Server.Port = *port
	cfg.Server.ReadTimeout = time.Duration(*readTimeout) * time.Second
	cfg.Server.WriteTimeout = time.Duration(*writeTimeout) * time.Second

	return cfg, nil
}
