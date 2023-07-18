package long

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	figs []int
}

func NewFromStr(num string) (Number, error) {
	for _, r := range num {
		if !unicode.IsDigit(r) {
			return Number{}, fmt.Errorf("num must consist from figs")
		}
	}
	figs := make([]int, len(num))
	j := len(figs) - 1
	for i := range num {
		v, _ := strconv.Atoi(num[i : i+1])
		figs[j] = v
		j--
	}
	return Number{figs: figs}, nil
}

func NewFromInt(num int) Number {
	if num == 0 {
		return Number{figs: []int{num}}
	}

	figs := make([]int, 0)
	for num != 0 {
		figs = append(figs, num%10)
		num /= 10
	}
	return Number{figs: figs}
}

func (n Number) String() string {
	sb := strings.Builder{}
	leadZeros := true
	for i := len(n.figs) - 1; i >= 0; i-- {
		if n.figs[i] != 0 {
			leadZeros = false
		}
		if leadZeros {
			continue
		}
		sb.WriteString(strconv.Itoa(n.figs[i]))
	}
	if sb.Len() == 0 {
		sb.WriteString("0")
	}
	return sb.String()
}

func Sum(a, b Number) Number {
	ret := Number{}
	ost := 0

	for i := 0; i < len(a.figs) || i < len(b.figs); i++ {
		var sum int
		if i < len(a.figs) {
			sum += a.figs[i]
		}
		if i < len(b.figs) {
			sum += b.figs[i]
		}
		sum += ost

		ret.figs = append(ret.figs, sum%10)
		ost = sum / 10
	}

	if ost > 0 {
		ret.figs = append(ret.figs, ost)
	}

	return ret
}

func Sub(a, b Number) Number {
	ret := Number{}
	take := 0

	for i := 0; i < len(a.figs); i++ {
		sub := a.figs[i] - take
		if i < len(b.figs) {
			sub -= b.figs[i]
		}

		if sub < 0 {
			sub += 10
			take = 1
		} else {
			take = 0
		}

		ret.figs = append(ret.figs, sub)
	}

	return ret
}

func Mul(a, b Number) Number {
	ret := Number{}

	for i := 0; i < len(a.figs); i++ {
		ost := 0
		t := Number{}

		for j := 0; j < len(b.figs); j++ {
			mul := a.figs[i]*b.figs[j] + ost
			t.figs = append(t.figs, mul%10)
			ost = mul / 10
		}

		if ost > 0 {
			t.figs = append(t.figs, ost)
		}

		for k := 0; k < i; k++ {
			t.figs = append([]int{0}, t.figs...)
		}

		ret = Sum(ret, t)
	}

	return ret
}

func Div(a, b Number) (Number, Number) {
	quotient := Number{}
	dividend := Number{}

	for i := len(a.figs) - 1; i >= 0; i-- {
		dividend.figs = append([]int{a.figs[i]}, dividend.figs...)

		q := 0
		for q = 0; q < 10; q++ {
			temp := Mul(b, NewFromInt(q))
			if Cmp(temp, dividend) > 0 {
				break
			}
		}

		q--
		temp := Mul(b, NewFromInt(q))
		dividend = Sub(dividend, temp)

		quotient.figs = append([]int{q}, quotient.figs...)
	}

	return quotient, dividend
}

func Cmp(a, b Number) int {

	if a.Len() > b.Len() {
		return 1
	} else if a.Len() < b.Len() {
		return -1
	}

	for i := len(a.figs) - 1; i >= 0; i-- {
		if a.figs[i] > b.figs[i] {
			return 1
		} else if a.figs[i] < b.figs[i] {
			return -1
		}
	}

	return 0
}

func (n Number) Len() int {
	for i := len(n.figs) - 1; i >= 0; i-- {
		if n.figs[i] != 0 {
			return i + 1
		}
	}
	return 0
}
