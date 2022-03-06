package basic4

type node struct {
	childen   []*node
	handler   handlerFunc
	matchFunc matchFunc
	patter    string
	nodeType  int
}
