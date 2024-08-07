package stringutil

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

// Creates an slice of slice values not included in the other given slice.
func Diff(base, exclude []string) (result []string) {
	excludeMap := make(map[string]bool)
	for _, s := range exclude {
		excludeMap[s] = true
	}
	for _, s := range base {
		if !excludeMap[s] {
			result = append(result, s)
		}
	}
	return result
}

func Unique(ss []string) (result []string) {
	smap := make(map[string]bool)
	for _, s := range ss {
		smap[s] = true
	}
	for s := range smap {
		result = append(result, s)
	}
	return result
}

func StringFind(array []string, str string) int {
	if str == "" {
		return -1
	}
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}

func StringIn(str string, array []string) bool {
	return StringFind(array, str) > -1
}

func StringReverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func StringJoin(array ...string) string {
	buf := new(strings.Builder)
	for _, s := range array {
		buf.WriteString(s)
	}
	return buf.String()
}

// example: prefix:"[", suffix:"]", separator:",", slice:[]string{"a", "b", "c"} => [a],[b],[c]
func StringJoinWithOvercoat(prefix string, suffix string, separator string, slice ...string) string {
	ss := make([]string, 0, len(slice))
	for _, s := range slice {
		ss = append(ss, fmt.Sprintf(`%s%s%s`, prefix, s, suffix))
	}
	return strings.Join(ss, separator)
}

func NewString(v string) *string {
	return &v
}

func SimplifyStringList(s []string) []string {
	b := s[:0]
	for _, x := range s {
		if x := SimplifyString(x); x != "" {
			b = append(b, x)
		}
	}
	return b
}

var reMoreSpace = regexp.MustCompile(`\s+`)

// "\ta  b  c" => "a b c"
func SimplifyString(s string) string {
	return reMoreSpace.ReplaceAllString(strings.TrimSpace(s), " ")
}

// 生成数字验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
