package commands

type CallbackFunction func(Context) error

type CheckFunction func(Context) (bool, error)

type Command struct {
	Name     string
	Aliases  []string
	Callback CallbackFunction
	Checks   []CheckFunction
}
