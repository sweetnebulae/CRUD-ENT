package helper

// ErrorPanic is a helper function that panics if the given error is not nil.
// It is typically used in low-level or startup code where failure is unrecoverable.
//
// Example usage:
//
//	file, err := os.Open("config.json")
//	helper.ErrorPanic(err)
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
