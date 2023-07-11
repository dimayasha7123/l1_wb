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

func NewFromStr(num string) (*Number, error) {
	for _, r := range num {
		if !unicode.IsDigit(r) {
			return nil, fmt.Errorf("num must consist from digits")
		}
	}
	figs := make([]int, 0, len(num))
	for i := range num {
		v, _ := strconv.Atoi(num[i : i+1])
		figs = append(figs, v)
	}
	for i, j := 0, len(figs)-1; i < j; i, j = i+1, j-1 {
		figs[i], figs[j] = figs[j], figs[i]
	}
	return &Number{figs: figs}, nil
}

func NewFromInt(num int) *Number {
	if num == 0 {
		return &Number{figs: []int{num}}
	}

	figs := make([]int, 0)
	for num != 0 {
		figs = append(figs, num%10)
		num /= 10
	}
	return &Number{figs: figs}
}

func (n *Number) String() string {
	sb := strings.Builder{}
	for i := len(n.figs) - 1; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(n.figs[i]))
	}
	return sb.String()
}

func Sum(a, b *Number) *Number {
	ret := &Number{}
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

func Sub(a, b *Number) *Number {
	ret := &Number{}
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

func Mul(a, b *Number) *Number {
	ret := &Number{}

	for i := 0; i < len(a.figs); i++ {
		ost := 0
		t := &Number{}

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
