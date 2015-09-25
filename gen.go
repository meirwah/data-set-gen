package main


import (
	"encoding/json"
	"flag"
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

var (
	bodyFieldSize = flag.Int("fieldsize", BODY_SIZE, "size of each String field")
	folder = flag.String("folder", RESULT_FOLDER, "Folder to place result data set")
	dataSetSize     = flag.Int("setsize", DATA_SET_SIZE, "amount of files to be created")
	filePrefix       = flag.String("fileprefix", "", "the each file prefix , suffix will be int ID")
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
    flag.Parse()

	fmt.Println("Starting set-sets creation")
	if _, err := os.Stat(*folder); os.IsNotExist(err) {
		err := os.Mkdir(*folder,0777)
		if err != nil {
			log.Fatal(err)
		}
	}



	for i := 0; i < *dataSetSize; i++ {
		m := Message{i,"Alice", randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize),
			randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize),
			randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize),
			randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize),
			randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize),
			randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize), randStringRunes(*bodyFieldSize),
			rand.Int63(),
		    }
		data, err := json.Marshal(m)

		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(*folder + "/" + *filePrefix+strconv.Itoa(i), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Created "+ strconv.Itoa(*dataSetSize) + " files")
}


var letterRunes = []rune(DICT)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
