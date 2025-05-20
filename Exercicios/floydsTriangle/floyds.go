package main

func floydsTriangle(n int) [][]int {
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangle[i][j] = 1
			} else {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}
	return triangle

}

func main() {
	triangle := floydsTriangle(5)
	print(triangle)
}
