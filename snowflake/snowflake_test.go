package snowflake

import (
	"fmt"
	localTime "github.com/go-estar/local-time"
	"testing"
)

func TestGetEpoch(t *testing.T) {
	ti, _ := localTime.ParseLocal("2020-01-01 00:00:00")
	fmt.Println(ti.UnixMillisecond())
}

func TestLen15(t *testing.T) {
	var generator = New(1, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 3,
		StepBits: 7,
	})
	id := generator.String()
	fmt.Println(id, len(id))
}

func TestLen16(t *testing.T) {
	var generator = New(1, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 4,
		StepBits: 9,
	})
	id := generator.String()
	fmt.Println(id, len(id))
}

func TestLen19(t *testing.T) {
	var generator = New(1, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 10,
		StepBits: 12,
	})
	id := generator.String()
	fmt.Println(id, len(id))
}
