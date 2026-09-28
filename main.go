package main

func main() {
	var cfg config
	cfg.commands = getCliCommands()
	startRepl(&cfg)

}
