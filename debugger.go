package gobugger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Config is to save the config for debugger
type Config struct {
	Prefix     string
	Path       string
	DebugStart time.Time
	DebugEnd   time.Time
	DebugData  []interface{}
}

// Output is to save the output
type Output struct {
	Prefix string        `json:"prefix"`
	Start  time.Time     `json:"start"`
	End    time.Time     `json:"end"`
	Data   []interface{} `json:"data"`
}

// Start is to start the debugger
func (c *Config) Start() {
	c.DebugStart = time.Now()
	c.Prefix = "gobugger"
}

// Add is to add a new entry to debugger
func (c *Config) Add(debugData interface{}) {
	c.DebugData = append(c.DebugData, debugData)
}

// End is to set the debug end date, format the output data
// And write the date in a json file
func (c *Config) End() error {

	c.DebugEnd = time.Now()

	output := Output{
		Prefix: c.Prefix,
		Start:  c.DebugStart,
		End:    c.DebugEnd,
		Data:   c.DebugData,
	}

	path := fmt.Sprintf("%s/%s", c.Path, c.DebugStart.Format("2006-01-02"))

	_ = os.Mkdir(path, os.ModePerm)

	convert, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		return err
	}

	file := fmt.Sprintf("%s/%s-%s.json", path, c.Prefix, c.DebugStart.Format("20060102T150405"))

	err = ioutil.WriteFile(file, convert, 0644)
	if err != nil {
		return err
	}

	return nil

}
