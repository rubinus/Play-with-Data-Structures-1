package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/05-Recursion/03-Recursion-Basics/listnode"
	"github.com/donng/Play-with-Data-Structures/05-Recursion/03-Recursion-Basics/solution3"
)

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := listnode.GetListNode(nums)
	fmt.Println(head)

	res := solution3.RemoveElements(head, 6)
	fmt.Println(res)
}
