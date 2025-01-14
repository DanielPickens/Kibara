package dtos

import (
	"fmt"
	"io"

	"github.com/kibara/swaggering"
)

type SingularityDeployStatisticsExtendedTaskState string

const (
	SingularityDeployStatisticsExtendedTaskStateTASK_LAUNCHED        SingularityDeployStatisticsExtendedTaskState = "TASK_LAUNCHED"
	SingularityDeployStatisticsExtendedTaskStateTASK_STAGING         SingularityDeployStatisticsExtendedTaskState = "TASK_STAGING"
	SingularityDeployStatisticsExtendedTaskStateTASK_STARTING        SingularityDeployStatisticsExtendedTaskState = "TASK_STARTING"
	SingularityDeployStatisticsExtendedTaskStateTASK_RUNNING         SingularityDeployStatisticsExtendedTaskState = "TASK_RUNNING"
	SingularityDeployStatisticsExtendedTaskStateTASK_CLEANING        SingularityDeployStatisticsExtendedTaskState = "TASK_CLEANING"
	SingularityDeployStatisticsExtendedTaskStateTASK_KILLING         SingularityDeployStatisticsExtendedTaskState = "TASK_KILLING"
	SingularityDeployStatisticsExtendedTaskStateTASK_FINISHED        SingularityDeployStatisticsExtendedTaskState = "TASK_FINISHED"
	SingularityDeployStatisticsExtendedTaskStateTASK_FAILED          SingularityDeployStatisticsExtendedTaskState = "TASK_FAILED"
	SingularityDeployStatisticsExtendedTaskStateTASK_KILLED          SingularityDeployStatisticsExtendedTaskState = "TASK_KILLED"
	SingularityDeployStatisticsExtendedTaskStateTASK_LOST            SingularityDeployStatisticsExtendedTaskState = "TASK_LOST"
	SingularityDeployStatisticsExtendedTaskStateTASK_LOST_WHILE_DOWN SingularityDeployStatisticsExtendedTaskState = "TASK_LOST_WHILE_DOWN"
	SingularityDeployStatisticsExtendedTaskStateTASK_ERROR           SingularityDeployStatisticsExtendedTaskState = "TASK_ERROR"
)

type SingularityDeployStatistics struct {
	DeployId                            *string                                       `json:"deployId,omitempty"`
	NumTasks                            *int32                                        `json:"numTasks,omitempty"`
	NumSuccess                          *int32                                        `json:"numSuccess,omitempty"`
	NumFailures                         *int32                                        `json:"numFailures,omitempty"`
	InstanceSequentialFailureTimestamps *map[int64][]int64                            `json:"instanceSequentialFailureTimestamps,omitempty"`
	LastTaskState                       *SingularityDeployStatisticsExtendedTaskState `json:"lastTaskState,omitempty"`
	RequestId                           *string                                       `json:"requestId,omitempty"`
	NumSequentialRetries                *int32                                        `json:"numSequentialRetries,omitempty"`
	LastFinishAt                        *int64                                        `json:"lastFinishAt,omitempty"`
	AverageRuntimeMillis                *int64                                        `json:"averageRuntimeMillis,omitempty"`
}

func (self *SingularityDeployStatistics) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployStatistics) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployStatistics); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployStatistics cannot copy the values from %#v", other)
}

func (self *SingularityDeployStatistics) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployStatistics) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployStatistics) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployStatistics", name)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = &v
			return nil
		}
		return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)

	case "numTasks", "NumTasks":
		v, ok := value.(int32)
		if ok {
			self.NumTasks = &v
			return nil
		}
		return fmt.Errorf("Field numTasks/NumTasks: value %v(%T) couldn't be cast to type int32", value, value)

	case "numSuccess", "NumSuccess":
		v, ok := value.(int32)
		if ok {
			self.NumSuccess = &v
			return nil
		}
		return fmt.Errorf("Field numSuccess/NumSuccess: value %v(%T) couldn't be cast to type int32", value, value)

	case "numFailures", "NumFailures":
		v, ok := value.(int32)
		if ok {
			self.NumFailures = &v
			return nil
		}
		return fmt.Errorf("Field numFailures/NumFailures: value %v(%T) couldn't be cast to type int32", value, value)

	case "instanceSequentialFailureTimestamps", "InstanceSequentialFailureTimestamps":
		v, ok := value.(map[int64][]int64)
		if ok {
			self.InstanceSequentialFailureTimestamps = &v
			return nil
		}
		return fmt.Errorf("Field instanceSequentialFailureTimestamps/InstanceSequentialFailureTimestamps: value %v(%T) couldn't be cast to type map[int64][]int64", value, value)

	case "lastTaskState", "LastTaskState":
		v, ok := value.(SingularityDeployStatisticsExtendedTaskState)
		if ok {
			self.LastTaskState = &v
			return nil
		}
		return fmt.Errorf("Field lastTaskState/LastTaskState: value %v(%T) couldn't be cast to type SingularityDeployStatisticsExtendedTaskState", value, value)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = &v
			return nil
		}
		return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)

	case "numSequentialRetries", "NumSequentialRetries":
		v, ok := value.(int32)
		if ok {
			self.NumSequentialRetries = &v
			return nil
		}
		return fmt.Errorf("Field numSequentialRetries/NumSequentialRetries: value %v(%T) couldn't be cast to type int32", value, value)

	case "lastFinishAt", "LastFinishAt":
		v, ok := value.(int64)
		if ok {
			self.LastFinishAt = &v
			return nil
		}
		return fmt.Errorf("Field lastFinishAt/LastFinishAt: value %v(%T) couldn't be cast to type int64", value, value)

	case "averageRuntimeMillis", "AverageRuntimeMillis":
		v, ok := value.(int64)
		if ok {
			self.AverageRuntimeMillis = &v
			return nil
		}
		return fmt.Errorf("Field averageRuntimeMillis/AverageRuntimeMillis: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularityDeployStatistics) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployStatistics", name)

	case "deployId", "DeployId":
		return *self.DeployId, nil
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "numTasks", "NumTasks":
		return *self.NumTasks, nil
		return nil, fmt.Errorf("Field NumTasks no set on NumTasks %+v", self)

	case "numSuccess", "NumSuccess":
		return *self.NumSuccess, nil
		return nil, fmt.Errorf("Field NumSuccess no set on NumSuccess %+v", self)

	case "numFailures", "NumFailures":
		return *self.NumFailures, nil
		return nil, fmt.Errorf("Field NumFailures no set on NumFailures %+v", self)

	case "instanceSequentialFailureTimestamps", "InstanceSequentialFailureTimestamps":
		return *self.InstanceSequentialFailureTimestamps, nil
		return nil, fmt.Errorf("Field InstanceSequentialFailureTimestamps no set on InstanceSequentialFailureTimestamps %+v", self)

	case "lastTaskState", "LastTaskState":
		return *self.LastTaskState, nil
		return nil, fmt.Errorf("Field LastTaskState no set on LastTaskState %+v", self)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "numSequentialRetries", "NumSequentialRetries":
		return *self.NumSequentialRetries, nil
		return nil, fmt.Errorf("Field NumSequentialRetries no set on NumSequentialRetries %+v", self)

	case "lastFinishAt", "LastFinishAt":
		return *self.LastFinishAt, nil
		return nil, fmt.Errorf("Field LastFinishAt no set on LastFinishAt %+v", self)

	case "averageRuntimeMillis", "AverageRuntimeMillis":
		return *self.AverageRuntimeMillis, nil
		return nil, fmt.Errorf("Field AverageRuntimeMillis no set on AverageRuntimeMillis %+v", self)

	}
}

func (self *SingularityDeployStatistics) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployStatistics", name)

	case "deployId", "DeployId":
		self.DeployId = nil

	case "numTasks", "NumTasks":
		self.NumTasks = nil

	case "numSuccess", "NumSuccess":
		self.NumSuccess = nil

	case "numFailures", "NumFailures":
		self.NumFailures = nil

	case "instanceSequentialFailureTimestamps", "InstanceSequentialFailureTimestamps":
		self.InstanceSequentialFailureTimestamps = nil

	case "lastTaskState", "LastTaskState":
		self.LastTaskState = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	case "numSequentialRetries", "NumSequentialRetries":
		self.NumSequentialRetries = nil

	case "lastFinishAt", "LastFinishAt":
		self.LastFinishAt = nil

	case "averageRuntimeMillis", "AverageRuntimeMillis":
		self.AverageRuntimeMillis = nil

	}

	return nil
}

func (self *SingularityDeployStatistics) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployStatisticsList []*SingularityDeployStatistics

func (self *SingularityDeployStatisticsList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployStatisticsList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployStatisticsList cannot copy the values from %#v", other)
}

func (list *SingularityDeployStatisticsList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployStatisticsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployStatisticsList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
