package swordtech

var initFunc []func()

func init() {
	initFunc = make([]func(), 0)
}

// AddStartup adds a function to the startup
func AddStartup(s func()) {
	initFunc = append(initFunc, s)
}

func executeStartup() {
	for _, f := range initFunc {
		f()
	}
}
