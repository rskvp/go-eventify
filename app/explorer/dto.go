package explorer

type ExplorerTreeNode struct {
	Key      string              `json:"key"`
	Label    string              `json:"label"`
	Children []*ExplorerTreeNode `json:"children"`
	Type     string              `json:"type"`
}

type ExplorerTree []*ExplorerTreeNode

type AddFlowRequest struct {
	Name string `json:"name" binding:"required"`
}

type AddEventRequest struct {
	Name   string `json:"name" binding:"required"`
	FlowId string `json:"flowId" binding:"required"`
}
