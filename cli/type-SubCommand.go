package cli

type SubCommand struct {
  CLI
  parents []Command
}
