package hackathon

import (
	"bytes"
	"math/rand"
	"strconv"
)

func convStringByte(sentence string) []byte {
	byteArray := []byte(sentence)
	return byteArray
}

func totalValueSum(table []byte) int {
	val := 1
	for _, i := range table {
		val += int(i)
	}
	return val
}

func xor(t1 []byte, t2 []byte) []byte {
	var t3 []byte
	for i := 0; i < len(t1); i++ {
		t3 = append(t3, t1[i]^t2[i])
	}
	return t3
}

func and(t1 []byte, t2 []byte) []byte {
	var t3 []byte
	for i := 0; i < len(t1); i++ {
		t3 = append(t3, t1[i]&t2[i])
	}
	return t3
}

func nand(t1 []byte, t2 []byte) []byte {
	var t3 []byte
	for i := 0; i < len(t1); i++ {
		t3 = append(t3, t1[i]&^t2[i])
	}
	return t3
}

func or(t1 []byte, t2 []byte) []byte {
	var t3 []byte
	for i := 0; i < len(t1); i++ {
		t3 = append(t3, t1[i]|t2[i])
	}
	return t3
}

func not(t []byte) []byte {
	var t2 []byte
	for i := 0; i < len(t); i++ {
		t2 = append(t2, ^t[i])
	}
	return t2
}

func shiftBytesLeft(a []byte) (dst []byte) {
	n := len(a)
	dst = make([]byte, n)
	for i := 0; i < n-1; i++ {
		dst[i] = a[i] << 1
		dst[i] = (dst[i] & 0xfe) | (a[i+1] >> 7)
	}
	dst[n-1] = a[n-1] << 1
	return dst
}

func operator(t []byte) ([]byte, []int) {
	var tAll [][]byte
	var tAllCopy [][]byte
	var tFinal []byte
	tCopy := t
	tEmpty := make([]byte, len(t))
	for !bytes.Equal(tCopy, tEmpty) {
		tAll = append(tAll, tCopy)
		tCopy = shiftBytesLeft(tCopy)
	}
	var operandsOrder []int
	for i := 0; i < len(t); i++ {
		operandsOrder = append(operandsOrder, rand.Intn(5-1)+1)
	}
	for i, o := range operandsOrder {
		switch o {
		case 1:
			if tAllCopy == nil {
				tAllCopy = append(tAllCopy, nand(tAll[i], tAll[i+1]))
			} else {
				tAllCopy = append(tAllCopy, nand(tAllCopy[i-1], tAll[i]))
			}
		case 2:
			if tAllCopy == nil {
				tAllCopy = append(tAllCopy, xor(tAll[i], tAll[i+1]))
			} else {
				tAllCopy = append(tAllCopy, xor(tAllCopy[i-1], tAll[i]))
			}
		case 3:
			if tAllCopy == nil {
				tAllCopy = append(tAllCopy, and(tAll[i], tAll[i+1]))
			} else {
				tAllCopy = append(tAllCopy, and(tAllCopy[i-1], tAll[i]))
			}
		case 4:
			if tAllCopy == nil {
				tAllCopy = append(tAllCopy, or(tAll[i], tAll[i+1]))
			} else {
				tAllCopy = append(tAllCopy, or(tAllCopy[i-1], tAll[i]))
			}
		}
	}
	tFinal = tAllCopy[len(tAllCopy)-1]
	if rand.Intn(2-1)+1 == 1 {
		return tFinal, operandsOrder
	} else {
		return not(tFinal), operandsOrder
	}
}

func AllForOne(hexa string) string {
	var hexas []string
	s := ""
	for _, c := range hexa {
		s += string(c)
		if len(s) == 15 {
			hexas = append(hexas, s)
			s = ""
		}
	}
	var ints []int64
	for _, h := range hexas {
		if HexaToInt(h) == 0 {
			return ""
		}
		ints = append(ints, HexaToInt(h))
	}
	var average int64 = 0
	for _, n := range ints {
		average += n
	}
	average /= int64(len(ints))
	return strconv.FormatInt(average, 16)
}

func valueCalc(t []byte, table []int) (string, int) {
	tVal := totalValueSum(t)
	value := 1
	diff := len(table)
	hexaFinal := ""
	for i := 0; i < len(table); i++ {
		value *= (table[i] * (10 * diff))
		diff -= 1
	}
	hexaFinal += strconv.FormatInt(int64(tVal*value*len(t)), 16)
	for i := 1; i < len(t); i++ {
		byyte := convStringByte(RandStringRunes(len(t)))
		tVal = totalValueSum(byyte)
		hexaFinal += strconv.FormatInt(int64(tVal*value*len(t)), 16)
	}
	hexaFinal = AllForOne(hexaFinal)
	return hexaFinal, value
}

func HexaToInt(hexa string) int64 {
	value, err := strconv.ParseInt(hexa, 16, 64)
	if err != nil {
		return 0
	}
	return value
}

func RandStringRunes(n int) string {
	letterRunes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMMNOPQRSTUVWXYZ0123456789&+/*%$!;:.,?~"
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(letterRunes[rand.Intn(len(letterRunes))])
	}
	return string(b)
}

func Crypt(s string) string {
	cryptedBytes, operators := operator(convStringByte(s))
	value, operator := valueCalc(cryptedBytes, operators)
	if value == "" {
		return Crypt(s)
	}
	return strconv.FormatInt(int64(operator), 16) + value
}
