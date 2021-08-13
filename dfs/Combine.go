package dfs

func combine(n int, k int) (rs [][]int) {
	var dfs func()
	tmp := make([]int, 0)
	m := make(map[int]struct{})
	dfs = func() {
		for i:= 1;i <= n;i++ {
			if len(tmp) == k {
				rs = append(rs, append([]int(nil), tmp...))
				return
			}
			if _,ok := m[i];ok {
				continue
			}
			if len(tmp) != 0 && i < tmp[len(tmp) - 1]{
				continue
			}
			tmp = append(tmp, i)
			m[i] = struct{}{}
			dfs()
			tmp = tmp[:len(tmp) - 1]
			delete(m, i)
		}
	}
	dfs()
	return
}