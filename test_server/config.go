package main

import (
	"fmt"
	"os"

	"github.com/gorilla/securecookie"
	"github.com/sore0159/star_system/data"
	db "github.com/sore0159/star_system/mockdb"
	ssr "github.com/sore0159/star_system/route"
)

const HTTP_PORT = ":8000"
const FILE_DIR_NAME = "FILES/"
const DEFAULT_DATA_FILE_NAME = FILE_DIR_NAME + "DATA_FILE.json"
const DEFAULT_PROXY_ADDR = ""
const DEFAULT_LOG_FILE_NAME = FILE_DIR_NAME + "LOG_FILE.txt"
const DEFAULT_TEMPLATE_DIR = FILE_DIR_NAME + "templates/"

type Config struct {
	LogFile     string
	HTTPPort    string
	StaticDir   string
	TemplateDir string
	DataFile    string
}

func GetConfig() (*Config, error) {
	return &Config{
		LogFile:     DEFAULT_LOG_FILE_NAME,
		HTTPPort:    HTTP_PORT,
		StaticDir:   FILE_DIR_NAME,
		TemplateDir: DEFAULT_TEMPLATE_DIR,
		DataFile:    DEFAULT_DATA_FILE_NAME,
	}, nil
}

type Resources struct {
	Log      *Logger
	Key      *securecookie.SecureCookie
	Provider data.Provider
}

func GetResources(cfg *Config) (*Resources, error) {
	l, err := NewLogger(cfg)
	if err != nil {
		return nil, fmt.Errorf("logger creation failure: %v\n", err)
	}
	p, err := db.LoadMockProvider(cfg.DataFile)
	if err != nil {
		if os.IsNotExist(err) {
			p = db.NewMockProvider()
		} else {
			return nil, fmt.Errorf("academy load failure: %v\n", err)
		}
	}
	key := ssr.NewCookieSecurity(priv_HASHKEY, priv_BLOCKKEY)
	return &Resources{
		Log:      l,
		Key:      key,
		Provider: p,
	}, nil
}
