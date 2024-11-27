package math

func IsPrime(num int64) bool {
	if num == 1 {
		return true
	}
	var i int64
	for i = 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
