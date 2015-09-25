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
	BODY_SIZE       = 250
	RESULT_FOLDER   = "tmp"
	FILE_PREFIX     = "data"
	DATA_SET_SIZE   = 10
	DICT            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Message struct {
	Id   int
	Name string
	Body1 string
	Body2 string
	Body3 string
	Body4 string
	Body5 string
	Body6 string
	Body7 string
	Body8 string
	Body9 string
	Body10 string
	Body11 string
	Body12 string
	Body13 string
	Body14 string
	Body15 string
	Body16 string
	Body17 string

	Time int64
}

func main() {

	fmt.Println("Starting set-sets creation")
	if _, err := os.Stat(RESULT_FOLDER); os.IsNotExist(err) {
		err := os.Mkdir(RESULT_FOLDER,0777)
		if err != nil {
			log.Fatal(err)
		}
	}



	for i := 0; i < DATA_SET_SIZE; i++ {
		m := Message{i,"Alice", randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE),
			randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE),
			randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE),
			randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE),
			randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE),
			randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE), randStringRunes(BODY_SIZE),
			rand.Int63(),
		    }
		data, err := json.Marshal(m)

		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(RESULT_FOLDER + "/" + FILE_PREFIX+strconv.Itoa(i), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Created "+ strconv.Itoa(DATA_SET_SIZE) + " files")
}


var letterRunes = []rune(DICT)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
