package tutorial

import (
	"github.com/spf13/pflag"
	"os"
)

type Test2 struct {
	Testvalue string
	flags     *pflag.FlagSet
}

func NewSet2(flags *pflag.FlagSet) *Test2 {
	return &Test2{
		flags: flags,
	}
}

func (t2 *Test2) Parse() error {
	err := t2.flags.Parse(os.Args)
	return err
}

func (t2 *Test2) GetValue() string {

	t2.AddFlags()
	return t2.Testvalue
}

func (t2 *Test2) AddFlags() {
	t2.flags.StringVar(&t2.Testvalue, "testvalue2", "test2", "test value2")
}
