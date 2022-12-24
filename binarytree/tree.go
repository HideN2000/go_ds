package binarytree

// OrderableFunc は 要素の大小順序を決定する関数である.
type OrderableFunc[T comparable] func(left, right T) bool

// Tree は 指定された比較可能 (comparable) かつ順序付き型 T の要素を効率的に管理するための構造体である.
type Tree[T comparable] struct {
	root, sentinel *node[T]
	op             OrderableFunc[T]
	size           int
}

// NewTree は 指定された比較可能 (comparable) 型 T とその大小順序を定義した関数 OrderableFunc[T] を引数にとり，
// T を効率的に管理する Tree[T] の構造体を返り値として返す.
// OrderableFunc[T] は Well-Defined でなくてはならない.
// すなわち 任意のTの要素 a,b,c に対して
//
//	'a (op) b' && 'b (op) c' -> 'a (op) c'
//
// が成り立たなくてはならない．
// 上記の条件が満たされない場合 Treeの正しい挙動は保証されない.
//
// 以下に T の型と Well-Defined な OrderableFunc[T] の 例を挙げる.
// <ex>
// [T = int]
//
//	func op(a, b int) bool {
//		if a < b {
//			return true
//		}
//		return false
//	}
//
// Time: O(1)
func NewTree[T comparable](operator OrderableFunc[T]) *Tree[T] {
	t := &Tree[T]{
		op: func(left, right T) bool {
			if left == right {
				return true
			}
			return operator(left, right)
		},
		sentinel: new(node[T]),
	}
	t.root = t.sentinel
	t.root.par = t.sentinel
	t.root.left = t.sentinel
	t.root.right = t.sentinel
	return t
}

// Len　は　呼び出し時点での要素数を返す
// Time: O(1)
func (t *Tree[T]) Len() int {
	return t.size
}

// Clear は Tree を初期化し，全要素を削除する
// Time : O(N)
func (t *Tree[T]) Clear() {
	t.root = t.sentinel
	t.root.par = t.sentinel
	t.root.left = t.sentinel
	t.root.right = t.sentinel
}

// Contains は　渡された value 値が Tree に含まれるかを判定する
// Time : O(log N)
func (t *Tree[T]) Contains(value T) bool {
	ptr := t.root
	for ptr != t.sentinel && ptr.value != value {
		if t.op(value, ptr.value) {
			ptr = ptr.left
		} else {
			ptr = ptr.right
		}
	}
	return ptr != t.sentinel && ptr.value == value
}

// GetKthElem は　 Tree に含まれる要素のうち k(0-index) 番目に小さい値と error 値 nil を返す.
// 与インデックス値は [0, SizeOfTree) の範囲になくてはならない.
// 上記の条件が守られない場合は ErrInvalidIndex が error 値として返される．
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Tree[T]) GetKthElem(k int) (T, error) {
	if k < 0 {
		k += t.size
	}
	if k < 0 || k >= t.size {
		return t.sentinel.value, ErrInvalidIndex
	}
	return t.kthElement(k + 1)
}

// Min は　 Tree に含まれる要素のうち最も小さい値と error 値 nil を返す.
// Tree に要素がない場合は ErrTreeEmpty が error 値として返される.
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Tree[T]) Min() (T, error) {
	if t.size == 0 {
		return t.sentinel.value, ErrTreeEmpty
	}
	ptr := t.root
	for ptr.left != t.sentinel {
		ptr = ptr.left
	}
	return ptr.value, nil
}

// Max は 　Tree に含まれる要素のうち最も大きい値と error 値 nil を返す.
// Tree に要素がない場合は ErrTreeEmpty が error 値として返される.
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Tree[T]) Max() (T, error) {
	if t.size == 0 {
		return t.sentinel.value, ErrTreeEmpty
	}
	ptr := t.root
	for ptr.right != t.sentinel {
		ptr = ptr.right
	}
	return ptr.value, nil
}

func (t *Tree[T]) LessThan(value T) int {
	if t.size == 0 {
		return 0
	}
	count, ptr := 0, t.root
	for ptr != t.sentinel {
		if t.op(value, ptr.value) {
			ptr = ptr.left
		} else {
			count += 1 + ptr.left.subtreeSize
			ptr = ptr.right
		}
	}
	return count
}

func (t Tree[T]) Between(left, right T) int {
	if !t.op(left, right) {
		return 0
	}
	return t.LessThan(right) - t.LessThan(left)
}

// Prev は　 Tree に含まれ,なおかつ,渡された value 値より真に小さいものの中での最大値と error 値 nil を返す.
// Tree に要素がない場合は ErrTreeEmpty が error 値として返される.
// 該当する要素がない場合は ErrNotFound が error 値として返される.
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Tree[T]) Prev(value T) (T, error) {
	if t.size == 0 {
		return t.sentinel.value, ErrTreeEmpty
	}
	ptr, retval, updated := t.root, t.sentinel.value, false
	for ptr != t.sentinel {
		if t.op(value, ptr.value) {
			ptr = ptr.left
		} else {
			retval, ptr, updated = ptr.value, ptr.right, true
		}
	}
	if !updated {
		return retval, ErrNotFound
	}
	return retval, nil
}

// Next は　 Tree に含まれ,なおかつ,渡された value 値より真に大きいものの中での最小値と error 値 nil を返す.
// Tree に要素がない場合は ErrTreeEmpty が error 値として返される.
// 該当する要素がない場合は ErrNotFound が error 値として返される.
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Tree[T]) Next(value T) (T, error) {
	if t.size == 0 {
		return t.sentinel.value, ErrTreeEmpty
	}
	ptr, retval, updated := t.root, t.sentinel.value, false
	for ptr != t.sentinel {
		if t.op(ptr.value, value) {
			ptr = ptr.right
		} else {
			retval, ptr, updated = ptr.value, ptr.left, true
		}
	}
	if !updated {
		return retval, ErrNotFound
	}
	return retval, nil
}

// Push は Tree に渡された value 値を新たに加える.
// Time: O(log N)
func (t *Tree[T]) Push(value T) {
	t.size++
	z, y, v := t.root, t.sentinel, newNode(value)
	for z != t.sentinel {
		y = z
		z.subtreeSize++
		if !t.op(z.value, value) {
			z = z.left
		} else {
			z = z.right
		}
	}
	v.par = y
	if y == t.sentinel {
		t.root = v
	} else if !t.op(y.value, value) {
		y.left = v
	} else {
		y.right = v
	}
	v.color, v.left, v.right = true, t.sentinel, t.sentinel
	t.fixUpInsert(v)
}

// Pop は Tree に渡された value 値を Tree から<1つだけ>削除する.
// 該当する要素がない場合は ErrNotFound が error 値として返され,操作は棄却される.
// 該当する要素がある場合は nil が error 値として返され,要素の削除が行われる．
// Time: O(log N)
func (t *Tree[T]) Pop(value T) error {
	z, err := t.findAddress(value)
	if err != nil {
		return err
	}
	t.size--
	y, yOriginalColor := z, z.color
	var p, q *node[T]
	if z.left == t.sentinel {
		p, q = z, z.right
		for p != t.sentinel {
			p.subtreeSize--
			p = p.par
		}
		t.transplant(z, q)
	} else if z.right == t.sentinel {
		p, q = z, z.left
		for p != t.sentinel {
			p.subtreeSize--
			p = p.par
		}
		t.transplant(z, q)
	} else {
		y = z.right
		for y.left != t.sentinel {
			y = y.left
		}
		p = y
		for p != t.sentinel {
			p.subtreeSize--
			p = p.par
		}
		y.subtreeSize, yOriginalColor, q = z.subtreeSize, y.color, y.right
		if y.par == z {
			q.par = y
		} else {
			t.transplant(y, y.right)
			y.right, z.right.par = z.right, y
		}
		t.transplant(z, y)
		y.left, z.left.par = z.left, y
		y.color = z.color
	}
	if !yOriginalColor {
		t.fixUpDelete(q)
	}
	return nil
}
