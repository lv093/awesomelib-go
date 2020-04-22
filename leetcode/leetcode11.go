package main

func main() {
	target := []int{1,8,6,2,5,4,8,3,7}
	println(maxArea(target))
}


func maxArea(height []int) int {
	mlen := len(height)

	max := 0
	for i:=0; i < mlen; i++ {
		for j := i+1; j < mlen; j++ {
			h := height[i]
			if height[j] < h {
				h = height[j]
			}
			temp := (j - i) * h
			if temp > max {
				max = temp
			}
		}
	}
	return max
}
