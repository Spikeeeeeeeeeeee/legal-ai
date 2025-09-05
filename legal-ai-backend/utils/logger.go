package utils

import "log"

func InitLogger() {
	log.SetPrefix("[legal-ai] ")
}

func LogInfo(msg string, v ...interface{}) {
	log.Printf("[INFO] %s: %v\n", msg, v)
}

func LogError(msg string, err error) {
	log.Printf("[ERROR] %s: %v\n", msg, err)
}
