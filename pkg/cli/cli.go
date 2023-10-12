package cli

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var ParsingTypeError = errors.New("cannot getting value")

type Flag struct {
	Name    string
	Default interface{}

	value interface{}
}

func (f *Flag) Bool() (bool, error) {
	if v, ok := f.value.(bool); ok {
		return v, nil
	}
	return false, ParsingTypeError
}

func (f *Flag) String() (string, error) {
	if v, ok := f.value.(string); ok {
		return v, nil
	}
	return "", ParsingTypeError
}

type Command struct {
	Name  string
	Short string
	Long  string
	Run   func()

	regFlags map[string]*Flag
}

func (c *Command) Env(name, value string) {
	if c.regFlags == nil {
		c.regFlags = make(map[string]*Flag)
	}
	c.regFlags[name] = &Flag{Name: name, Default: value}
}

func (c *Command) EnvBool(name string) {
	if c.regFlags == nil {
		c.regFlags = make(map[string]*Flag)
	}
	c.regFlags[name] = &Flag{Name: name, Default: false}
}

func (c *Command) EnvMulti(name string, values ...string) {
	if c.regFlags == nil {
		c.regFlags = make(map[string]*Flag)
	}
	c.regFlags[name] = &Flag{Name: name, Default: values}
}

func (c *Command) parseFlags(args []string) {
	for i := 0; i < len(args); i++ {
		hasPrefix := strings.HasPrefix(args[i], "--")
		if hasPrefix && i < len(args)-1 && !strings.HasPrefix(args[i+1], "--") {
			flag, ok := c.regFlags[args[i][2:]]
			if !ok {
				log.Fatalf("recognize flag error: '%s'", args[i][2:])
			}
			t, ok := flag.Default.([]string)
			if ok {
				for _, s := range t {
					if s == args[i+1] {
						break
					}
					log.Fatalf("multi-flag value error: '%s' not in [%s]\n", args[i+1], t)
				}
			}
			flag.value = args[i+1]
		} else if hasPrefix {
			flag, ok := c.regFlags[args[i][2:]]
			if !ok {
				log.Fatalf("recognize flag error: '%s'", args[i][2:])
			}
			flag.value = true
		}
	}
}

func Parse(cmds []*Command) {
	var (
		active *Command
	)
	//start --host 127.0.0.1 --port 8000 --methods GET
	if len(os.Args) < 2 {
		log.Fatalln("command error: enter a valid command")
	}

	for _, cmd := range cmds {
		if cmd.Name == os.Args[1] {
			active = cmd
			break
		}
	}
	if active == nil {
		log.Fatalln("command error: enter a valid command")
	}

	active.parseFlags(os.Args[2:])

	for _, flag := range active.regFlags {
		fmt.Println(flag)
		//err := os.Setenv(flag.Name, flag.value.(string))
		//if err != nil {
		//	fmt.Println(err)
		//}
	}

	active.Run()
}
