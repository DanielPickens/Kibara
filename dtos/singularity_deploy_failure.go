package dtos

import (
	"fmt"
	"io"

	"github.com/kibara/swaggering"
)

type SingularityDeployFailureSingularityDeployFailureReason string

const (
	SingularityDeployFailureSingularityDeployFailureReasonTASK_FAILED_ON_STARTUP         SingularityDeployFailureSingularityDeployFailureReason = "TASK_FAILED_ON_STARTUP"
	SingularityDeployFailureSingularityDeployFailureReasonTASK_FAILED_HEALTH_CHECKS      SingularityDeployFailureSingularityDeployFailureReason = "TASK_FAILED_HEALTH_CHECKS"
	SingularityDeployFailureSingularityDeployFailureReasonTASK_COULD_NOT_BE_SCHEDULED    SingularityDeployFailureSingularityDeployFailureReason = "TASK_COULD_NOT_BE_SCHEDULED"
	SingularityDeployFailureSingularityDeployFailureReasonTASK_NEVER_ENTERED_RUNNING     SingularityDeployFailureSingularityDeployFailureReason = "TASK_NEVER_ENTERED_RUNNING"
	SingularityDeployFailureSingularityDeployFailureReasonTASK_EXPECTED_RUNNING_FINISHED SingularityDeployFailureSingularityDeployFailureReason = "TASK_EXPECTED_RUNNING_FINISHED"
	SingularityDeployFailureSingularityDeployFailureReasonDEPLOY_CANCELLED               SingularityDeployFailureSingularityDeployFailureReason = "DEPLOY_CANCELLED"
	SingularityDeployFailureSingularityDeployFailureReasonDEPLOY_OVERDUE                 SingularityDeployFailureSingularityDeployFailureReason = "DEPLOY_OVERDUE"
	SingularityDeployFailureSingularityDeployFailureReasonFAILED_TO_SAVE_DEPLOY_STATE    SingularityDeployFailureSingularityDeployFailureReason = "FAILED_TO_SAVE_DEPLOY_STATE"
	SingularityDeployFailureSingularityDeployFailureReasonLOAD_BALANCER_UPDATE_FAILED    SingularityDeployFailureSingularityDeployFailureReason = "LOAD_BALANCER_UPDATE_FAILED"
	SingularityDeployFailureSingularityDeployFailureReasonPENDING_DEPLOY_REMOVED         SingularityDeployFailureSingularityDeployFailureReason = "PENDING_DEPLOY_REMOVED"
)

type SingularityDeployFailure struct {
	Reason  *SingularityDeployFailureSingularityDeployFailureReason `json:"reason,omitempty"`
	TaskId  *SingularityTaskId                                      `json:"taskId,omitempty"`
	Message *string                                                 `json:"message,omitempty"`
}

func (self *SingularityDeployFailure) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployFailure) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployFailure); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployFailure cannot copy the values from %#v", other)
}

func (self *SingularityDeployFailure) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployFailure) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployFailure) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployFailure", name)

	case "reason", "Reason":
		v, ok := value.(SingularityDeployFailureSingularityDeployFailureReason)
		if ok {
			self.Reason = &v
			return nil
		}
		return fmt.Errorf("Field reason/Reason: value %v(%T) couldn't be cast to type SingularityDeployFailureSingularityDeployFailureReason", value, value)

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			return nil
		}
		return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityDeployFailure) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployFailure", name)

	case "reason", "Reason":
		return *self.Reason, nil
		return nil, fmt.Errorf("Field Reason no set on Reason %+v", self)

	case "taskId", "TaskId":
		return self.TaskId, nil
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityDeployFailure) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployFailure", name)

	case "reason", "Reason":
		self.Reason = nil

	case "taskId", "TaskId":
		self.TaskId = nil

	case "message", "Message":
		self.Message = nil

	}

	return nil
}

func (self *SingularityDeployFailure) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployFailureList []*SingularityDeployFailure

func (self *SingularityDeployFailureList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployFailureList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployFailureList cannot copy the values from %#v", other)
}

func (list *SingularityDeployFailureList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployFailureList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployFailureList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
