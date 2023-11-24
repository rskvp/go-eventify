package explorer

type ExplorerTreeNode struct {
	Key      string              `json:"key"`
	Title    string              `json:"title"`
	IsLeaf   bool                `json:"isLeaf"`
	Children []*ExplorerTreeNode `json:"children"`
}

type ExplorerTree []*ExplorerTreeNode
