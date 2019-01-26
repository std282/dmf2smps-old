package smpsbuild

/*
	This file contains address/pointer/reference helper functions for patterns.
*/

import (
	"container/list"
)

// addRef adds address to reference list.
//
// When address is in the list, it will be later notified by updateRef about
// pattern actual position, so address can point to this pattern.
func (pat *Pattern) addRef(addr address) {
	pat.references.PushBack(addr)
}

func (pat *Pattern) removeRef(addr address) {
	var nodeRemove *list.Element
	for node := pat.references.Front(); node != nil; node = node.Next() {
		if a := node.Value.(address); a == addr {
			nodeRemove = node
			break
		}
	}

	if nodeRemove != nil {
		pat.references.Remove(nodeRemove)
	}
}

// updateRef notifies all pattern references about figuring out pattern position.
func (pat *Pattern) updateRef(pos uint) {
	pat.foreachRef(func(addr address) {
		addr.evaluate(pos)
	})

	pat.setInnerPointers(pos)
}

// isReferenced returns true if pattern was referenced by some address.
func (pat *Pattern) isReferenced() bool {
	return pat.references.Len() != 0
}
