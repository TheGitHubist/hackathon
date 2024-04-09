package hackathon

import (
	"fmt"
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

func invertString(s string) string {
	newS := ""
	for s != "" {
		newS += string(s[len(s)-1])
		s = s[:len(s)-1]
	}
	return newS
}

func power(n int, p int) int {
	pow := 1
	for i := 0; i < p; i++ {
		pow *= n
	}
	return pow
}

func crypt(sentence string, ranNumber int) string {
	ByteTable := convStringByte(sentence)
	outsideTable := make([]byte, len(sentence))
	finalTable := ByteTable
	finalTable = nand(xor(not(xor(ByteTable, outsideTable)), not(or(and(ByteTable, outsideTable), not(and(ByteTable, outsideTable))))), not(xor(not(and(ByteTable, outsideTable)), or(ByteTable, outsideTable))))
	finalValue := totalValueSum(finalTable)
	return strconv.FormatInt(int64(finalValue*(len(finalTable)*power(ranNumber, 4))), 16)
}

func FinalCrypt(sentence string) string {
	inv := invertString(sentence)
	cr := ""
	ran := rand.Intn(999999999-999999) + 999999
	fmt.Println(ran)
	for i := 0; i < len(sentence); i++ {
		cr += crypt(string(sentence[i]), ran)
		cr += crypt(string(inv[i]), ran)
	}
	return gatherHexaToOne(cr)
}

func averageHexa(hexa1 string, hexa2 string) string {
	hexaChar := "0123456789abcdef"
	newHexa := ""
	for i, c := range hexa1 {
		ch1 := 0
		ch2 := 0
		for j := 0; j < len(hexaChar); j++ {
			if hexaChar[j] == byte(c) {
				ch1 = j
			} else if i < len(hexa2) && hexaChar[j] == hexa2[i] {
				ch2 = j
			}
		}
		newHexa += strconv.FormatInt(int64((ch1+ch2)/2), 16)
	}
	return newHexa
}

func gatherHexaToOne(hexa string) string {
	table := []string{}
	sign := ""
	for i, c := range hexa {
		sign += string(c)
		if len(sign) == 64 {
			table = append(table, sign)
			sign = ""
		} else if i == len(hexa)-1 {
			table = append(table, sign)
			sign = ""
		}
	}
	newHexa := ""
	if len(table) == 1 {
		return table[0]
	}
	for len(table) > 0 {
		if newHexa == "" {
			newHexa += averageHexa(table[len(table)-1], table[len(table)-2])
			table = table[:len(table)-1]
			table = table[:len(table)-1]
		} else {
			if len(table[len(table)-1]) < 64 {
				break
			}
			newHexa += averageHexa(newHexa, table[len(table)-1])
			table = table[:len(table)-1]
		}
	}
	return newHexa
}
