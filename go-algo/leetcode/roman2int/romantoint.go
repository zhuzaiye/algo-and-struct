package roman2int

/**
1、一般情况，罗马数字中“小的数字在大的数字的右边”
2、特例就是4，40， 400 和9， 90， 900者六个数字是“小的数字在大的数字的左边”
*/

//数字量很少，采用hash表不会提升多少速度，采用go的switch
func getValue(s byte) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

/**
 1、前值小于后值，减去前值
 2、前值大于或者等于后值，加上前值
 3、最后一个值必须加上
 */
func romanToInt(s string) int {
	rst := 0
	preNumb := getValue(s[0])
	for i:= 1; i<len(s); i++{
		num := getValue(s[i])
		if preNumb < num {
			rst -= preNumb
		}else {
			rst += preNumb
		}
		preNumb = num
	}
	rst += preNumb
	return rst
}
