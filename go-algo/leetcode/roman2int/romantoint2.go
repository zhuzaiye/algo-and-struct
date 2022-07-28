package roman2int

func romanToInt2(s string) int {
	if s == ""{
		return 0
	}
	hashMap := map[byte]int{  //这里使用rune，是为了防止数字溢出。byte uint8 而 rune uint32
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	rst := 0
	for i:= 0; i<len(s); i++ {
		if i <len(s)-1 && hashMap[s[i]] < hashMap[s[1+i]]{
			rst -= hashMap[s[i]]
		}else {
			rst += hashMap[s[i]]
		}
	}
	return rst
}
