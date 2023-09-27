package memento

func ExampleGame() {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.Status()
	progress := game.Save()

	game.Play(-2, -3)
	game.Status()

	//游戏过程中保存一次
	progress1 := game.Save()
	game.Status()

	//查看进度1的状态
	game.Load(progress1)
	game.Status()

	//查看初始状态
	game.Load(progress)
	game.Status()
	// Output:
	// Current HP:10, MP:10
	// Current HP:7, MP:8
	// Current HP:7, MP:8
	// Current HP:7, MP:8
	// Current HP:10, MP:10
}
