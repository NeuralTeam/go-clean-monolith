package dotenv

import (
	"bytes"
	"errors"
	"fmt"
	"go-clean-monolith/constants"
	"io"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"syscall"
)

// LoadInEnv read all env and load to ENV for this process
func LoadInEnv(filenames ...string) error {
	filenames = filenamesOrDefault(filenames)

	for _, filename := range filenames {
		err := loadFile(filename, false)
		if err != nil {
			return err
		}
	}

	return nil
}

// LoadInVar read all env and write to 'map[string]string'
func LoadInVar(filenames ...string) (envMap map[string]string, err error) {
	filenames = filenamesOrDefault(filenames)
	envMap = make(map[string]string)

	for _, filename := range filenames {
		individualEnvMap, individualErr := readFile(filename)

		if individualErr != nil {
			err = individualErr
			return
		}

		for key, value := range individualEnvMap {
			envMap[key] = value
		}
	}

	return
}

// LoadInStruct read all env and write to struct
func LoadInStruct(env any, filenames ...string) error {
	filenames = filenamesOrDefault(filenames)
	envDumpMap := make(map[string]string)

	for _, filename := range filenames {
		individualEnvMap, individualErr := readFile(filename)

		if individualErr != nil {
			return individualErr
		}

		for key, value := range individualEnvMap {
			envDumpMap[key] = value
		}
	}

	return validateEnv(env, envDumpMap)
}

// LoadInStructFromENV read process ENV and write to struct
func LoadInStructFromENV(env any) error {
	envMap := make(map[string]string)

	el := reflect.ValueOf(env).Elem()
	for i := 0; i < el.NumField(); i++ {
		key := el.Type().Field(i).Tag.Get("json")
		if value, ok := syscall.Getenv(key); ok {
			envMap[key] = value
		}
	}

	return validateEnv(env, envMap)
}

func validateEnv(env any, envMap map[string]string) error {
	el := reflect.ValueOf(env).Elem()
	for i := 0; i < el.NumField(); i++ {
		f := el.Field(i)
		tf := el.Type().Field(i)

		value, ok := envMap[tf.Tag.Get("json")]

		dotenvTag := tf.Tag.Get("dotenv")
		if dotenvTag != "" {
			opts := strings.Split(dotenvTag, ",")
			for _, opt := range opts {
				optSl := strings.Split(opt, ":")
				if len(optSl) == 2 && optSl[0] == "default" && value == "" {
					value = optSl[1]
					ok = true
					continue
				}
				switch opt {
				case "required":
					if !ok {
						return errors.New(fmt.Sprintf(`var='%s' must be required`, tf.Name))
					}
				case "notnull":
					if value == "" {
						return errors.New(fmt.Sprintf(`var='%s' cant be empty or nil or 0`, tf.Name))
					}
				}
			}
		}

		switch tf.Type.String() {
		case "int":
			tmp, err := strconv.Atoi(value)
			if err != nil {
				return errors.New("can not parse INT value")
			}
			f.Set(reflect.ValueOf(tmp))
		case "string":
			f.SetString(value)
		case "[]string":
			s := strings.Split(value, constants.EnvSeparator)
			for si, v := range s {
				s[si] = strings.Trim(v, " ")
			}
			f.Set(reflect.AppendSlice(f, reflect.ValueOf(s)))
		case "net.IP":
			tmp := net.ParseIP(value)
			if tmp == nil {
				return errors.New("can not parse IP value")
			}
			f.Set(reflect.ValueOf(tmp))
		}
	}
	return nil
}

// parse reads an env file from io.Reader, returning a map of keys and values
func parse(r io.Reader) (map[string]string, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}

	return UnmarshalBytes(buf.Bytes())
}

// Unmarshal reads an env file from a string, returning a map of keys and values
func Unmarshal(str string) (envMap map[string]string, err error) {
	return UnmarshalBytes([]byte(str))
}

// UnmarshalBytes parses env file from byte slice of chars, returning a map of keys and values
func UnmarshalBytes(src []byte) (map[string]string, error) {
	out := make(map[string]string)
	err := parseBytes(src, out)

	return out, err
}

func loadFile(filename string, overload bool) error {
	envMap, err := readFile(filename)
	if err != nil {
		return err
	}

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] || overload {
			_ = os.Setenv(key, value)
		}
	}

	return nil
}

func readFile(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	return parse(file)
}

func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}
