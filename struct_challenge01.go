package main

import (
    "fmt"
)

type Astro struct {
    name string
    age int
    mission string
    isneeded bool
}

func main() {
    ast1 := Astro{ "Person1", 35, "Apollo", false}
    ast2 := Astro{ "Person2", 40, "Apollo11", false}
    ast3 := Astro{ "Person3", 45, "Voyager", true}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)
}
