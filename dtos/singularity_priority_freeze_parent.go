package dtos

import (
	"fmt"
	"io"

	"github.com/kibara/swaggering"
)

type SingularityPriorityFreezeParent struct {
	PriorityFreeze *SingularityPriorityFreeze `json:"priorityFreeze,omitempty"`
	Timestamp      *int64                     `json:"timestamp,omitempty"`
	User           *string                    `json:"user,omitempty"`
}

func (self *SingularityPriorityFreezeParent) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPriorityFreezeParent) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPriorityFreezeParent); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPriorityFreezeParent cannot copy the values from %#v", other)
}

func (self *SingularityPriorityFreezeParent) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPriorityFreezeParent) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPriorityFreezeParent) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPriorityFreezeParent", name)

	case "priorityFreeze", "PriorityFreeze":
		v, ok := value.(*SingularityPriorityFreeze)
		if ok {
			self.PriorityFreeze = v
			return nil
		}
		return fmt.Errorf("Field priorityFreeze/PriorityFreeze: value %v(%T) couldn't be cast to type *SingularityPriorityFreeze", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityPriorityFreezeParent) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPriorityFreezeParent", name)

	case "priorityFreeze", "PriorityFreeze":
		return self.PriorityFreeze, nil
		return nil, fmt.Errorf("Field PriorityFreeze no set on PriorityFreeze %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityPriorityFreezeParent) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPriorityFreezeParent", name)

	case "priorityFreeze", "PriorityFreeze":
		self.PriorityFreeze = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "user", "User":
		self.User = nil

	}

	return nil
}

func (self *SingularityPriorityFreezeParent) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPriorityFreezeParentList []*SingularityPriorityFreezeParent

func (self *SingularityPriorityFreezeParentList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPriorityFreezeParentList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPriorityFreezeParentList cannot copy the values from %#v", other)
}

func (list *SingularityPriorityFreezeParentList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPriorityFreezeParentList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPriorityFreezeParentList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
