package model

type StatusEnum int

const (
	Started = iota
	Ongoing
	Completed
	Removed
)

var stateName = map[StatusEnum]string{
	Started:   "started",
	Ongoing:   "ongoing",
	Completed: "completed",
	Removed:   "removed",
}

func (ss StatusEnum) String() string {
	return stateName[ss]
}

type Todo struct {
	Name   string
	Status StatusEnum
}
