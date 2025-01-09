package prototype

type Inode interface {
	print(string)
	clone() Inode
}
