package dtos

import (
	"fmt"
	"io"

	"github.com/kibara/swaggering"
)

type SingularityBounceRequest struct {
	Incremental               *bool                    `json:"incremental,omitempty"`
	SkipHealthchecks          *bool                    `json:"skipHealthchecks,omitempty"`
	RunShellCommandBeforeKill *SingularityShellCommand `json:"runShellCommandBeforeKill,omitempty"`
	DurationMillis            *int64                   `json:"durationMillis,omitempty"`
	ActionId                  *string                  `json:"actionId,omitempty"`
	Message                   *string                  `json:"message,omitempty"`
}

func (self *SingularityBounceRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityBounceRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityBounceRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityBounceRequest cannot copy the values from %#v", other)
}

func (self *SingularityBounceRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityBounceRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityBounceRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityBounceRequest", name)

	case "incremental", "Incremental":
		v, ok := value.(bool)
		if ok {
			self.Incremental = &v
			return nil
		}
		return fmt.Errorf("Field incremental/Incremental: value %v(%T) couldn't be cast to type bool", value, value)

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		v, ok := value.(*SingularityShellCommand)
		if ok {
			self.RunShellCommandBeforeKill = v
			return nil
		}
		return fmt.Errorf("Field runShellCommandBeforeKill/RunShellCommandBeforeKill: value %v(%T) couldn't be cast to type *SingularityShellCommand", value, value)

	case "durationMillis", "DurationMillis":
		v, ok := value.(int64)
		if ok {
			self.DurationMillis = &v
			return nil
		}
		return fmt.Errorf("Field durationMillis/DurationMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityBounceRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityBounceRequest", name)

	case "incremental", "Incremental":
		return *self.Incremental, nil
		return nil, fmt.Errorf("Field Incremental no set on Incremental %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		return self.RunShellCommandBeforeKill, nil
		return nil, fmt.Errorf("Field RunShellCommandBeforeKill no set on RunShellCommandBeforeKill %+v", self)

	case "durationMillis", "DurationMillis":
		return *self.DurationMillis, nil
		return nil, fmt.Errorf("Field DurationMillis no set on DurationMillis %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityBounceRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityBounceRequest", name)

	case "incremental", "Incremental":
		self.Incremental = nil

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		self.RunShellCommandBeforeKill = nil

	case "durationMillis", "DurationMillis":
		self.DurationMillis = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "message", "Message":
		self.Message = nil

	}

	return nil
}

func (self *SingularityBounceRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityBounceRequestList []*SingularityBounceRequest

func (self *SingularityBounceRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityBounceRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityBounceRequestList cannot copy the values from %#v", other)
}

func (list *SingularityBounceRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityBounceRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityBounceRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
