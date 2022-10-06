package tutorial

import (
	"github.com/spf13/pflag"
	"os"
)

type Test struct {
	Testvalue string
	flags     *pflag.FlagSet
}

func NewSet(flags *pflag.FlagSet) *Test {
	return &Test{
		flags: flags,
	}
}

func (t *Test) Parse() error {
	err := t.flags.Parse(os.Args)
	return err
}

func (t *Test) GetValue() string {

	t.AddFlags()
	t.Parse()
	return t.Testvalue
}

func (t *Test) AddFlags() {
	t.flags.StringVar(&t.Testvalue, "testvalue", "test", "test value")
}
