package main

import (
  "fmt"
  "github.com/google/uuid"
  "github.com/chokunlim/go-again/sound/greeting"
)

func main() {
  id := uuid.New()
  fmt.Printf("Generated UUID: %s\n", id)
  greeting.introduce()
  fmt.Println("hi")
}