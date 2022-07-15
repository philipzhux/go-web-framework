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

/* to insert a pattern, only single matched node is required */
func (this *node) singleMatch(name string) (*node){
	for _,child:=range this.children {
		if child.name==name {
			return child
		}
	}
	return nil
}

/* to match a pattern, an exhausted search at non-leaf level is needed */
func (this *node) exhaustiveMatch(name string) ([]*node){
	ret := make([]*node,0)
	for _,child:=range this.children {
		if child.name==name {
			ret = append(ret, child)
		}
	}
	return ret
}

func (this *node) insertPattern(names []string,full_pattern string,depth int) {
	if depth==len(names){
		this.leaf_pattern = full_pattern
		return
	}
	matched_child := this.singleMatch(names[depth])
	if matched_child==nil {
		matched_child = newNode(names[depth],names[depth][0]=='*'||names[depth][0]==':')
		this.children = append(this.children,matched_child)
	}
	matched_child.insertPattern(names, full_pattern, depth+1)
}

func (this *node) getPattern(names []string,depth int) string{
	if depth==len(names){
		return this.leaf_pattern
	}
	matched_children := this.exhaustiveMatch(names[depth])
	for _,matched_child := range matched_children{
		ret := matched_child.getPattern(names, depth+1)
		if ret!= PATTERN_NON_EXIST {
			return ret
		}
	}
	return PATTERN_NON_EXIST
}

func parsePath(path string) []string {
	
}