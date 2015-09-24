package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "DoorBuzzerOpen",
        "POST",
        "/door/buzzer/open",
        DoorBuzzerOpen,
    },
    Route{
        "DoorLockLock",
        "POST",
        "/door/lock/lock",
        DoorLockLock,
    },
    Route{
        "DoorLockUnLock",
        "POST",
        "/door/lock/unlock",
        DoorLockUnlock,
    },
    Route{
        "DoorLockLock",
        "POST",
        "/door/lock/open",
        DoorLockOpen,
    },
}
