package dtos

import (
	"fmt"
	"io"

	"github.com/kibara/swaggering"
)

type MesosFileChunkObject struct {
	Data       *string `json:"data,omitempty"`
	Offset     *int64  `json:"offset,omitempty"`
	NextOffset *int64  `json:"nextOffset,omitempty"`
}

func (self *MesosFileChunkObject) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *MesosFileChunkObject) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosFileChunkObject); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosFileChunkObject cannot copy the values from %#v", other)
}

func (self *MesosFileChunkObject) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *MesosFileChunkObject) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *MesosFileChunkObject) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosFileChunkObject", name)

	case "data", "Data":
		v, ok := value.(string)
		if ok {
			self.Data = &v
			return nil
		}
		return fmt.Errorf("Field data/Data: value %v(%T) couldn't be cast to type string", value, value)

	case "offset", "Offset":
		v, ok := value.(int64)
		if ok {
			self.Offset = &v
			return nil
		}
		return fmt.Errorf("Field offset/Offset: value %v(%T) couldn't be cast to type int64", value, value)

	case "nextOffset", "NextOffset":
		v, ok := value.(int64)
		if ok {
			self.NextOffset = &v
			return nil
		}
		return fmt.Errorf("Field nextOffset/NextOffset: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *MesosFileChunkObject) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MesosFileChunkObject", name)

	case "data", "Data":
		return *self.Data, nil
		return nil, fmt.Errorf("Field Data no set on Data %+v", self)

	case "offset", "Offset":
		return *self.Offset, nil
		return nil, fmt.Errorf("Field Offset no set on Offset %+v", self)

	case "nextOffset", "NextOffset":
		return *self.NextOffset, nil
		return nil, fmt.Errorf("Field NextOffset no set on NextOffset %+v", self)

	}
}

func (self *MesosFileChunkObject) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosFileChunkObject", name)

	case "data", "Data":
		self.Data = nil

	case "offset", "Offset":
		self.Offset = nil

	case "nextOffset", "NextOffset":
		self.NextOffset = nil

	}

	return nil
}

func (self *MesosFileChunkObject) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type MesosFileChunkObjectList []*MesosFileChunkObject

func (self *MesosFileChunkObjectList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosFileChunkObjectList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosFileChunkObjectList cannot copy the values from %#v", other)
}

func (list *MesosFileChunkObjectList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *MesosFileChunkObjectList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MesosFileChunkObjectList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
