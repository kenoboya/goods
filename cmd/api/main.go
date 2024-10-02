package main

import "goods/internal/api"

const configDIR = "api/configs"
const envDIR = "api/.env"

func main() {
	api.Run(configDIR, envDIR)
}
