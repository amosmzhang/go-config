package config

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	First  string `yaml:"first"`
	Second string `yaml:"second"`
	Third  string `yaml:"third"`
}

type TestConfig2 struct {
	First  string  `yaml:"first"`
	Second string  `yaml:"second"`
	Third  int     `yaml:"third"`
	Fourth bool    `yaml:"fourth"`
	Fifth  float64 `yaml:"fifth"`
}

type TestSubStruct struct {
	A string `yaml:"a"`
	B int    `yaml:"b"`
}

type TestConfig3 struct {
	First  []string      `yaml:"first"`
	Second []string      `yaml:"second"`
	Third  TestSubStruct `yaml:"third"`
}

type TestConfig4 struct {
	Tough []TestSubStruct `yaml:"tough"`
}

type TestJSON struct {
	A string `json:"a"`
	B int    `json:"b"`
}

type TestPointer struct {
	First  []string       `yaml:"first"`
	Second []string       `yaml:"second"`
	Third  *TestSubStruct `yaml:"third"`
}

func TestReadYMLConfig(t *testing.T) {
	var c TestConfig

	err := GetConfig(&c, "config_test.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "configItem1", c.First)
	assert.Equal(t, "configItem2", c.Second)
}

func TestReadJSONConfig(t *testing.T) {
	var c TestJSON

	err := GetConfig(&c, "config_test.json")

	assert.Equal(t, nil, err)
	assert.Equal(t, "a string", c.A)
	assert.Equal(t, 1234, c.B)
}

func TestReadConfigEnv(t *testing.T) {
	var c TestConfig

	os.Setenv("SECOND", "osConfig2")
	os.Setenv("THIRD", "osConfig3")

	err := GetConfig(&c, "config_test.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "configItem1", c.First)
	assert.Equal(t, "osConfig2", c.Second)
	assert.Equal(t, "osConfig3", c.Third)

	os.Clearenv()
}

func TestReadConfig2(t *testing.T) {
	var c TestConfig2

	err := GetConfig(&c, "config_test2.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "configItem1", c.First)
	assert.Equal(t, "configItem2", c.Second)
	assert.Equal(t, 3, c.Third)
	assert.Equal(t, true, c.Fourth)
	assert.Equal(t, 5.5, c.Fifth)
}

func TestReadConfigEnv2(t *testing.T) {
	var c TestConfig2

	os.Setenv("SECOND", "osConfig2")
	os.Setenv("THIRD", "333")
	os.Setenv("FOURTH", "false")
	os.Setenv("FIFTH", "5.0505")

	err := GetConfig(&c, "config_test2.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "configItem1", c.First)
	assert.Equal(t, "osConfig2", c.Second)
	assert.Equal(t, 333, c.Third)
	assert.Equal(t, false, c.Fourth)
	assert.Equal(t, 5.0505, c.Fifth)

	os.Clearenv()
}

func TestReadConfig3(t *testing.T) {
	var c TestConfig3

	err := GetConfig(&c, "config_test3.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "item1", c.First[0])
	assert.Equal(t, "item2", c.First[1])
	assert.Equal(t, "subitem1", c.Third.A)
	assert.Equal(t, 2, c.Third.B)
}

func TestReadConfigEnv3(t *testing.T) {
	var c TestConfig3

	os.Setenv("FIRST_1", "osItem1")
	os.Setenv("THIRD_A", "osSubItem1")

	err := GetConfig(&c, "config_test3.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "item1", c.First[0])
	assert.Equal(t, "osItem1", c.First[1])
	assert.Equal(t, "osSubItem1", c.Third.A)

	os.Clearenv()
}

func TestReadConfig4(t *testing.T) {
	var c TestConfig4

	err := GetConfig(&c, "config_test4.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "subitem1", c.Tough[0].A)
	assert.Equal(t, 2, c.Tough[1].B)
}

func TestReadConfigEnv4(t *testing.T) {
	var c TestConfig4

	os.Setenv("TOUGH_0_A", "ositem1")
	os.Setenv("TOUGH_1_B", "22")

	err := GetConfig(&c, "config_test4.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "ositem1", c.Tough[0].A)
	assert.Equal(t, 22, c.Tough[1].B)

	os.Clearenv()
}

func TestReadConfig3Pointer(t *testing.T) {
	var c TestPointer

	err := GetConfig(&c, "config_test3.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "item1", c.First[0])
	assert.Equal(t, "item2", c.First[1])
	assert.Equal(t, "subitem1", c.Third.A)
	assert.Equal(t, 2, c.Third.B)
}

func TestReadConfigEnv3Pointer(t *testing.T) {
	var c TestPointer

	os.Setenv("FIRST_1", "osItem1")
	os.Setenv("THIRD_A", "osSubItem1")

	err := GetConfig(&c, "config_test3.yml")

	assert.Equal(t, nil, err)
	assert.Equal(t, "item1", c.First[0])
	assert.Equal(t, "osItem1", c.First[1])
	assert.Equal(t, "osSubItem1", c.Third.A)

	os.Clearenv()
}
