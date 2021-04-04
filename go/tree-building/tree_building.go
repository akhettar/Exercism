package tree

import (
	"errors"
	"sort"
)



type Record struct {
	ID int
	Parent int
}

type Node struct {
	ID int
	Children []*Node
}


func Build(recs []Record)(*Node, error) {

	// Sort the record
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].ID < recs[j].ID
	})

	if len(recs) == 0 {
		return nil, nil
	}
	
	// root has parent
	if recs[0].Parent > 0 {
		return nil, errors.New("Root has parent")
	}

	if len(recs) > 2 {
		// duplicate root
		if recs[0].ID == recs[1].ID {
			return nil, errors.New("duplicate root")
		}
	} 

	// has no root
	if recs[0].ID > 0 {
		return nil, errors.New("recods has no root")
	}

	root := Node{0, nil}
	references := make(map[int]*Node)

	for  i := 1 ; i < len(recs); i++ {
		child := Node{recs[i].ID, nil}

		// check for duplicate node
		if _, found := references[child.ID]; found {
			return nil, errors.New("duplicate node")
		}

		// check for duplicate root
		if child.ID == root.ID {
			return nil, errors.New("duplicate root")
		}

		// check for cycle directly
		if (child.ID == recs[i].Parent) {
			return nil, errors.New("cycle directly")
		}
		
		// higher id parent of lower id
		if child.ID < recs[i].Parent {
			return nil, errors.New("higher id parent of lower id")
		}
		// update the reference map
		references[child.ID] = &child

		// is the record a child of the root
		if recs[i].Parent == 0 {
			root.Children = append(root.Children, &child)
			continue
		}

		parent := references[recs[i].Parent]
		parent.Children = append(parent.Children, &child)
	}
	return &root, nil
}




