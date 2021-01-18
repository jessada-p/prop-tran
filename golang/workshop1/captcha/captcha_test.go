package captcha_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pallat/todos/captcha"
)

func TestCaptchaPattern1(t *testing.T) {

	operands := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nice"}
	operators := []string{"+", "-", "*"}

	for opnIndex, opnValue := range operands {
		for opIndex, opVal := range operators {
			t.Run(fmt.Sprintf("operands %d operator %v", opnIndex, opVal), func(t *testing.T) {
				pt := 1
				lo := 1
				op := opIndex + 1
				ro := opnIndex
				want := fmt.Sprintf("1 %v %v", opVal, opnValue)

				cc := captcha.New(pt, lo, op, ro)
				get := cc.String()
				assert.Equal(t, get, want, fmt.Sprintf("given %d %d %d %d want %q but get %q", pt, lo, op, ro, want, get))
			})
		}
	}
}
func TestCaptchaPattern2(t *testing.T) {

	operands := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nice"}
	operators := []string{"+", "-", "*"}

	for opnIndex, opnValue := range operands {
		for opIndex, opVal := range operators {
			t.Run(fmt.Sprintf("operands %d operator %v", opnIndex, opVal), func(t *testing.T) {
				pt := 2
				lo := opnIndex
				op := opIndex + 1
				ro := 1
				want := fmt.Sprintf("%v %v 1", opnValue, opVal)

				cc := captcha.New(pt, lo, op, ro)
				get := cc.String()
				assert.Equal(t, get, want, fmt.Sprintf("given %d %d %d %d want %q but get %q", pt, lo, op, ro, want, get))
			})
		}
	}
}
