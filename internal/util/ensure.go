package util

// Ensure provides a wrapper for validating conditions that should
// always evaluate to true. An example is initializing an database connection.
func Ensure(condition bool, message string) {
	if !condition {
		panic(message)
	}
}
