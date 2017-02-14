package logger

import "log"

// LogRequest output request information to stderr
func LogRequest(remoteAddr, method, url string) {
	log.Printf("%s %s %s", remoteAddr, method, url)
}
