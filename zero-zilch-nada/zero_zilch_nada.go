package zero

// EmptyBool returns an empty (zero value) bool
func EmptyBool() bool {
	var ebool bool
	return ebool
}

// EmptyInt returns an empty (zero value) int
func EmptyInt() int {
	var eint int
	return eint
}

// EmptyString returns an empty (zero value) string
func EmptyString() string {
	var estring string
	return estring
}

// EmptyFunc returns an empty (zero value) func
func EmptyFunc() func() {
	var efunc func()
	return efunc
}

// EmptyPointer returns an empty (zero value) pointer
func EmptyPointer() *int {
	var epint *int
	return epint
}

// EmptyMap returns an empty (zero value) map
func EmptyMap() map[int]int {
	var emap map[int]int
	return emap
}

// EmptySlice returns an empty (zero value) slice
func EmptySlice() []int {
	var eslice []int
	return eslice
}

// EmptyChannel returns an empty (zero value) channel
func EmptyChannel() chan int {
	var ch chan int
	return ch
}

// EmptyInterface returns an empty (zero value) interface
func EmptyInterface() interface{} {
	var einter interface{}
	return einter
}
