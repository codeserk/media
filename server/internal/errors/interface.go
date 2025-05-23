package errors

// Error interface that is compatible with Error GoLang interface.
// The motivation for this interface is being able to determine what part of
// the message should go the clients, and which part is internal (and should
// be logged, but not exposed in any place).
// Using this interface should prevent sensitive information to be exposed.
type Interface interface {
	Error() string

	Public() string
	Internal() string
}
