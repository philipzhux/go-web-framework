package gweb

import "strings"

const (
	PATTERN_NON_EXIST string = ""
)

type node struct {
	name string;
	leaf_pattern string;
	children []*node;
	wildcard bool
}

func newNode(name string, wildcard bool) (*node){
	return &node{
		name: name,
		leaf_pattern: "",
		children: make([]*node,0),
		wildcard: wildcard,
	}
}

/* for insert, only single matched node is required */
/* insert must match name strictly */
func (n *node) singleMatch(name string) (*node){
	for _,child:=range n.children {
		if child.name==name {
			return child
		}
	}
	return nil
}

/* to match a pattern, an exhausted search at non-leaf level is needed */
func (n *node) exhaustiveMatch(name string) ([]*node){
	ret := make([]*node,0)
	for _,child:=range n.children {
		if child.wildcard || child.name==name{
			ret = append(ret, child)
		}
	}
	return ret
}

func (n *node) insertPattern(names []string,full_pattern string,depth int) {
	if depth==len(names) || strings.HasPrefix(n.name,"*"){
		n.leaf_pattern = full_pattern
		return
	}
	matched_child := n.singleMatch(names[depth])
	if matched_child==nil {
		matched_child = newNode(names[depth],strings.HasPrefix(names[depth],"*")||strings.HasPrefix(names[depth],":"))
		if strings.HasPrefix(names[depth],"*") {
			matched_child.name = strings.Join(names[depth:],"/")
		}
		n.children = append(n.children,matched_child)
	}
	matched_child.insertPattern(names, full_pattern, depth+1)
}

func (n *node) getPattern(names []string,depth int) string{
	if depth==len(names) || strings.HasPrefix(n.name,"*"){
		return n.leaf_pattern
	}
	matched_children := n.exhaustiveMatch(names[depth])
	for _,matched_child := range matched_children{
		ret := matched_child.getPattern(names, depth+1)
		if ret!= PATTERN_NON_EXIST {
			return ret
		}
	}
	return PATTERN_NON_EXIST
}

func parsePath(path string) []string {
	splits := strings.Split(path,"/")
	var i int
	var v string
	for i,v = range splits {
		if len(v)>0 {
			break
		}
	}
	return splits[i:]
}