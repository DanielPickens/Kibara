package dtos

import (
	"fmt"
	"io"

	"github.com/kibara/swaggering"
)

type SingularityDeployRequest struct {
	Deploy                    *SingularityDeploy  `json:"deploy,omitempty"`
	Message                   *string             `json:"message,omitempty"`
	UpdatedRequest            *SingularityRequest `json:"updatedRequest,omitempty"`
	UnpauseOnSuccessfulDeploy *bool               `json:"unpauseOnSuccessfulDeploy,omitempty"`
}

func (self *SingularityDeployRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployRequest cannot copy the values from %#v", other)
}

func (self *SingularityDeployRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployRequest", name)

	case "deploy", "Deploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.Deploy = v
			return nil
		}
		return fmt.Errorf("Field deploy/Deploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "updatedRequest", "UpdatedRequest":
		v, ok := value.(*SingularityRequest)
		if ok {
			self.UpdatedRequest = v
			return nil
		}
		return fmt.Errorf("Field updatedRequest/UpdatedRequest: value %v(%T) couldn't be cast to type *SingularityRequest", value, value)

	case "unpauseOnSuccessfulDeploy", "UnpauseOnSuccessfulDeploy":
		v, ok := value.(bool)
		if ok {
			self.UnpauseOnSuccessfulDeploy = &v
			return nil
		}
		return fmt.Errorf("Field unpauseOnSuccessfulDeploy/UnpauseOnSuccessfulDeploy: value %v(%T) couldn't be cast to type bool", value, value)

	}
}

func (self *SingularityDeployRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployRequest", name)

	case "deploy", "Deploy":
		return self.Deploy, nil
		return nil, fmt.Errorf("Field Deploy no set on Deploy %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "updatedRequest", "UpdatedRequest":
		return self.UpdatedRequest, nil
		return nil, fmt.Errorf("Field UpdatedRequest no set on UpdatedRequest %+v", self)

	case "unpauseOnSuccessfulDeploy", "UnpauseOnSuccessfulDeploy":
		return *self.UnpauseOnSuccessfulDeploy, nil
		return nil, fmt.Errorf("Field UnpauseOnSuccessfulDeploy no set on UnpauseOnSuccessfulDeploy %+v", self)

	}
}

func (self *SingularityDeployRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployRequest", name)

	case "deploy", "Deploy":
		self.Deploy = nil

	case "message", "Message":
		self.Message = nil

	case "updatedRequest", "UpdatedRequest":
		self.UpdatedRequest = nil

	case "unpauseOnSuccessfulDeploy", "UnpauseOnSuccessfulDeploy":
		self.UnpauseOnSuccessfulDeploy = nil

	}

	return nil
}

func (self *SingularityDeployRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployRequestList []*SingularityDeployRequest

func (self *SingularityDeployRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployRequestList cannot copy the values from %#v", other)
}

func (list *SingularityDeployRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
