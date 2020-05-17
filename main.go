package main

func main() {
	cfg := createConfig("assets/config.ini")

	game := createEngine(cfg)
	game.init()
	game.run()
}