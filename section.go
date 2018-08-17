package simpleConf

import (
	"sync"
)

type Section struct {
	m sync.Map
}

func initSection() *Section {
	section := &Section{}
	return section
}

func (s *Section) GetInt(key string) int {
	value, ok := s.m.Load(key)
	if ok {
		intVal, valid := value.(int)
		if valid {
			return intVal
		}
	}
	return 0
}

func (s *Section) GetStr(key string) string {
	value, ok := s.m.Load(key)
	if ok {
		intVal, valid := value.(string)
		if valid {
			return intVal
		}
	}
	return ""
}

func (s *Section) GetFloat(key string) float64 {
	value, ok := s.m.Load(key)
	if ok {
		intVal, valid := value.(float64)
		if valid {
			return intVal
		}
	}
	return 0
}

func (s *Section) GetInterface(key string) interface{} {
	value, _ := s.m.Load(key)
	return value
}
