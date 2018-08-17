package simpleIni

import (
	"fmt"
	"os"
	"testing"
)

func Test_GetConf_1(t *testing.T) {
	GetConf("test.ini")
	rangeSection("db1@test", t)
	rangeSection("db2@online", t)
}

func Test_GetConf_2(t *testing.T) {
	GetConf("./testData/test1.ini")
	rangeSection("type@test1", t)
}

func Test_GetConf_3(t *testing.T) {
	chs := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		chs[i] = make(chan int)
		c := chs[i]
		go func(i int) {
			GetConf(fmt.Sprintf("./testData/test%d.ini", i+2))
			c <- 1
		}(i)
	}
	for _, c := range chs {
		<-c
	}
	rangeSection("user@db", t)
	rangeSection("user6@db", t)
}

func Test_GetConf_4(t *testing.T) {
	HotLoad()
	file, _ := os.OpenFile("./test2.ini", os.O_WRONLY|os.O_CREATE, 0755)
	defer file.Close()
	file.WriteString("[testHotload]\n")
	file.WriteString("port=3306\n")
	file.Sync()
	GetConf("./test2.ini")
	t.Log(GetSection("testHotload").GetInt("port"))
	file.WriteString("port=80\n")
	file.Sync()
	t.Log(GetSection("testHotload").GetInt("port"))
	os.Remove("./test2.ini")
}

func rangeSection(sectionName string, t *testing.T) {
	t.Log("***************************************************")
	t.Log("section:", sectionName)
	t.Log("===================================================")
	s := GetSection(sectionName)
	s.m.Range(func(key, value interface{}) bool {
		_, ok := value.(string)
		if ok {
			t.Log(key, "= (string)", value)
		}
		_, ok = value.(int)
		if ok {
			t.Log(key, "= (int)", value)
		}
		_, ok = value.(float64)
		if ok {
			t.Log(key, "= (float64)", value)
		}
		return true
	})
	t.Log("***************************************************")
	t.Log("")
}
