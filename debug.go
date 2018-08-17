package simpleConf

import (
	"log"
)

func debugPrintf(format string, values ...interface{}) {
	if IsDebug {
		log.Printf("[simpleIni-debug] "+format, values...)
	}
}

func debugPrintln(values ...interface{}) {
	if IsDebug {
		log.Print("[simpleIni-debug] ")
		log.Println(values...)
	}
}
