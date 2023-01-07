package set

import (
	errors "github.com/hiden2000/go_ds/errors"
	internal "github.com/hiden2000/go_ds/internal/set"
)

// OrderableFunc は 要素の大小順序を決定する関数である.
type OrderableFunc[T comparable] func(left, right T) bool

// Set は 指定された比較可能 (comparable) かつ順序付き型 T の要素を効率的に管理するための構造体である.
type Set[T comparable] struct {
	root, sentinel *internal.Node[T]
	op             OrderableFunc[T]
	size           int
}

// New は 指定された比較可能 (comparable) 型 T とその大小順序を定義した関数 OrderableFunc[T] を引数にとり，
// T を効率的に管理する Set[T] の構造体を返り値として返す.
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
func New[T comparable](operator OrderableFunc[T]) *Set[T] {
	t := &Set[T]{
		op: func(left, right T) bool {
			if left == right {
				return true
			}
			return operator(left, right)
		},
		sentinel: new(internal.Node[T]),
	}
	t.root = t.sentinel
	t.root.Par = t.sentinel
	t.root.Left = t.sentinel
	t.root.Right = t.sentinel
	return t
}

// Len　は　呼び出し時点での要素数を返す
// Time: O(1)
func (t *Set[T]) Len() int {
	return t.size
}

// Clear は Set を初期化し，全要素を削除する
// Time : O(N)
func (t *Set[T]) Clear() {
	t.root = t.sentinel
	t.root.Par = t.sentinel
	t.root.Left = t.sentinel
	t.root.Right = t.sentinel
}

// Contains は　渡された value 値が Set に含まれるかを判定する
// Time : O(log N)
func (t *Set[T]) Contains(value T) bool {
	ptr := t.root
	for ptr != t.sentinel && ptr.Value != value {
		if t.op(value, ptr.Value) {
			ptr = ptr.Left
		} else {
			ptr = ptr.Right
		}
	}
	return ptr != t.sentinel && ptr.Value == value
}

// GetKthElem は　 Set に含まれる要素のうち k(0-index) 番目に小さい値と error 値 nil を返す.
// 与インデックス値は [0, SizeOfSet) の範囲になくてはならない.
// 上記の条件が守られない場合は ErrInvalidIndex が error 値として返される．
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Set[T]) GetKthElem(k int) (T, error) {
	if k < 0 {
		k += t.size
	}
	if k < 0 || k >= t.size {
		return t.sentinel.Value, errors.ErrInvalidIndex
	}
	return t.kthElement(k + 1)
}

// Min は　 Set に含まれる要素のうち最も小さい値と error 値 nil を返す.
// Set に要素がない場合は ErrNotFound が error 値として返される.
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Set[T]) Min() (T, error) {
	if t.size == 0 {
		return t.sentinel.Value, errors.ErrNotFound
	}
	ptr := t.root
	for ptr.Left != t.sentinel {
		ptr = ptr.Left
	}
	return ptr.Value, nil
}

// Max は 　Set に含まれる要素のうち最も大きい値と error 値 nil を返す.
// Set に要素がない場合は ErrNotFound が error 値として返される.
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Set[T]) Max() (T, error) {
	if t.size == 0 {
		return t.sentinel.Value, errors.ErrNotFound
	}
	ptr := t.root
	for ptr.Right != t.sentinel {
		ptr = ptr.Right
	}
	return ptr.Value, nil
}

func (t *Set[T]) LessThan(value T) int {
	if t.size == 0 {
		return 0
	}
	count, ptr := 0, t.root
	for ptr != t.sentinel {
		if t.op(value, ptr.Value) {
			ptr = ptr.Left
		} else {
			count += 1 + ptr.Left.SubtreeSize
			ptr = ptr.Right
		}
	}
	return count
}

func (t Set[T]) Between(left, right T) int {
	if !t.op(left, right) {
		return 0
	}
	return t.LessThan(right) - t.LessThan(left)
}

// Prev は　 Set に含まれ,なおかつ,渡された value 値より真に小さいものの中での最大値と error 値 nil を返す.
// Set に要素がない場合は ErrNotFound が error 値として返される.
// 該当する要素がない場合は ErrNotFound が error 値として返される.
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Set[T]) Prev(value T) (T, error) {
	if t.size == 0 {
		return t.sentinel.Value, errors.ErrNotFound
	}
	ptr, retval, updated := t.root, t.sentinel.Value, false
	for ptr != t.sentinel {
		if t.op(value, ptr.Value) {
			ptr = ptr.Left
		} else {
			retval, ptr, updated = ptr.Value, ptr.Right, true
		}
	}
	if !updated {
		return retval, errors.ErrNotFound
	}
	return retval, nil
}

// Next は　 Set に含まれ,なおかつ,渡された value 値より真に大きいものの中での最小値と error 値 nil を返す.
// Set に要素がない場合は ErrNotFound が error 値として返される.
// 該当する要素がない場合は ErrNotFound が error 値として返される.
// error 値が nil でない場合の返り値は不定である.
// Time: O(log N)
func (t *Set[T]) Next(value T) (T, error) {
	if t.size == 0 {
		return t.sentinel.Value, errors.ErrNotFound
	}
	ptr, retval, updated := t.root, t.sentinel.Value, false
	for ptr != t.sentinel {
		if t.op(ptr.Value, value) {
			ptr = ptr.Right
		} else {
			retval, ptr, updated = ptr.Value, ptr.Left, true
		}
	}
	if !updated {
		return retval, errors.ErrNotFound
	}
	return retval, nil
}

// Push は Set に渡された value 値を新たに加える.
// Time: O(log N)
func (t *Set[T]) Push(value T) {
	t.size++
	z, y, v := t.root, t.sentinel, internal.NewNode(value)
	for z != t.sentinel {
		y = z
		z.SubtreeSize++
		if !t.op(z.Value, value) {
			z = z.Left
		} else {
			z = z.Right
		}
	}
	v.Par = y
	if y == t.sentinel {
		t.root = v
	} else if !t.op(y.Value, value) {
		y.Left = v
	} else {
		y.Right = v
	}
	v.Color, v.Left, v.Right = true, t.sentinel, t.sentinel
	t.fixUpInsert(v)
}

// Pop は Set に渡された value 値を Set から<1つだけ>削除する.
// 該当する要素がない場合は ErrNotFound が error 値として返され,操作は棄却される.
// 該当する要素がある場合は nil が error 値として返され,要素の削除が行われる．
// Time: O(log N)
func (t *Set[T]) Pop(value T) error {
	z, err := t.findAddress(value)
	if err != nil {
		return err
	}
	t.size--
	y, yOriginalColor := z, z.Color
	var p, q *internal.Node[T]
	if z.Left == t.sentinel {
		p, q = z, z.Right
		for p != t.sentinel {
			p.SubtreeSize--
			p = p.Par
		}
		t.transplant(z, q)
	} else if z.Right == t.sentinel {
		p, q = z, z.Left
		for p != t.sentinel {
			p.SubtreeSize--
			p = p.Par
		}
		t.transplant(z, q)
	} else {
		y = z.Right
		for y.Left != t.sentinel {
			y = y.Left
		}
		p = y
		for p != t.sentinel {
			p.SubtreeSize--
			p = p.Par
		}
		y.SubtreeSize, yOriginalColor, q = z.SubtreeSize, y.Color, y.Right
		if y.Par == z {
			q.Par = y
		} else {
			t.transplant(y, y.Right)
			y.Right, z.Right.Par = z.Right, y
		}
		t.transplant(z, y)
		y.Left, z.Left.Par = z.Left, y
		y.Color = z.Color
	}
	if !yOriginalColor {
		t.fixUpDelete(q)
	}
	return nil
}

func (t *Set[T]) kthElement(k int) (T, error) {
	ptr := t.root
	for ptr != t.sentinel {
		lsize := ptr.Left.SubtreeSize + 1
		if k == lsize {
			return ptr.Value, nil
		} else if k < lsize {
			ptr = ptr.Left
		} else {
			k -= lsize
			ptr = ptr.Right
		}
	}
	return t.sentinel.Value, errors.ErrUnexpected //Unexpected Error
}

func (t *Set[T]) findAddress(value T) (*internal.Node[T], error) {
	ptr := t.root
	for ptr != t.sentinel && value != ptr.Value {
		if t.op(value, ptr.Value) {
			ptr = ptr.Left
		} else {
			ptr = ptr.Right
		}
	}
	if ptr == t.sentinel {
		return ptr, errors.ErrNotFound
	} else {
		return ptr, nil
	}
}

func (t *Set[T]) fixUpInsert(z *internal.Node[T]) {
	for zp := z.Par; zp.Color; zp = z.Par {
		if zpp := zp.Par; zp == zpp.Left {
			y := zpp.Right
			if y.Color {
				zp.Color, y.Color, zpp.Color = false, false, true
				z = zpp
			} else {
				if z == zp.Right {
					z = zp
					t.rotateLeft(z)
				}
				zp = z.Par
				zpp = zp.Par
				zp.Color, zpp.Color = false, true
				t.rotateRight(zpp)
			}
		} else {
			y := zpp.Left
			if y.Color {
				zp.Color, y.Color, zpp.Color = false, false, true
				z = zpp
			} else {
				if z == zp.Left {
					z = zp
					t.rotateRight(z)
				}
				zp = z.Par
				zpp = zp.Par
				zp.Color, zpp.Color = false, true
				t.rotateLeft(zpp)
			}
		}
	}
	t.root.Color = false
}

func (t *Set[T]) fixUpDelete(v *internal.Node[T]) {
	for v != t.root && !v.Color {
		if vp := v.Par; v == vp.Left {
			w := vp.Right
			if w.Color {
				w.Color, vp.Color = false, true
				t.rotateLeft(vp)
				w = vp.Right
			}
			if wl, wr := w.Left, w.Right; !wl.Color && !wr.Color {
				w.Color, v = true, vp
			} else {
				if !wr.Color {
					wl.Color, w.Color = false, true
					t.rotateRight(w)
					w = vp.Right
				}
				w.Color, vp.Color, wr.Color = vp.Color, false, false
				t.rotateLeft(vp)
				v = t.root
			}
		} else {
			w := vp.Left
			if w.Color {
				w.Color, vp.Color = false, true
				t.rotateRight(vp)
				w = vp.Left
			}
			if wl, wr := w.Left, w.Right; !wl.Color && !wr.Color {
				w.Color, v = true, vp
			} else {
				if !wl.Color {
					wr.Color, w.Color = false, true
					t.rotateLeft(w)
					w = vp.Left
				}
				w.Color, vp.Color, wl.Color = vp.Color, false, false
				t.rotateRight(vp)
				v = t.root
			}
		}
	}
	v.Color = false
}

func (t *Set[T]) transplant(u, v *internal.Node[T]) {
	if up := u.Par; up == t.sentinel {
		t.root = v
	} else if u == up.Left {
		up.Left = v
	} else {
		up.Right = v
	}
	v.Par = u.Par
}

func (t *Set[T]) rotateLeft(x *internal.Node[T]) {
	y := x.Right
	x.Right = y.Left
	if yl := y.Left; yl != t.sentinel {
		yl.Par = x
	}
	y.Par = x.Par
	if x.Par == t.sentinel {
		t.root = y
	} else if xp := x.Par; x == xp.Left {
		xp.Left = y
	} else {
		xp.Right = y
	}
	y.Left, x.Par = x, y
	y.SubtreeSize = x.SubtreeSize
	x.SubtreeSize = x.Left.SubtreeSize + x.Right.SubtreeSize + 1
}

func (t *Set[T]) rotateRight(x *internal.Node[T]) {
	y := x.Left
	x.Left = y.Right
	if yr := y.Right; yr != t.sentinel {
		yr.Par = x
	}
	y.Par = x.Par
	if x.Par == t.sentinel {
		t.root = y
	} else if xp := x.Par; x == xp.Right {
		xp.Right = y
	} else {
		xp.Left = y
	}
	y.Right, x.Par = x, y
	y.SubtreeSize = x.SubtreeSize
	x.SubtreeSize = x.Left.SubtreeSize + x.Right.SubtreeSize + 1
}
