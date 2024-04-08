package common

import (
	"log"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogIfError(err error) {
	if err != nil {
		log.Println(err)
	}
}
