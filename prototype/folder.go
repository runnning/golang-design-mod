package prototype

import "fmt"

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indent string) {
	fmt.Println(indent + f.name)
	for _, i := range f.children {
		i.print(indent + indent)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		tempChildren = append(tempChildren, i.clone())
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
