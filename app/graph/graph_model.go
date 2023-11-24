package graph

const (
	NodeTypeInput      = "input"
	NodeTypeOutput     = "output"
	NodeTypeDefault    = "default"
	NodeTargetPosition = "left"
	NodeSourcePosition = "right"
	EdgeMarkerEndType  = "arrowclosed"
)

type FlowGraph struct {
	Nodes []*FlowGraphNode  `json:"nodes"`
	Edges []*FlowGraphEdges `json:"edges"`
}

type FlowGraphNode struct {
	Id             string                `json:"id"`
	Position       FlowGraphNodePosition `json:"position"`
	Data           FlowGraphNodeData     `json:"data"`
	TargetPosition string                `json:"targetPosition"`
	SourcePosition string                `json:"sourcePosition"`
	Type           string                `json:"type"`
}

type FlowGraphNodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type FlowGraphNodeData struct {
	Label string `json:"label"`
}

type FlowGraphEdges struct {
	Id        string              `json:"id"`
	Source    string              `json:"source"`
	Target    string              `json:"target"`
	MarkerEnd FlowGraphEdgeMarker `json:"markerEnd"`
}

type FlowGraphEdgeMarker struct {
	Type string `json:"type"`
}
