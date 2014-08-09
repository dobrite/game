package game

type system interface {
	enqueue(message)
	run()
}
