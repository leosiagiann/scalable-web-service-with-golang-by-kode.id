package main

import "lat-http-request/routers"

func main() {
	PORT := ":9000"

	routers.StartServer().Run(PORT)
}
