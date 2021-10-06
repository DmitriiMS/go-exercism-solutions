// Package tree provides structures and function to build a tree of online
// forum threads out of a list of records.
package tree

import (
	"errors"
	"sort"
)

// Structure that represents a node in a tree
type Node struct {
	ID       int
	Children []*Node
}

// Structure of an record in an input
type Record struct {
	ID     int
	Parent int
}

// Build function takes a slice of records and transforms it into a tree of Nodes
func Build(records []Record) (*Node, error) {
	// check if input is empty
	if len(records) == 0 {
		return nil, nil
	}
	// sort slice by record IDs
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	// check for root node
	if records[0].ID != 0 {
		return nil, errors.New("no root node")
	}
	// check for continuity errors
	var prev int
	for i, r := range records {
		if i != 0 && r.ID-1 != prev {
			return nil, errors.New("records have continuity issues")
		}
		prev = r.ID
	}

	// MapOfNodes is a map of nodes organized by parents of records.
	// While building a map, this loop also checks for various errors.
	mapOfNodes := make(map[int][]*Node)
	for _, record := range records {
		if record.ID == 0 { // some root node checks
			if record.Parent != 0 { // root node parent shuould be 0 -- empty value for int
				return nil, errors.New("root node has parent")
			}
			if _, ok := mapOfNodes[-1]; ok { // if a root record already exists, this means we have two roots
				return nil, errors.New("duplicate root node")

			}
			// If everything is ok, make a special root entry that should have
			// root record ID and empty node pointer for children.
			// -1 is imaginary root's parent.
			mapOfNodes[-1] = append(mapOfNodes[-1], &Node{record.ID, []*Node{}})
			continue
		}
		// for non-root nodes
		// First check if parent ID of a node is smaller than the record ID.
		// If not, earlier record has newer record as a parent, or it's its on parent.
		// This is wrong and will lead to cycles.
		if record.Parent >= record.ID {
			return nil, errors.New("parent ID is higher than or equal to record ID")
		}
		// If an entry in a map already exists, check that the list of nodes for such parents
		// doesn't contain duplicates.
		if _, ok := mapOfNodes[record.Parent]; ok {
			for _, v := range mapOfNodes[record.Parent] {
				if record.ID == v.ID {
					return nil, errors.New("duplicate node")
				}
			}
		}
		// If all checks are cleared, append a record from input slice to correct
		// parent's entry in a map.
		mapOfNodes[record.Parent] = append(mapOfNodes[record.Parent], &Node{record.ID, []*Node{}})
	}

	// Walk each slice of the map and link nodes to make a tree
	for _, v := range mapOfNodes {
		for _, node := range v {
			node.Children = mapOfNodes[node.ID]
		}
	}
	return mapOfNodes[-1][0], nil
}
