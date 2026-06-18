package main

import (
	"fmt"
	"slices"
)

func main() {
	// fmt.Println("=-=-=-=-=-=-=--=-=-=-=-=-=-=-=")
	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println("slice cha")
	// fmt.Println("Check the slice: ", reflect.TypeOf(slice).Kind() == reflect.Slice)
	// fmt.Println("slice len: ", len(slice))
	// fmt.Println("slice cap: ", cap(slice))
	// fmt.Println("")
	// fmt.Println("=-=-=-=-=-=-=--=-=-=-=-=-=-=-=")
	// fmt.Println("slice con")
	// subSlice := slice[3:6]
	// fmt.Println(subSlice)
	// fmt.Println("Check the subSlice: ", reflect.TypeOf(subSlice).Kind() == reflect.Slice)
	// fmt.Println("subSlice cap: ", cap(subSlice))
	// fmt.Println("subSlice len: ", len(subSlice))

	// tìm hiểu thêm về các hàm slice
	// ** clone: tạo bản sao cho slice **
	cloneSlice := slices.Clone([]int{1, 2, 3})
	fmt.Println(cloneSlice)

	// So sánh hai slide xem có giống nhau không
	//** equal: so sánh hai slice có giống nhau không **
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	compareSlice := slices.Equal(slice1, slice2)
	fmt.Println(compareSlice)

	//Tìm vị trí đầu tiên của phần tử
	//** Index: tìm vị trí đầu tiên của phần tử trong slice **
	index := slices.Index([]int{1, 2, 4, 5, 3, 3}, 3)
	fmt.Println(index)

	//Kiểm tra phần tử có tồn tại trong slice hay không
	//** Contains: kiểm tra phần tử có tồn tại trong slice hay không **
	contains := slices.Contains([]int{1, 2, 4, 5, 3}, 9)
	fmt.Println(contains)

	//Chèn phần tủ vào slice
	//** Insert: chèn phần tử vào slice **
	insert := slices.Insert([]int{1, 2, 4, 5, 3}, 2, 9)
	fmt.Println(insert)

	//Xóa phần tử trong slice
	//** Delete: xóa phần tử trong slice **
	delete := slices.Delete([]int{1, 2, 9, 4, 5, 3}, 2, 5)
	fmt.Println(delete)

	// Đảo ngược slice
	//** Reverse: đảo ngược slice **
	reverse := []int{1, 2, 3, 4, 5}
	slices.Reverse(reverse)
	fmt.Println(reverse)

	// Sắp xếp slice theo thứ tự tăng dần
	//** Sort: sắp xếp slice **
	sort := []int{5, 6, 3, 1, 2, 4}
	slices.Sort(sort)
	fmt.Println(sort)

	//sắp xếp slice theo điều kiện tùy chỉnh
	//** SortFunc: sắp xếp slice theo điều kiện tùy chỉnh **
	sortFunc := []int{5, 6, 3, 1, 2, 4}
	slices.SortFunc(sortFunc, func(i, j int) int {
		return j - i
	})
	fmt.Println(sortFunc)

	//lấy giá trị lớn nhất trong slice
	//** Max: lấy giá trị lớn nhất trong slice **
	max := slices.Max([]int{5, 6, 3, 1, 2, 4})
	fmt.Println(max)

	//lấy giá trị nhỏ nhất trong slice
	//** Min: lấy giá trị nhỏ nhất trong slice **
	min := slices.Min([]int{5, 6, 3, 1, 2, 4})
	fmt.Println(min)
}
