package binarytree

type node[T comparable] struct {
	value            T
	par, left, right *node[T]
	subtreeSize      int
	color            bool
}

func newNode[T comparable](value T) *node[T] {
	p := &node[T]{
		value:       value,
		subtreeSize: 1,
	}
	return p
}

func (t *Tree[T]) kthElement(k int) (T, error) {
	ptr := t.root
	for ptr != t.sentinel {
		lsize := ptr.left.subtreeSize + 1
		if k == lsize {
			return ptr.value, nil
		} else if k < lsize {
			ptr = ptr.left
		} else {
			k -= lsize
			ptr = ptr.right
		}
	}
	return t.sentinel.value, ErrUnexpected //Unexpected Error
}

func (t *Tree[T]) findAddress(value T) (*node[T], error) {
	ptr := t.root
	for ptr != t.sentinel && value != ptr.value {
		if t.op(value, ptr.value) {
			ptr = ptr.left
		} else {
			ptr = ptr.right
		}
	}
	if ptr == t.sentinel {
		return ptr, ErrNotFound
	} else {
		return ptr, nil
	}
}

func (t *Tree[T]) fixUpInsert(z *node[T]) {
	for zp := z.par; zp.color; zp = z.par {
		if zpp := zp.par; zp == zpp.left {
			y := zpp.right
			if y.color {
				zp.color, y.color, zpp.color = false, false, true
				z = zpp
			} else {
				if z == zp.right {
					z = zp
					t.rotateLeft(z)
				}
				zp = z.par
				zpp = zp.par
				zp.color, zpp.color = false, true
				t.rotateRight(zpp)
			}
		} else {
			y := zpp.left
			if y.color {
				zp.color, y.color, zpp.color = false, false, true
				z = zpp
			} else {
				if z == zp.left {
					z = zp
					t.rotateRight(z)
				}
				zp = z.par
				zpp = zp.par
				zp.color, zpp.color = false, true
				t.rotateLeft(zpp)
			}
		}
	}
	t.root.color = false
}

func (t *Tree[T]) fixUpDelete(v *node[T]) {
	for v != t.root && !v.color {
		if vp := v.par; v == vp.left {
			w := vp.right
			if w.color {
				w.color, vp.color = false, true
				t.rotateLeft(vp)
				w = vp.right
			}
			if wl, wr := w.left, w.right; !wl.color && !wr.color {
				w.color, v = true, vp
			} else {
				if !wr.color {
					wl.color, w.color = false, true
					t.rotateRight(w)
					w = vp.right
				}
				w.color, vp.color, wr.color = vp.color, false, false
				t.rotateLeft(vp)
				v = t.root
			}
		} else {
			w := vp.left
			if w.color {
				w.color, vp.color = false, true
				t.rotateRight(vp)
				w = vp.left
			}
			if wl, wr := w.left, w.right; !wl.color && !wr.color {
				w.color, v = true, vp
			} else {
				if !wl.color {
					wr.color, w.color = false, true
					t.rotateLeft(w)
					w = vp.left
				}
				w.color, vp.color, wl.color = vp.color, false, false
				t.rotateRight(vp)
				v = t.root
			}
		}
	}
	v.color = false
}

func (t *Tree[T]) transplant(u, v *node[T]) {
	if up := u.par; up == t.sentinel {
		t.root = v
	} else if u == up.left {
		up.left = v
	} else {
		up.right = v
	}
	v.par = u.par
}

func (t *Tree[T]) rotateLeft(x *node[T]) {
	y := x.right
	x.right = y.left
	if yl := y.left; yl != t.sentinel {
		yl.par = x
	}
	y.par = x.par
	if x.par == t.sentinel {
		t.root = y
	} else if xp := x.par; x == xp.left {
		xp.left = y
	} else {
		xp.right = y
	}
	y.left, x.par = x, y
	y.subtreeSize = x.subtreeSize
	x.subtreeSize = x.left.subtreeSize + x.right.subtreeSize + 1
}

func (t *Tree[T]) rotateRight(x *node[T]) {
	y := x.left
	x.left = y.right
	if yr := y.right; yr != t.sentinel {
		yr.par = x
	}
	y.par = x.par
	if x.par == t.sentinel {
		t.root = y
	} else if xp := x.par; x == xp.right {
		xp.right = y
	} else {
		xp.left = y
	}
	y.right, x.par = x, y
	y.subtreeSize = x.subtreeSize
	x.subtreeSize = x.left.subtreeSize + x.right.subtreeSize + 1
}
