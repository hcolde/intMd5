/*
*Auth:hcolde
*Date:2019/12/24
*Getting MD5 as number instead of string.
*/
package intMd5

type MD5 struct {
	md51 string
	md52 string
	temp int
}

func Str2Dec(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}

func Str2Int(i int32) (string, string) {
	if i > '9' {
		return string(i - int32(16)), "1"
	} else {
		return string(i), "0"
	}
}

func Md52int(s string) (b string, t string) {
	for _, j := range s {
		m, n := Str2Int(j)
		b += m
		t += n
	}
	return
}

func GetMd5(s string) (md5 MD5) {
	m, n := Md52int(s)
	md5.md51 = m[:16]
	md5.md52 = m[16:]
	md5.temp = Str2Dec(n)
	return
}

func Reset(md5 MD5) (m string) {
	temp := md5.temp
	s := ""
	for temp > 1 {
		s = string(temp % 2 + 48) + s
		temp /= 2
	}
	s = "1" + s

	for i, j := range md5.md51 {
		if s[i] == '1' {
			m += string(j + 16)
		} else {
			m += string(j)
		}
	}

	for i, j := range md5.md52 {
		if s[i+16] == '1' {
			m += string(j + 16)
		} else {
			m += string(j)
		}
	}
	return
}