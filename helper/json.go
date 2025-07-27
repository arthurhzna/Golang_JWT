package helper

import "fmt"

func ErrorConditionCheck(err error) {
    if err != nil {
        fmt.Printf("Error occurred: %v\n", err)  // Tambahkan logging
        panic(err)
    }
}