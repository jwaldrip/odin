package cli

type boolFlag interface {
  FlagValue
  IsBoolFlag() bool
}
