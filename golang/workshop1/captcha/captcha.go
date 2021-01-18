package captcha

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Captcha struct
type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

// New Captcha
func New(pattern, leftOperand, operator, rightOperand int) Captcha {
	return Captcha{pattern, leftOperand, operator, rightOperand}
}

func (cc *Captcha) String() string {
	operands := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nice"}
	operators := []string{"+", "-", "*"}

	if cc.pattern == 1 {
		return fmt.Sprintf("%v %v %v", cc.leftOperand, operators[cc.operator-1], operands[cc.rightOperand])
	} else {
		return fmt.Sprintf("%v %v %v", operands[cc.leftOperand], operators[cc.operator-1], cc.rightOperand)
	}
}

var src = rand.NewSource(time.Now().UnixNano())
var rdn = rand.New(src)
var stores = map[string]int{}

// Question random captha
func Question() (string, string) {
	pt, lo, op, ro := rdn.Intn(2)+1, rdn.Intn(9)+1, rdn.Intn(3)+1, rdn.Intn(9)+1
	cc := New(pt, lo, op, ro)
	ans := 0

	switch op {
	case 1:
		ans = lo + ro
	case 2:
		ans = lo - ro
	case 3:
		ans = lo * ro
	}

	key := uuid.New().String()
	stores[key] = ans

	log.Printf("new key %v", key)
	log.Printf("set val %v", ans)

	return key, cc.String()
}

// Answer question
func Answer(key string, ans int) bool {

	if val, ok := stores[key]; ok {
		log.Printf("val: %v\n", val)
		if val == ans {
			delete(stores, key)
			return true
		}
		return false
	}

	log.Printf("key not found %v\n", key)

	return false
}
