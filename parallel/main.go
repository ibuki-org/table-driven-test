package parallel

import "time"

func main() {
	Add(1, 2)
}

func Add(a, b int) int {
	result := a + b
	time.Sleep(3 * time.Second)
	return result
}
