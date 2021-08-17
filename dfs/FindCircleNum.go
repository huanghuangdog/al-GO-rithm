package dfs

/**
20210817
 */
func findCircleNum(isConnected [][]int) (rs int) {
	size := len(isConnected)
	isVisited := make([]int, size)
	var dfs func(index int)
	dfs = func(index int) {
		isVisited[index] = 1
		for i:= 0;i < size;i++ {
			if isVisited[i] == 1 || i == index || isConnected[index][i] == 0 {
				continue
			}
			dfs(i)
		}
	}
	for i:= 0;i < size;i++ {
		if isVisited[i] == 0 {
			dfs(i)
			rs++
		}
	}
	return
}
