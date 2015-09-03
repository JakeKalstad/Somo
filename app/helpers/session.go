package helpers

import (
	"github.com/garyburd/redigo/redis"
    "encoding/json"
    "bytes"
    "fmt"
)

type Session struct {
	C redis.Conn
}

func InitSession () *Session {
	s := Session{}
	return &s
}
func (s Session) InitRedis() {
	c, err := redis.Dial("tcp", ":6379")
	s.C = c
	if (err != nil) {
		fmt.Printf("error", err)
	}
}
func (s Session) Get(sessionId string, model interface{})  {
	c, err := redis.Dial("tcp", ":6379")
	str, err := redis.String(c.Do("GET", sessionId))
	if err != nil {
    	fmt.Println(err)
	}
	b := bytes.NewBufferString(str)
	decoder := json.NewDecoder(b)
    decoder.Decode(model)
}
func (s Session) Set(sessionId string, content interface{}) {
	b, err := json.Marshal(content)
	if err != nil {
    	fmt.Println(err)
	}
	if(s.C == nil) {
		fmt.Println("NULLLL")
	}
	c, err := redis.Dial("tcp", ":6379")
    c.Do("SET", sessionId, string(b))
}