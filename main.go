package main

import (
    "vidar.sh/server"
)

func main() {
    // Start server
    svr := server.StartServer()

    if err := srv.Start(); err != nil {
        srv.Logger.Fatal("Error starting server: ", err)
    }
}
