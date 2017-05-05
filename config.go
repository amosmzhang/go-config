package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func GetConfig(c interface{}, cf string) error {
	if reflect.TypeOf(c).Kind() != reflect.Ptr {
		return errors.New("config.GetConfig() expects a pointer arg")
	}

	// read config file
	var raw []byte
	var err error

	raw, err = ioutil.ReadFile(cf)
	if err != nil {
		return err
	}

	if strings.HasSuffix(cf, ".yml") {
		// unmarshall yaml
		err = yaml.Unmarshal(raw, c)
		if err != nil {
			return err
		}
	} else if strings.HasSuffix(cf, ".json") {
		// unmarshall json
		err = json.Unmarshal(raw, c)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Config file must end in either .yml or .json")
	}

	// read env vars
	err = updateEnvFields(reflect.ValueOf(c), "")
	if err != nil {
		return err
	}

	return nil
}

func updateEnvFields(v reflect.Value, prefix string) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("Not a pointer value")
	}

	v = reflect.Indirect(v)

	switch v.Kind() {
	case reflect.Ptr:
		if err := updateEnvFields(reflect.Indirect(v).Addr(), prefix); err != nil {
			return err
		}
	case reflect.Int:
		if val := os.Getenv(prefix); val != "" {
			conv, err := strconv.Atoi(val)
			if err == nil {
				v.SetInt(int64(conv))
			}
		}
	case reflect.Float64:
		if val := os.Getenv(prefix); val != "" {
			conv, err := strconv.ParseFloat(val, 64)
			if err == nil {
				v.SetFloat(conv)
			}
		}
	case reflect.String:
		if val := os.Getenv(prefix); val != "" {
			v.SetString(val)
		}
	case reflect.Bool:
		if val := os.Getenv(prefix); val != "" {
			conv, err := strconv.ParseBool(val)
			if err == nil {
				v.SetBool(conv)
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			vi := v.Index(i)
			name := strconv.Itoa(i)
			if prefix != "" {
				name = prefix + "_" + name
			}
			if err := updateEnvFields(vi.Addr(), name); err != nil {
				return err
			}
		}
	case reflect.Struct:
		vt := reflect.TypeOf(v.Interface())
		for i := 0; i < vt.NumField(); i++ {
			ft := vt.Field(i)
			fv := v.Field(i)
			name := strings.ToUpper(ft.Name)
			if prefix != "" {
				name = prefix + "_" + name
			}
			if err := updateEnvFields(fv.Addr(), name); err != nil {
				return err
			}
		}
	}

	return nil
}
