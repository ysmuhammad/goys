import "sort"

func solution(dishes [][]string) [][]string {
	var res [][]string
	var keys []string
	m := make(map[string][]string)
	for i := 0; i < len(dishes); i++ {
		for j := 1; j < len(dishes[i]); j++ {
			m[dishes[i][j]] = append(m[dishes[i][j]], dishes[i][0])
			sort.Strings(m[dishes[i][j]])
		}
	}
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for z := 0; z < len(keys); z++ {
		if len(m[keys[z]]) >= 2 {
			var final []string
			final = append(final, keys[z])
			for _, f := range m[keys[z]] {
				final = append(final, f)
			}
			res = append(res, final)
		}
	}
	if len(res) == 0 {
		return [][]string{}
	}
	return res
}