package framework

import (
	"errors"
	"strings"
)

type Node struct {
	isLast   bool                // 本身是否是最后一个节点, 即, 是否可以成为一个独立的url
	segment  string              // url中的字符串
	handlers []ControllerHandler // 中间件+控制器
	children []*Node             // 子节点
}

type Tree struct {
	root *Node // 根节点
}

// NewNode 初始化一个节点
func NewNode() *Node {
	return &Node{
		isLast:   false,
		segment:  "",
		children: []*Node{},
	}
}

// NewTree 初始化一棵字典树
func NewTree() *Tree {
	root := NewNode()
	return &Tree{root: root}
}

// 判断一个segment是否以:开头
func isAppropriateSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *Node) filterChildNodes(segment string) []*Node {
	if len(n.children) == 0 {
		return nil
	}
	// 这就是一个通配符, 可以匹配任何类型, 既然能匹配任何类型, 那就有可能是我们需要的
	// 所以将子节点全部返回
	if isAppropriateSegment(segment) {
		return n.children
	}
	nodes := make([]*Node, 0, len(n.children))
	for _, cNode := range n.children {
		// 如果满足1.下一层子节点有通配符; 2.下一层子节点没有通配符, 但是文本完全匹配, 则均可加入到nodes中去
		condition := isAppropriateSegment(cNode.segment) || cNode.segment == segment
		if condition {
			nodes = append(nodes, cNode)
		}
	}
	return nodes
}

// 判断路由是否在节点的所有子节点中存在
func (n *Node) matchNode(url string) *Node {
	segments := strings.SplitN(url, "/", 2)
	// 取出第一个部分
	segment := segments[0]
	if !isAppropriateSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// 得到匹配符合的下一层子节点
	cNodes := n.filterChildNodes(segment)
	if cNodes == nil || len(cNodes) == 0 {
		return nil
	}
	// 若只有一个segment, 那就是最后一个标记
	if len(segments) == 1 {
		// 那就判断cNode有无isLast标记
		for _, tn := range cNodes {
			if tn.isLast {
				return tn
			}
		}
		// 都不是最后一个节点
		return nil
	}
	// 如果有两个segment, 那就递归遍历每个子节点, 继续寻找
	for _, tn := range cNodes {
		// 这种写法是合理的, 因为本来传入进来的就是一个url, 返回的却是一个node
		// segments[1]里面可能包含很多个路由中的字段, 我们要做的就是寻找有没有匹配的
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

// AddRouter 添加路由, 添加前先看看路由是否已经存在
func (tree *Tree) AddRouter(url string, handlers []ControllerHandler) error {
	node := tree.root
	if node.matchNode(url) != nil {
		return errors.New("router exists:" + url)
	}
	segments := strings.Split(url, "/")
	for index, segment := range segments {
		if !isAppropriateSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *Node // 标记是否有合适的子节点
		childNodes := node.filterChildNodes(segment)
		// 如果有匹配的子节点
		if len(childNodes) > 0 {
			// 如果有segment相同的子节点，则选择这个子节点
			for _, cNode := range childNodes {
				if cNode.segment == segment {
					objNode = cNode
					break
				}
			}
		}
		if objNode == nil {
			// 创建一个当前Node的节点
			cNode := NewNode()
			cNode.segment = segment
			if isLast {
				cNode.isLast = true
				cNode.handlers = handlers
			}
			node.children = append(node.children, cNode)
			objNode = cNode
		}
		node = objNode
	}
	return nil
}

// FindHandler 查找控制器
func (tree *Tree) FindHandler(url string) []ControllerHandler {
	matchedNode := tree.root.matchNode(url)
	if matchedNode == nil {
		return nil
	}
	return matchedNode.handlers
}
