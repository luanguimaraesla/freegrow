package node

type nodePhase string

type NodeStatus struct {
	Phase nodePhase
}

const (
	NodePhaseRunning nodePhase = "running"
	NodePhaseStopped nodePhase = "stopped"
)
