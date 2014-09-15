package cli

import "sort"
import "errors"
import "fmt"

var ErrHelp = errors.New("flag: help requested")

// sortFlags returns the flags as a slice in lexicographical sorted order.
func sortFlags(flags map[string]*Flag) []*Flag {
  list := make(sort.StringSlice, len(flags))
  i := 0
  for _, f := range flags {
    list[i] = f.Name
    i++
  }
  list.Sort()
  result := make([]*Flag, len(list))
  for i, name := range list {
    result[i] = flags[name]
  }
  return result
}

// defaultUsage is the default function to print a usage message.
func defaultUsage(f *FlagSet) {
  if f.name == "" {
    fmt.Fprintf(f.out(), "Usage:\n")
  } else {
    fmt.Fprintf(f.out(), "Usage of %s:\n", f.name)
  }
  f.PrintDefaults()
}
