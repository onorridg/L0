package inMemory

import (
	"fmt"
	"time"
)

type Node struct {
	Id        uint64
	updatedAt time.Time
	prev      *Node
	next      *Node
}

func (n *Node) Append(id uint64) *Node {
	if n == nil {
		return initList(id)
	}
	n = n.FindLastNode()
	n.next = &Node{
		Id:        id,
		updatedAt: time.Now(),
		prev:      n,
		next:      nil,
	}
	return n.next
}

func (n *Node) FindLastNode() *Node {
	if n == nil {
		return nil
	}

	var lastNode *Node
	for n != nil {
		lastNode = n
		n = n.next
	}
	return lastNode
}

func (n *Node) FindFirstNode() *Node {
	if n == nil {
		return nil
	}

	var head *Node
	for n != nil {
		head = n
		n = n.prev
	}
	return head
}

func (n *Node) FindNode(id uint64) *Node {
	if n == nil {
		return nil
	}

	n = n.FindFirstNode()
	for n != nil {
		//fmt.Println("Now id:", n.Id, "| Find id:", id)
		if n.Id == id {
			return n
		}
		n = n.next
	}
	return nil
}

func (n *Node) Update(id uint64) *Node {
	//fmt.Println("Update func:", node)
	if n == nil {
		return &Node{Id: id, updatedAt: time.Now(), prev: nil, next: nil}
	}

	node := n.FindNode(id)
	if node.prev == nil {
		if node.next == nil {
			return node
		} else if node.next != nil {
			node.next.prev = nil
			n.Append(id)
			return node.next
		}
	} else if &node.prev != nil {
		if node.next == nil {
			return node
		} else if node.next != nil {
			node.prev.next = node.next

			n.Append(id)
			return node.prev
		}
	}
	return nil
}

func (n *Node) Delete() *Node {
	if n == nil {
		return nil
	}
	head := n.FindFirstNode()
	if head.next == nil {
		return nil
	}
	head.next.prev = nil
	return head.next
}

func (n *Node) Len() uint64 {
	if n == nil {
		return 0
	}

	var listLen uint64

	node := n.FindFirstNode()
	for node != nil {
		fmt.Println("Node Id:", node.Id, "| updatedAt:", node.updatedAt)
		listLen += 1
		node = node.next
	}
	return listLen
}

func initList(id uint64) *Node {
	head := &Node{
		Id:        id,
		updatedAt: time.Now(),
		prev:      nil,
		next:      nil,
	}
	return head
}
