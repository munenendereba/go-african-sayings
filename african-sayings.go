package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Languages struct {
	Languages []Lang `json:"languages"`
}

type Lang struct {
	Lang string `json:"lang"`
}

type Saying struct {
	Saying      string `json:"saying"`
	Translation string `json:"translation"`
}

type Sayings struct {
	Sayings []Saying `json:"sayings"`
}

func AfricanSaying(langIn string, allSayings bool) string {
	homeFolder := "./resources/"
	fileSuffix := "-sayings.json"

	if langIn != "" {
		langFilePath := homeFolder + langIn + fileSuffix

		sayings, err := GetSayings(langFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(sayings)
	} else {
		jsonFile, err := os.Open(homeFolder + "languages.json")

		if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("Successfully Opened languages.json")
		}

		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Fatalln(err)
		}

		var languages Languages

		json.Unmarshal(byteValue, &languages)

		for i := 0; i < len(languages.Languages); i++ {
			filepath := homeFolder + languages.Languages[i].Lang + fileSuffix

			sayings, err := GetSayings(filepath)

			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(sayings)
		}

	}

	return ""
}

func GetSayings(filename string) (Sayings, error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err.Error())
		return Sayings{}, err
	}

	defer file.Close()

	fileBytesValue, err2 := io.ReadAll(file)

	if err2 != nil {
		log.Fatalln(err2.Error())

		return Sayings{}, err2
	}

	var sayings Sayings

	errUnmarshall := json.Unmarshal(fileBytesValue, &sayings)

	if errUnmarshall != nil {
		log.Fatalln(errUnmarshall.Error())

		return Sayings{}, errUnmarshall
	}

	return sayings, nil
}
