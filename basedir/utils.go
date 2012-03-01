package basedir

type environ map[string]string

func (map1 environ) Eq(map2 environ) bool {
	if !(len(map1) == len(map2)) {
		return false
	}
	for key,value := range map1 {
		v, ok := map2[key]
		if !ok || v != value {
			return false
		}
	}
	return true
}
