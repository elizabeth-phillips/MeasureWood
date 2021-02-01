package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elizabeth-phillips/woodMeasure/types"
)

func ReadPlan(fileName string) (cuts []types.Cut) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened firepitChair.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cuts)

	return
}
