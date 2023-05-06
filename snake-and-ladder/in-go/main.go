package main

import "snake-and-ladder/in-go/engine"

func main() {

	gameEngine := engine.InitEngine(10, 10, 200)
	gameEngine.AddPlayer("Naruto")
	gameEngine.AddPlayer("Sasuke")
	gameEngine.AddPlayer("Sakura")
	gameEngine.AddPlayer("Negi")
	gameEngine.Play()

}
