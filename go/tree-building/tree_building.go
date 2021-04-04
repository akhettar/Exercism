package tree

// Node instance type
type Node struct {
	ID       int
	Children []*Node
}

// Record instance type
type Record struct {
	ID     int
	Parent int
}

func Build(records []Record) (*Node, error) {
	var root *Node
	for _, r := range records {
		if r.ID == 0 {

		}
	}

	return nil, nil
}
