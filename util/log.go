package util

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

//LogLevelInfo Level INFO
const LogLevelInfo = "INFO"

//LogLevelDebug Level DEBUG
const LogLevelDebug = "DEBUG"

//LogLevelError Level ERROR
const LogLevelError = "ERROR"

type jsonMsg struct {
	Msg   string    `json:"msg"`
	Time  time.Time `json:"time"`
	Level string    `json:"level"`
}

func (jMsg *jsonMsg) MarshalJSON() ([]byte, error) {
	log.Println("FOO")

	type Alias jsonMsg

	return json.Marshal(&struct {
		*Alias
		Time string `json:"time"`
	}{
		Time: time.Now().Format(time.RFC3339),
		//initializing internal Alias struct
		Alias: (*Alias)(jMsg),
	})
}

//Log Simple log wrapper
func Log(message string, level string) {
	var jsonMessage = jsonMsg{
		Msg:   message,
		Time:  time.Now(),
		Level: strings.ToUpper(level),
	}

	result, marshalError := json.Marshal(&jsonMessage)

	if marshalError != nil {
		Log("Error decoding message.", LogLevelError)
	} else {
		log.Println(result)
	}
}
