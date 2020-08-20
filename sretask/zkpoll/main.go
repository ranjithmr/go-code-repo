package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sync"
	"time"
)

type broker struct {
	Name     string
	Services []innerField
}

type innerField struct {
	Service   string
	Container interface{}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(filename string) {

	javacommand := "pwd > " + filename
	cmd := exec.Command("/bin/bash", "-c", javacommand)
	err := cmd.Run()
	checkError(err)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	// executing system command
	runCommand("/tmp/testfile_atl")
	runCommand("/tmp/testfile_eat")
	//reading from json file
	wg.Add(2)
	brokers := jsonDecoder("myatl.json")
	go checkError(updatedb(brokers, "ATL", &wg))

	brokers = jsonDecoder("myeat.json")
	go checkError(updatedb(brokers, "EAT", &wg))
	wg.Wait()
	fmt.Println("Time: ", time.Since(start))
}
func jsonDecoder(filename string) []broker {

	data, err := ioutil.ReadFile(filename)
	checkError(err)
	//Json decoding
	var cont map[string]interface{}
	checkError(json.Unmarshal(data, &cont))

	var bkr broker
	var brokers []broker
	for _, reg := range cont {
		v := reg.(map[string]interface{})
		for k, val := range v {

			bkr = broker{
				Name:     k,
				Services: nestedParser(val), //additional parsing the values
			}
			brokers = append(brokers, bkr)
		}
	}
	return brokers
}

func nestedParser(val interface{}) []innerField {
	var innerFields []innerField

	services, ok := val.(map[string]interface{})
	if ok {
		for svcname, values := range services {
			inrFld := innerField{
				Service:   svcname,
				Container: values,
			}
			innerFields = append(innerFields, inrFld)
		}
	}
	return innerFields
}
