# simpleIni
Simply read the ini type file

# example

```go
package main

import "github.com/me1onrind/simpleIni"
import "fmt"

func init() {
    // Global variable is recommended
    // But you can use the variable you defined
    simpleIni.Instance = simplinIni.InitSections()
    // Only hot load the file loaded after HotLoad()
    // It's necessary
    Instance.HotLoad()
    // There are no restrictions on suffixes
    Instance.GetConf("./conf/db.ini")
    Instance.GetConf("./conf/api.conf")
}

func main() {
    section  := simpleIni.Instace.GetSection("db@user")
    port     := section.GetInt("port")
    host t   := section.GetStr("host")
    username := section.GetStr("username")
    password := section.GetStr("password")
    table    := section.GetSte("table")
    fmt.Printf("host is %s", host)
    fmt.Printf("port is %d", port)
    fmt.Printf("username is %s", username)
    fmt.Printf("password is %s", password)
    fmt.Printf("table is %s", table)
    return
}
```



# api

# option
