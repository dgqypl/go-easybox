package _map

func PutAll(src map[any]any, dest map[any]any) {
	for k, v := range src {
		dest[k] = v
	}
}
