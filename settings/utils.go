package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// AppSettings app settnigs
	AppSettings Settings
)

// ReadSettings to init app settings
func ReadSettings(filepath string) Settings {
	var appParams Settings
	doc, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Printf("err is = %e\n", err)
	}

	err = json.Unmarshal(doc, &appParams)
	if err != nil {
		//		panic(err)
		fmt.Printf("err is = %e\n", err)
	}

	return appParams
}

