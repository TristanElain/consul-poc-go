package utils

import (
	"log"
)

// HandleError - affiche l'erreur et stope le programme
func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
