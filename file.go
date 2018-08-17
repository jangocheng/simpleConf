package simpleConf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var r1, r2 *regexp.Regexp
var lineBreak = byte('\n')
var comment = ";#"

func initReg() {
	if r1 == nil {
		r1 = regexp.MustCompile(fmt.Sprintf(`^\[([^\[\]\s%s]+)\]\s*([%s].*)*$`, comment, comment))                           // regex [sectionName]
		r2 = regexp.MustCompile(fmt.Sprintf(`^([^\[\]\s%s]+)\s*=\s*([^\[\]\s%s]+)\s*([%s].*)*$`, comment, comment, comment)) // regex key = value
	}
}

/**
 * read line to struct Sections
 */
func readFile(filename string) {
	initReg()
	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	defer file.Close()
	if err != nil {
		debugPrintln(err)
		return
	}

	var lineNumber int
	sectionName := "Default"
	buf := bufio.NewReader(file)
	for {
		lineNumber++
		line, err := buf.ReadString(lineBreak)
		if err != nil {
			if err != io.EOF {
				debugPrintln(err)
			}
			return
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.IndexByte(comment, line[0]) != -1 {
			continue
		}

		tmp := getSectionName(line)
		if tmp != "" {
			sectionName = tmp
			continue
		}

		k, v := getKeyValue(line)
		if k == "" {
			debugPrintf("%d: syntax error: %s", lineNumber, line)
			continue
		}
		instance.setValue(sectionName, k, v)
	}
}

func getSectionName(line string) string {
	res := r1.FindStringSubmatch(line)
	if len(res) == 0 {
		return ""
	} else {
		return res[1]
	}
}

func getKeyValue(line string) (string, string) {
	res := r2.FindStringSubmatch(line)
	if len(res) == 0 {
		return "", ""
	} else {
		return res[1], res[2]
	}
}
