package dotenv

import (
	"bytes"
	"errors"
	"go-clean-monolith/constants"
	"io"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// TODO: экспорт ENV в файл .env
// TODO: значения по умолчанию

// LoadInEnv will read your env file(s) and load them into ENV for this process.
func LoadInEnv(filenames ...string) (err error) {
	filenames = filenamesOrDefault(filenames)

	for _, filename := range filenames {
		err = loadFile(filename, false)
		if err != nil {
			return // return early on a spazout
		}
	}
	return
}

// LoadInVar read all env (with same file loading semantics as LoadInEnv) but return values as
// a map rather than automatically writing values into env.
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

func LoadInStruct(env any, filenames ...string) error {
	filenames = filenamesOrDefault(filenames)
	envMap := make(map[string]string)

	for _, filename := range filenames {
		individualEnvMap, individualErr := readFile(filename)

		if individualErr != nil {
			return individualErr
		}

		for key, value := range individualEnvMap {
			envMap[key] = value
		}
	}

	el := reflect.ValueOf(env).Elem()
	for i := 0; i < el.NumField(); i++ {
		f := el.Field(i)
		tf := el.Type().Field(i)

		value, ok := envMap[tf.Tag.Get("json")]
		if !ok {
			continue
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

	return Parse(file)
}

// Parse reads an env file from io.Reader, returning a map of keys and values.
func Parse(r io.Reader) (map[string]string, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}

	return UnmarshalBytes(buf.Bytes())
}

// Unmarshal reads an env file from a string, returning a map of keys and values.
func Unmarshal(str string) (envMap map[string]string, err error) {
	return UnmarshalBytes([]byte(str))
}

// UnmarshalBytes parses env file from byte slice of chars, returning a map of keys and values.
func UnmarshalBytes(src []byte) (map[string]string, error) {
	out := make(map[string]string)
	err := parseBytes(src, out)

	return out, err
}

func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}
