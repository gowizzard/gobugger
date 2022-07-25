package gobugger

import (
	"log"
	"testing"
)

const (
	path = "/home/user/path/to/folder"
)

// Test is to test the debugger function
// We add a map with test data into the debugger
func Test(t *testing.T) {

	c := Config{
		Path: path,
	}

	c.Start()

	testDataMap := make(map[string]interface{})

	testDataMap["Name"] = "Outdoor Elements™ II Flannel Shirt für Männer"
	testDataMap["Sku"] = "outdoor-elements-ii-flannel-shirt-fur-manner-1959661"
	testDataMap["Price"] = "79,99"
	testDataMap["Color"] = "Canyon Gold Oversize Tartan"
	testDataMap["Size"] = "XL"
	testDataMap["Description"] = "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look."

	c.Add(testDataMap)

	err := c.End()
	if err != nil {
		log.Fatalln(err)
	}

}
