package cli

type boolFlag interface {
	Value
	IsBoolFlag() bool
}
