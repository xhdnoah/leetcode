package main

func replaceSpace(s string) string {
	b := []byte(s)
	res := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, b[i])
		}
	}
	return string(res)
}