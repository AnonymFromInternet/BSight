package sortedlist

type Node struct {
	Value int
	Next  *Node
}

type Sortedlist struct {
	Head *Node
}

func New() *Sortedlist {
	return &Sortedlist{Head: nil}
}

func (s *Sortedlist) Insert(value int) {
	newNode := &Node{Value: value}

	if s.Head == nil || s.Head.Value >= value {
		newNode.Next = s.Head
		s.Head = newNode
		return
	}

	current := s.Head
	for current.Next != nil && current.Next.Value < value {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
}

func (s *Sortedlist) Delete(value int) bool {
	if s.Head == nil {
		return false
	}

	if s.Head.Value == value {
		s.Head = s.Head.Next
		return true
	}

	current := s.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	if current.Next == nil {
		return false
	}
	current.Next = current.Next.Next
	return true
}

func (s *Sortedlist) IsExists(value int) bool {
	current := s.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (s *Sortedlist) Search(value int) *Node {
	current := s.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}
