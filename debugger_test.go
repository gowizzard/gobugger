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
	"fmt"
	"testing"
)

type TestStruct struct {
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	Price       string `json:"price"`
	Color       string `json:"color"`
	Size        string `json:"size"`
	Description string `json:"description"`
}

func Test(t *testing.T) {

	// Create new debugger
	c := Config{
		Prefix: "products",
		Path:   "test-files",
	}

	// Start debugger
	c.Start()

	// Add test data with a map
	testDataMap := make(map[string]interface{})

	testDataMap["Name"] = "Outdoor Elements™ II Flannel Shirt für Männer"
	testDataMap["Sku"] = "outdoor-elements-ii-flannel-shirt-fur-manner-1959661"
	testDataMap["Price"] = "79,99"
	testDataMap["Color"] = "Canyon Gold Oversize Tartan"
	testDataMap["Size"] = "XL"
	testDataMap["Description"] = "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look."

	c.Add(testDataMap)

	// Add test data with an array
	testDataArray := []string{
		"Outdoor Elements™ II Flannel Shirt für Männer",
		"outdoor-elements-ii-flannel-shirt-fur-manner-1959661",
		"79,99",
		"Canyon Gold Oversize Tartan",
		"XL",
		"Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look.",
	}

	c.Add(testDataArray)

	// Add test data with a struct
	testDataStruct := TestStruct{
		Name:        "Outdoor Elements™ II Flannel Shirt für Männer",
		Sku:         "outdoor-elements-ii-flannel-shirt-fur-manner-1959661",
		Price:       "79,99",
		Color:       "Canyon Gold Oversize Tartan",
		Size:        "XL",
		Description: "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look.",
	}

	c.Add(testDataStruct)

	// End debugger
	err := c.End()
	if err != nil {
		fmt.Println(err)
	}

}
