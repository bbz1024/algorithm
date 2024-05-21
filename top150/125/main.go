package main

func main() {
	//1011,0001
	reverseBits(0b10110001)
}
func reverseBits2(n uint8) (rev uint8) {
	for i := 0; i < 8 && n > 0; i++ {
		rev |= n & 1 << (7 - i) //从低位开始写入rev的高位
		n >>= 1                 // 右移
	}
	return
}

func reverseBits(num uint32) (answer uint32) {
	for i := 0; i < 31 && num > 0; i++ {
		answer |= num & 1 << (31 - i)
		num >>= 1
	}
	return answer
}
