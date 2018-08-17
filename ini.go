package simpleIni

var instance *sections

func init() {
	instance = initSections()
}

func GetSection(name string) *Section {
	return instance.getSection(name)
}

func GetConf(filename string) {
	instance.getConf(filename)
}

func GetBatchConf(filenames F) {
	instance.getBatchConf(filenames)
}
