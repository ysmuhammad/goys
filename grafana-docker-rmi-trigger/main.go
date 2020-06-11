package main

import (
        "encoding/json"
        "os/exec"
        "log"
        "fmt"
        "net/http"
)

type Message struct {
        Id int64
        State string `json:"state"`
}

func Cleaner(w http.ResponseWriter, r *http.Request) {
        b := json.NewDecoder(r.Body)
        defer r.Body.Close()

        var msg Message
        err := b.Decode(&msg)

        fmt.Println(msg.State)
        if msg.State == "alerting" {
                fmt.Println("MAYDAY MAYDAY !!")
                exec.Command("/bin/sh","rm-docker-image.sh").Run()
        } else {
                fmt.Println("Chill")
        }

        output, err := json.Marshal(msg)
        if err != nil {
                http.Error(w, err.Error(), 500)
                return
        }
        w.Header().Set("content-type", "application/json")
        w.Write(output)
}

func main() {
        http.HandleFunc("/", Cleaner)
        address := ":8000"
        log.Println("Starting server on address", address)
        err := http.ListenAndServe(address, nil)
        if err != nil {
                panic(err)
        }
}