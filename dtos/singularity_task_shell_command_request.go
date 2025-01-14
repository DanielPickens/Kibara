package dtos

import (
	"fmt"
	"io"

	"github.com/kibara/swaggering"
)

type SingularityTaskShellCommandRequest struct {
	TaskId       *SingularityTaskId       `json:"taskId,omitempty"`
	User         *string                  `json:"user,omitempty"`
	ShellCommand *SingularityShellCommand `json:"shellCommand,omitempty"`
	Timestamp    *int64                   `json:"timestamp,omitempty"`
}

func (self *SingularityTaskShellCommandRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskShellCommandRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskShellCommandRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskShellCommandRequest cannot copy the values from %#v", other)
}

func (self *SingularityTaskShellCommandRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskShellCommandRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskShellCommandRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandRequest", name)

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			return nil
		}
		return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "shellCommand", "ShellCommand":
		v, ok := value.(*SingularityShellCommand)
		if ok {
			self.ShellCommand = v
			return nil
		}
		return fmt.Errorf("Field shellCommand/ShellCommand: value %v(%T) couldn't be cast to type *SingularityShellCommand", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularityTaskShellCommandRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskShellCommandRequest", name)

	case "taskId", "TaskId":
		return self.TaskId, nil
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "shellCommand", "ShellCommand":
		return self.ShellCommand, nil
		return nil, fmt.Errorf("Field ShellCommand no set on ShellCommand %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityTaskShellCommandRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandRequest", name)

	case "taskId", "TaskId":
		self.TaskId = nil

	case "user", "User":
		self.User = nil

	case "shellCommand", "ShellCommand":
		self.ShellCommand = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	}

	return nil
}

func (self *SingularityTaskShellCommandRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskShellCommandRequestList []*SingularityTaskShellCommandRequest

func (self *SingularityTaskShellCommandRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskShellCommandRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskShellCommandRequestList cannot copy the values from %#v", other)
}

func (list *SingularityTaskShellCommandRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskShellCommandRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskShellCommandRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
