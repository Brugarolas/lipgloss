package tree

// Enumerator returns the branch and tree prefixes of a given item.
type Enumerator func(i int, last bool) (indent string, prefix string)

// DefaultEnumerator enumerates items.
func DefaultEnumerator(_ int, last bool) (indent, prefix string) {
	if last {
		return "   ", "└──"
	}
	return "│  ", "├──"
}