package prototype

import "testing"

func TestPrototype(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	folder1 := &Folder{
		children: []Inode{file1, file2},
		name:     "Folder1",
	}
	cloneFolder := folder1.clone()
	folder1.print(" ")
	cloneFolder.print(" ")
}
