package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var name, address string

    fmt.Print("Enter a name: ")
    fmt.Scanln(&name)

    fmt.Print("Enter an address: ")
    fmt.Scanln(&address)

    person := make(map[string]string)
    person["name"] = name
    person["address"] = address

    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Here is the JSON data")
    fmt.Println(string(jsonData))
}

