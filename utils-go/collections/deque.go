package collections

type Deque[T any] struct {
	arr  []T
	size int
}

func (d Deque[T]) PeekLeft() T {
	if d.size == 0 {
		panic("Deque is empty, cannot peek left")
	}
	return d.arr[0]
}

func (d Deque[T]) PeekRight() T {
	if d.size == 0 {
		panic("Deque is empty, cannot peek right")
	}
	return d.arr[d.size-1]
}

func (d *Deque[T]) PopLeft() T {
	if d.size == 0 {
		panic("Deque is empty, cannot pop left")
	}
	val := d.arr[0]
	d.arr = d.arr[1:]
	d.size--
	return val
}

func (d *Deque[T]) PopRight() T {
	if d.size == 0 {
		panic("Deque is empty, cannot pop right")
	}
	val := d.arr[d.size-1]
	d.arr = d.arr[:d.size-1]
	d.size--
	return val
}

func (d *Deque[T]) AppendLeft(val T) {
	d.arr = append([]T{val}, d.arr...)
	d.size++
}

func (d *Deque[T]) AppendRight(val T) {
	d.arr = append(d.arr, val)
	d.size++
}
