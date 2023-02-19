package sol

import "fmt"

const (
	EndProcNormal      EndProcReason = 0  // end process normal, find all solutions.
	EndProcByMaxStack  EndProcReason = -1 // end process by max stack size limit.
	EndProcByMaxSearch EndProcReason = -2 // end process by max search size limit.
)

type EndProcReason int

func (e EndProcReason) String() string {
	switch e {
	case EndProcNormal:
		return "ok"
	case EndProcByMaxStack:
		return "stack"
	case EndProcByMaxSearch:
		return "search"
	}
	return fmt.Sprintf("unknown:%d", e)
}

// Record : analyze record
type Record struct {
	// Init Cards
	InitCards string `gorm:"column:init_cards"`
	// solution count
	SolutionCount int `gorm:"column:solution_count"`
	// Average step
	AverageStep int `gorm:"column:average_step"`
	// Step difference
	DiffStep int `gorm:"column:diff_step"`
	// Min step
	MinStep int `gorm:"column:min_step"`
	// Max step
	MaxStep int `gorm:"column:max_step"`
	// Max stack size
	MaxStackSize int `gorm:"column:max_stack_size"`
	// max search size
	MaxSearchSize int `gorm:"column:max_search_size"`
	// search count
	SearchCount int `gorm:"column:search_count"`
	// end proc reason
	EndProcReason int `gorm:"column:end_proc_reason"`
}

func (a *Record) TableName() string {
	return "record"
}

func (a *Record) String() string {
	return fmt.Sprintf("init:%s\nsolution:%d\naverage:%d\ndiff:%d\nmin:%d\nmax:%d\nstack:%d\nsearch:%d\ncount:%d\nreason:%d",
		a.InitCards, a.SolutionCount, a.AverageStep, a.DiffStep, a.MinStep, a.MaxStep, a.MaxStackSize, a.MaxSearchSize, a.SearchCount, a.EndProcReason)
}
