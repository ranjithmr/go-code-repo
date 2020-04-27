package reverse 

func Reverse(arr []int) []int{
	var res []int
	for i := len(arr)-1; i>=0; i-- {
		res = append(res, arr[i])
	}
	return res
}
