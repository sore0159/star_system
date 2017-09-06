package main

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
}

func GetConfig() (*Config, error) {
	return &Config{
		LogFile:     DEFAULT_LOG_FILE_NAME,
		HTTPPort:    HTTP_PORT,
		StaticDir:   FILE_DIR_NAME,
		TemplateDir: DEFAULT_TEMPLATE_DIR,
	}, nil
}
