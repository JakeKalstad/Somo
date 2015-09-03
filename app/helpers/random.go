package helpers

import (
	"fmt"
	"encoding/base64"
    "crypto/rand"
)
func GenerateHandler() string {
	size := 32
   	rb := make([]byte,size)
   	_, err := rand.Read(rb)
   	if err != nil {
    	fmt.Println(err)
   	}
   	rs := base64.URLEncoding.EncodeToString(rb)
   	return rs
}