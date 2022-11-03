package main

import "fmt"

// 从1.18版本开始，Go添加了对泛型的支持，也即类型参数

// MapKeys 作为泛型函数的示例，MapKeys 接受任意类型的Map并返回其Key的切片
// 这个函数有2个类型参数 - K 和 V； K 是 comparable 类型，也就是说我们可以通过 == 和 != 操作符对这个类型的值进行比较
// 这是Go中Map的Key所必须的
// V 是 any 类型，意味着它不受任何限制 (any 是 interface{} 的别名类型)
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// List 作为泛型类型的示例， List 是一个 具有任意类型值的单链表
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push 可以像在常规类型上一样定义泛型类型的方法 但必须保留类型参数
// 这个类型是 List[T]，而不是 List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// 当调用泛型函数的时候, 经常可以使用类型推断
	// 注意，当调用 MapKeys 的时候， 不需要为 K 和 V 指定类型 - 编译器会进行自动推断
	fmt.Println("keys m:", MapKeys(m))

	// 虽然也可以明确指定这些类型
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(1114)
	lst.Push(1118)
	lst.Push(1225)
	fmt.Println("list:", lst.GetAll())
}
