package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("before")
	ret := m.Run()
	log.Println("after")
	os.Exit(ret)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
