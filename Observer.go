package main

type Observer interface {
	handleEvent([]string)
}
