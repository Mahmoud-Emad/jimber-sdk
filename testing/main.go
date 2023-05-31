package main

import (
	jimberSDK "jimber.com/sdk"
)

func main() {
	jimber := jimberSDK.NewJimber("localhost", "8080")
	jimber.RunServer()
}
