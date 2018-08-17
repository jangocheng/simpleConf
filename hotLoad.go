package simpleIni

import (
	"sync"
	"syscall"
	"unsafe"
)

const (
	IN_MASK    = syscall.IN_MODIFY
	HANDLE_NUM = 1
)

type hotLoadMap struct {
	sync.RWMutex
	m map[int32]string
}

var h hotLoadMap
var fd int

func HotLoad() (err error) {
	fd, err = syscall.InotifyInit()
	h.m = make(map[int32]string)
	if err != nil {
		debugPrintln(err)
		return err
	}
	go loop()
	return nil
}

func addWatchFile(filename string) {
	if fd == 0 {
		debugPrintln("HotLoad() is stopping")
		return
	}
	wd, err := syscall.InotifyAddWatch(fd, filename, IN_MASK)
	if err != nil {
		debugPrintln(err)
		return
	}
	h.Lock()
	h.m[int32(wd)] = filename
	h.Unlock()
}

func loop() {
	var buffer [HANDLE_NUM * syscall.SizeofInotifyEvent]byte
	for {
		syscall.Read(fd, buffer[:])
		event := (*syscall.InotifyEvent)(unsafe.Pointer(&buffer))
		h.RLock()
		filename := h.m[event.Wd]
		h.RUnlock()
		instance.getConf(filename)
	}
}
