package simpleConf

import (
	"log"
)

func warning(format string, values ...interface{}) {
	if IsDebug {
		log.Printf("[simpleConf-WARNING] "+format, values...)
	}
}

//func info(format string, values ...interface{}) {
//if IsDebug {
//log.Printf("[simpleIni-info] "+format, values...)
//}
//}

//func notice(format string, values ...interface{}) {
//if IsDebug {
//log.Printf("[simpleIni-notice] "+format, values...)
//}
