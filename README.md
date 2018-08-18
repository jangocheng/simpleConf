# simpleConf
​	**go version >= 1.9   because use sync.Map to make sure thread safe**	

​	Simply way to load the fixed-type like ini file and get value

```ini
; use ';' or '#' to comment
[test] ; section name
name  = simpleConf # string
float = 12.0       # float
str   = "12.0"     # string
int   = 12         # int
str   = "12"       # string, will overwrite last str
```

​	After load configure file,  you can get section and get value by key in any file (thread safe)

```go	
// main.go
simpleConf.GetConf("./conf/myconf.conf")
```



```go
// other.go
name := simpleConf.GetSection("test").GetStr("name")
```

hot load configure file if you need (thread safe)

```go
simpleConf.HotLoad()
```



# example

```go
package main

import "github.com/me1onrind/simpleIni"
import "fmt"

func init() {
    / console output debug info
	// default is false
	simpleConf.IsDebug = true

	// Only hot load the file loaded after HotLoad()
	// It's not necessary
	simpleConf.HotLoad()

	// load single file
	simpleConf.GetConf("./test.ini")

	// load files
	simpleConf.GetBatchConf(simpleConf.F{
		"./test.ini",
		"./test1.ini",
	})
}

func main() {
	/*
	 * [db@test]
	 * host=127.0.0.1
	 * port = 3306
	 */
	section := simpleConf.GetSection("db1@test")
	fmt.Println(section.GetInt("port")) // 3306
	fmt.Println(section.GetStr("port")) // ""
	fmt.Println(section.GetStr("host")) // 127.0.0.1
}
```



# API

GetConf : load single configure file  
GetBatchConf: load configure files  
HotLoad: use inotify system call to litsen file change  
GetSection: get section (return an Section struct pointer)  
GetStr/GetInt/GetFloat/GetInterface: get value under section

# option

IsDebug : show debug info in console

