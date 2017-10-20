package main

import (
    "fmt"
)

type World struct {
    Hello
}

func (w *World) Get() string {
    return fmt.Sprintf("world %s", w.Hello.Get())
}

type Hello struct {
    Name string

}

func (h *Hello) Get() string {
    return h.Name
}

func (h *Hello) Set(s string) {
    h.Name = s
}


func main() {

    fmt.Println("starting")

    w := &World{}
    w.Set("Mundo")

    display(w)

    //h := &Hello{Name:"Miguel"}
    //fmt.Printf("%v", h)
    //h.Set("Rodrigo")

    //display(h)

}

type entity interface {
    Get() string
    Set(string)
}

func display(e entity)  {
    fmt.Printf("Hello - %s\n", e.Get())
}