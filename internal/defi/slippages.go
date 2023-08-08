package defi

import (
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

var SlippageMap = map[v1.TaskType]SlippagePercent{
	v1.TaskType_EzkaliburSwap:  SlippagePercent05,
	v1.TaskType_IzumiSwap:      SlippagePercentZero,
	v1.TaskType_MaverickSwap:   SlippagePercent2,
	v1.TaskType_MuteioSwap:     SlippagePercent01,
	v1.TaskType_SpaceFISwap:    SlippagePercent01,
	v1.TaskType_StargateBridge: SlippagePercent05,
	v1.TaskType_SyncSwap:       SlippagePercent01,
	v1.TaskType_VelocoreSwap:   SlippagePercent05,
	v1.TaskType_VeSyncSwap:     SlippagePercent01,
	v1.TaskType_ZkSwap:         SlippagePercent05,
}
