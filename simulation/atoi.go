package simulation

import "math"

// 实现字符串转整型
// 状态机

type State int
type CharType int

// 定义状态
const (
	STATE_START State = iota
	STATE_SIGN
	STATE_NUMBER
	STATE_END
)

// 字符类型
const (
	CHAR_SPACE CharType = iota
	CHAR_NUMBER
	CHAR_SIGN
	CHAR_OTHER
)

var transfer = map[State]map[CharType]State{
	STATE_START: {
		CHAR_SPACE:  STATE_START,
		CHAR_SIGN:   STATE_SIGN,
		CHAR_NUMBER: STATE_NUMBER,
	},
	STATE_SIGN: {
		CHAR_NUMBER: STATE_NUMBER,
	},
	STATE_NUMBER: {
		CHAR_NUMBER: STATE_NUMBER,
	},
}

func toCharType(ch byte) CharType {
	if ch >= '0' && ch <= '9' {
		return CHAR_NUMBER
	} else if ch == '+' || ch == '-' {
		return CHAR_SIGN
	} else if ch == ' ' {
		return CHAR_SPACE
	} else {
		return CHAR_OTHER
	}
}

func myAtoi(s string) int {
	ans := 0
	sign := 1
	state := STATE_START
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return sign * ans
		} else {
			if transfer[state][typ] == STATE_SIGN && s[i] == '-' {
				sign = -1
			}
			if transfer[state][typ] == STATE_NUMBER {
				digit := int(s[i] - '0')
				if ans > (math.MaxInt32-digit)/10 { // 判断是否越界
					if sign == 1 {
						return math.MaxInt32
					} else {
						return math.MinInt32
					}
				}
				ans = ans*10 + digit
			}
			state = transfer[state][typ] // 下一状态
		}
	}
	return sign * ans
}
