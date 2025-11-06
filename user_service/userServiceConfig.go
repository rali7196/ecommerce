package main

type userServiceConfig struct {
	Salt string
}

var Config userServiceConfig = userServiceConfig{
	Salt: "spiders",
}
