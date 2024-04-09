package hackathon

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

func turnHexa(n int) string {
	hexaChar := "0123456789abcdef"
	base := ""
	for n > 15 {
		base += string(hexaChar[n%16])
		n -= n % 16
		n /= 16
	}
	base += string(hexaChar[n%16])
	return invertString(base)
}

func crypt(sentence string, ranNumber int) string {
	ByteTable := convStringByte(sentence)
	outsideTable := make([]byte, len(sentence))
	finalTable := ByteTable
	finalTable = nand(xor(not(xor(ByteTable, outsideTable)), not(or(and(ByteTable, outsideTable), not(and(ByteTable, outsideTable))))), not(xor(not(and(ByteTable,
		outsideTable)), or(ByteTable, outsideTable))))
	finalValue := totalValueSum(finalTable)
	return turnHexa(finalValue * (len(finalTable) * ranNumber))
}

func FinalCrypt(sentence string, salt int) string {
	inv := invertString(sentence)
	cr := ""
	for i := 0; i < len(sentence); i++ {
		cr += crypt(string(sentence[i]), salt)
		cr += crypt(string(inv[i]), salt)
	}
	return cr
}
