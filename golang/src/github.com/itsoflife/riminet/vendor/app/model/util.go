package model

// stateless utility functions

func CountOnes(num uint32) uint8 {
	count := uint8(0)
	for num > 0 {
		count += 1
		num = num & (num - 1)
	}
	return count
}
