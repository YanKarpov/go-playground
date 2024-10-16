package main

type State interface {
	Handle() string
}
