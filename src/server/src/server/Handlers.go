package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/stianeikeland/go-rpio"
)

func DoPin(pin integer,delay1 integer, delay2 integer) {
	time.Sleep(delay1 * time.Second)
	err := rpio.Open()
	pin := rpio.Pin(pin)
	pin.Output()
	pin.High()
	time.Sleep(delay2 * time.Second)
	pin.Low()
	err.Close()
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Show template")
}

func DoorBuzzerOpen(w http.ResponseWriter, r *http.Request) {
	DoPin(1,15,5)
}

func DoorLockLock(w http.ResponseWriter, r *http.Request) {
	DoPin(2,0,1)
}

func DoorLockUnlock(w http.ResponseWriter, r *http.Request) {
	DoPin(3,0,1)
}

func DoorLockOpen(w http.ResponseWriter, r *http.Request) {
	DoPin(4,0,1)
}
