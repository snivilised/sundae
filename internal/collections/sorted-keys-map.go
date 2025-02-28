package collections

import (
	"cmp"
	"slices"

	"github.com/snivilised/sundae/internal/third/lo"
)

// https://go.dev/blog/comparable

// OrderedKeysMap provides a deterministic way to iterate the contents
// of a map. The problem with a regular iteration is that the order
// the elements are presented is not guaranteed. This can be
// detrimental when consistency is paramount. Any map that needs to
// be iterated deterministically, should use the Keys() function for
// the iteration loop and obtain the value using this key to index
// into the map.
type OrderedKeysMap[K cmp.Ordered, V any] map[K]V

func (m *OrderedKeysMap[K, V]) Keys() []K {
	keys := lo.Keys(*m)

	slices.Sort(keys)

	return keys
}
