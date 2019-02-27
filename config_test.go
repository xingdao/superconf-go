package superconf

import (
	"fmt"
	"time"
	"testing"
)


func TestMain(m *testing.M) {

	type response struct {
		Page   int      `json:"page"`
		Fruits []string `json:"fruits"`
	}
	var inta int

	var stringb string

	var structc response

	allConfig = map[string]interface{} {
		"/mirror/int": &inta,
		"/mirror/string": &stringb,
		"/mirror/struct": &structc}

	Init()
	go func() {
		for {
			fmt.Printf("string value: %s\n", stringb)
			fmt.Printf("int value: %d\n", inta)
			fmt.Printf("struct value: %+v\n", structc)
			time.Sleep(time.Second)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}

