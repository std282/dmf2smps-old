package smpsbuild

import "container/list"

// addRef adds address to references list, so it will be set when the position
// of pattern will become clear
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

// updateRef notifies all addresses about figuring out pattern position
func (pat *Pattern) updateRef(pos uint) {
	pat.foreachRef(func(addr address) {
		addr.evaluate(pos)
	})
}

func (pat *Pattern) isReferenced() bool {
	return pat.references.Len() != 0
}
