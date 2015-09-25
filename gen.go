package main


import (
	"encoding/json"
	//"flag"
	"fmt"
    "math/rand"
	"log"
	"io/ioutil"
	"strconv"
	"os"
)

const (
	BODY_SIZE       = 5000
	RESULT_FOLDER   = "tmp"
	FILE_PREFIX     = "data"
	DATA_SET_SIZE   = 1000000
	DICT            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Message struct {
	Id   int64
	Name string
	Body string
	Time int64
}

func main() {

	fmt.Println("Starting set-sets creation")
	if _, err := os.Stat(RESULT_FOLDER); os.IsNotExist(err) {
		err := os.Mkdir(RESULT_FOLDER,0644)
		if err != nil {
			log.Fatal(err)
		}
	}



	for i := 0; i < DATA_SET_SIZE; i++ {
		m := Message{1,"Alice", RandStringRunes(BODY_SIZE), rand.Int63()}
		data, err := json.Marshal(m)

		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(RESULT_FOLDER + "/" + FILE_PREFIX+strconv.Itoa(i), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

}


var letterRunes = []rune(DICT)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
