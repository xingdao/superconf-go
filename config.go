package superconf

import (
	"strconv"
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

var allConfig = map[string]interface{} {}


func must(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *zk.Conn {
	servers := strings.Split("192.168.10.139:31081", ",")
	conn, _, err := zk.Connect(servers, time.Second)
	must(err)
	return conn
}


func watch(conn *zk.Conn, path string, key interface{})  {

	dates := make(chan []byte)
	errors := make(chan error)

	go func() {
		for {
			t, _, events, err := conn.GetW(path)
			if err != nil {
				errors <- err
				return
			}
			dates <- t
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}
			fmt.Println("GetW ")
		}
	}()

	go func() {
		for {
			select {
			case data := <-dates:
				switch key.(type) {
				case *int:
					t, _ := key.(*int)
					*t, _ = strconv.Atoi(string(data[:]))
				case *string:
					t, _ := key.(*string)
					*t = string(data[:])
				default:
					json.Unmarshal(data, &key)
				}
			case err := <-errors:
				fmt.Printf("watchStr error %+v\n", err)
				conn.Close()
				//panic(err)
				return
			}
		}
	}()

	data, _, _ := conn.Get(path)
	switch key.(type) {
	case *int:
		t, _ := key.(*int)
		*t, _ = strconv.Atoi(string(data[:]))
	case *string:
		t, _ := key.(*string)
		*t = string(data[:])
	default:
		json.Unmarshal(data, &key)
	}

}

func Init()  {
	conn := connect()
	for k, v := range allConfig {
		watch(conn, k, v)
	}
}
