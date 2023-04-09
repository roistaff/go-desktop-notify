package main

import (
    "github.com/gen2brain/beeep"
)

func main() {
    title := "Hello"
    message := "Hello,World!"
    iconPath := "hellocreater.png"
    err := beeep.Notify(title, message, iconPath)
    if err != nil {
        panic(err)
    }
}
