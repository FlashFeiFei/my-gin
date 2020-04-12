package help


//如果是小写字母, 则变换为大写字母
func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122  {
		strArry[0] -=  32
	}
	return string(strArry)
}