/*
*Auth:hcolde
*Date:2019/12/24
*Getting MD5 as number instead of string.
*/
package intMd5

import (
	"fmt"
	"errors"
	"strconv"
)

type MD5 struct {
	Md51 int
	Md52 int
	Temp int
}

func Str2Dec(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}

func Str2Int(i int32) (string, string, error) {
	if i & 0x0a > 0x09 {
		return "", "", errors.New("Characters must be between a and i!")
	}
	switch i >> 4 {
		case 3:
			return string(i), "0", nil
		case 4:
			return string(i - int32(16)), "1", nil
		case 6:
			return string(i - int32(48)), "1", nil
		default:
			fmt.Println(i, i>>4)
			return "", "", errors.New("Characters must be between a and i!")
	}
}

func Md52int(s string) (string, string, error) {
	b := ""
	t := ""
	for _, j := range s {
		m, n, err := Str2Int(j)
		if err != nil {
			return "", "", err
		}
		b += m
		t += n
	}
	return b, t, nil
}

func GetMd5(s string) (md5 MD5, err error) {
	if len(s) != 32 {
		err = errors.New("Please pass in a 32-char string!")
		return
	}
	m, n, err := Md52int(s)
	if err != nil {
		return
	}

	md5.Md51, err = strconv.Atoi(m[:16])
	if err != nil {
		return
	}

	md5.Md52, err = strconv.Atoi(m[16:])
	if err != nil {
		return
	}
	md5.Temp = Str2Dec(n)
	return
}

func Reset(md5 MD5) (m string) {
	temp := md5.Temp
	s := ""
	for temp > 1 {
		s = string(temp % 2 + 48) + s
		temp /= 2
	}
	s = "1" + s

	for i, j := range strconv.Itoa(md5.Md51) {
		if s[i] == '1' {
			m += string(j + 16)
		} else {
			m += string(j)
		}
	}

	for i, j := range strconv.Itoa(md5.Md52) {
		if s[i+16] == '1' {
			m += string(j + 16)
		} else {
			m += string(j)
		}
	}
	return
}