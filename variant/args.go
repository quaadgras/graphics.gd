package variant

// Arguments efficiently represents a variable number of arguments, for use in trampolines.
type Arguments struct{}

// Get the nth argument from the arguments list.
func (args Arguments) Get(index int) Any {
	return Any{}
}
