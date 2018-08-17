package simpleConf

import (
	"strconv"
	"strings"
	"sync"
)

type sections struct {
	m sync.Map
}

type F []string

func initSections() *sections {
	sections := &sections{}
	return sections
}

func (s *sections) getSection(name string) *Section {
	section, _ := s.m.LoadOrStore(name, initSection())
	sec := section.(*Section)
	return sec
}

func (s *sections) setValue(sectionName string, key string, value string) {
	section := s.getSection(sectionName)
	v1, err := strconv.Atoi(value)
	if err == nil {
		section.m.Store(key, v1)
		return
	}

	v2, err := strconv.ParseFloat(value, 64)
	if err == nil {
		section.m.Store(key, v2)
		return
	}
	section.m.Store(key, strings.Trim(value, "\"'"))
}

func (s *sections) getConf(filename string) {
	readFile(filename)
	addWatchFile(filename)
}

func (s *sections) getBatchConf(filenames F) {
	for _, filename := range filenames {
		s.getConf(filename)
	}
}
