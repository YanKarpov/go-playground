package main

type Handler interface {
    SetNext(handler Handler) Handler
    UseItem(item string)
}
