package main


import (
	"encoding/json"
	//"flag"
	"fmt"
    "math/rand"
	"log"
	"io/ioutil"
)

const (
	BODY_SIZE       = 5000
	FILE_PREFIX     = "dat"
	DATA_SET_SIZE   = 1000
)

type Message struct {
	Id   int64
	Name string
	Body string
	Time int64
}

func main() {

	fmt.Println("Starting set-sets creation")

	m := Message{1,"Alice", RandStringRunes(BODY_SIZE), rand.Int63()}
	data, err := json.Marshal(m)

	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(FILE_PREFIX, data, 0644)

	if err != nil {
		log.Fatal(err)
	}

}


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
