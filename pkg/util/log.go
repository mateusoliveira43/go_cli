package util

import "log"

func init() {
	log.SetFlags(log.Lmsgprefix)
}

func Debug(message string) {
	log.SetPrefix("DEBUG: \t")
	log.Println(message)
}

func Info(message string) {
	log.SetPrefix("INFO: \t")
	log.Println(message)
}

func Warn(message string) {
	log.SetPrefix("WARN: \t")
	log.Println(message)
}

func Error(message string) {
	log.SetPrefix("ERROR: \t")
	log.Println(message)
}

func Fatal(message string) {
	log.SetPrefix("FATAL: \t")
	log.Fatalln(message)
}
