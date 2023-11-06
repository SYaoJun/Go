package mystring

// string 和 []byte的转换

func Reverse_string(s string) string {
	bytes := []byte(s)
	len := len(s)
	for i := 0; i < len/2; i++ {
		var temp = bytes[i]
		bytes[i] = s[len-1-i]
		bytes[len-1-i] = temp
	}
	return string(bytes)
}
