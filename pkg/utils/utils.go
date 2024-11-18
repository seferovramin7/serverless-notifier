package utils

import "log"

func LogError(err error) {
	if err != nil {
		log.Printf("Error %v", err)
	}
}

func LogInfo(info string) {
	log.Printf("Info %s", info)
}
