# gobugger

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/gobugger.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/gobugger/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/gobugger/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/gobugger)](https://goreportcard.com/report/github.com/gowizzard/gobugger) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/gobugger.svg)](https://pkg.go.dev/github.com/gowizzard/gobugger) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/gobugger)](https://github.com/gowizzard/gobugger/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/gobugger)](https://github.com/gowizzard/gobugger/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/gobugger)](https://github.com/gowizzard/gobugger/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/gobugger)](https://github.com/gowizzard/gobugger/blob/master/LICENSE) 

With this small library it should be possible to debug easier in json files. Instead of always using the console.

# Install

```console
go get github.com/gowizzard/gobugger
```

# How to use?

Actually, the library is super simple. You start a logging process and feed it with data via another function. If you want to stop the debugging, then there is also a function for this. Now the data is stored in a folder structure directly at the desired location.

Here you will find an example:

```go
// Create new debugger
debugger := gobugger.Config{
    Path: "/Users/jonaskwiedor/Downloads",
}

// Start debugger
debugger.Start()

// Add test data with a struct
testDataStruct := TestStruct{
    Name:        "Outdoor Elements™ II Flannel Shirt für Männer",
    Sku:         "outdoor-elements-ii-flannel-shirt-fur-manner-1959661",
    Price:       "79,99",
    Color:       "Canyon Gold Oversize Tartan",
    Size:        "XL",
    Description: "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look.",
}

debugger.Add(testDataStruct)

// End debugger
err := debugger.End()
if err != nil {
    fmt.Println(err)
}
```

The whole thing is now saved in a JSON file and looks like this:

```json
{
    "debugger": "gobugger",
    "start": "2021-11-13T13:09:26.972983+01:00",
    "end": "2021-11-13T13:09:26.972983+01:00",
    "data": [
        {
            "name": "Outdoor Elements™ II Flannel Shirt für Männer",
            "sku": "outdoor-elements-ii-flannel-shirt-fur-manner-1959661",
            "price": "79,99",
            "color": "Canyon Gold Oversize Tartan",
            "size": "XL",
            "description": "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look."
        }
    ]
}
```

The nice thing is that you can also debug a wide variety of data types. Here is a larger example:

```go
// Create new debugger
debugger := gobugger.Config{
    Path: "/Users/jonaskwiedor/Downloads",
}

// Start debugger
debugger.Start()

// Add test data with a map
testDataMap := make(map[string]interface{})

testDataMap["Name"] = "Outdoor Elements™ II Flannel Shirt für Männer"
testDataMap["Sku"] = "outdoor-elements-ii-flannel-shirt-fur-manner-1959661"
testDataMap["Price"] = "79,99"
testDataMap["Color"] = "Canyon Gold Oversize Tartan"
testDataMap["Size"] = "XL"
testDataMap["Description"] = "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look."

debugger.Add(testDataMap)

// Add test data with an array
testDataArray := []string{
    "Outdoor Elements™ II Flannel Shirt für Männer",
    "outdoor-elements-ii-flannel-shirt-fur-manner-1959661",
    "79,99",
    "Canyon Gold Oversize Tartan",
    "XL",
    "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look.",
}

debugger.Add(testDataArray)

// Add test data with a struct
testDataStruct := TestStruct{
    Name:        "Outdoor Elements™ II Flannel Shirt für Männer",
    Sku:         "outdoor-elements-ii-flannel-shirt-fur-manner-1959661",
    Price:       "79,99",
    Color:       "Canyon Gold Oversize Tartan",
    Size:        "XL",
    Description: "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look.",
}

debugger.Add(testDataStruct)

// End debugger
err := debugger.End()
if err != nil {
    fmt.Println(err)
}
```

The whole thing then looks like this when you look at the debug file:

```json
{
    "debugger": "gobugger",
    "start": "2021-11-13T13:09:26.972983+01:00",
    "end": "2021-11-13T13:09:26.972983+01:00",
    "data": [
        {
            "Color": "Canyon Gold Oversize Tartan",
            "Description": "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look.",
            "Name": "Outdoor Elements™ II Flannel Shirt für Männer",
            "Price": "79,99",
            "Size": "XL",
            "Sku": "outdoor-elements-ii-flannel-shirt-fur-manner-1959661"
        },
        [
            "Outdoor Elements™ II Flannel Shirt für Männer",
            "outdoor-elements-ii-flannel-shirt-fur-manner-1959661",
            "79,99",
            "Canyon Gold Oversize Tartan",
            "XL",
            "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look."
        ],
        {
            "name": "Outdoor Elements™ II Flannel Shirt für Männer",
            "sku": "outdoor-elements-ii-flannel-shirt-fur-manner-1959661",
            "price": "79,99",
            "color": "Canyon Gold Oversize Tartan",
            "size": "XL",
            "description": "Mit seinem traditionellen Karomuster, dem klassischen Kragen und den langen Ärmeln bringt dir dieses Hemd einen Casual-Look."
        }
    ]
}
```
