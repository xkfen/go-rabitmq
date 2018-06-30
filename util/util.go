package util

import (
	"log"
	"fmt"
)
func FailOnErr(err error, msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}