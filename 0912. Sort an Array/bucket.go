package main

// 将数组分到有限数量的桶子里。每个桶子再个别排序
func bucket(nums *[]int) {
	bucket := make([]int, 2 *5*10001)
	for _, v := range *nums {
		bucket[v+50000]++
	}
	i := 0
	for k, v := range bucket {
		for v > 0 {
			(*nums)[i] = k-50000
			v--
			i++
		}
	}
	return
}