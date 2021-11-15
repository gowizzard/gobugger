//********************************************************************************************************************//
//
// This file is part of gobugger.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

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
	Debugger string        `json:"debugger"`
	Prefix   string        `json:"prefix"`
	Start    time.Time     `json:"start"`
	End      time.Time     `json:"end"`
	Data     []interface{} `json:"data"`
}

// Start is to start the debugger
func (c *Config) Start() {

	// Set start time
	c.DebugStart = time.Now()

}

// Add is to add a new entry to debugger
func (c *Config) Add(debugData interface{}) {

	// Add debug data
	c.DebugData = append(c.DebugData, debugData)

}

// End is to end the debugger
func (c *Config) End() error {

	// Set end time
	c.DebugEnd = time.Now()

	// Check prefix length
	if len(c.Prefix) == 0 {
		c.Prefix = "gobugger"
	}

	// Create debug information
	output := Output{
		Debugger: "gobugger",
		Prefix:   c.Prefix,
		Start:    c.DebugStart,
		End:      c.DebugEnd,
		Data:     c.DebugData,
	}

	// Create path
	path := fmt.Sprintf("%s/%s", c.Path, c.DebugStart.Format("2006-01-02"))

	// Create folder
	os.Mkdir(path, os.ModePerm)

	// Marshal json data
	convert, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		return err
	}

	// Create file path
	file := fmt.Sprintf("%s/%s-%s.json", path, c.Prefix, c.DebugStart.Format("2006-01-02T15-04-05"))

	// Write file
	err = ioutil.WriteFile(file, convert, 0644)
	if err != nil {
		return err
	}

	// Return nothing
	return nil

}
