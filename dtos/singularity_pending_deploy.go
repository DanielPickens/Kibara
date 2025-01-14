package dtos

import (
	"fmt"
	"io"

	"github.com/kibara/swaggering"
)

type SingularityPendingDeployDeployState string

const (
	SingularityPendingDeployDeployStateSUCCEEDED             SingularityPendingDeployDeployState = "SUCCEEDED"
	SingularityPendingDeployDeployStateFAILED_INTERNAL_STATE SingularityPendingDeployDeployState = "FAILED_INTERNAL_STATE"
	SingularityPendingDeployDeployStateCANCELING             SingularityPendingDeployDeployState = "CANCELING"
	SingularityPendingDeployDeployStateWAITING               SingularityPendingDeployDeployState = "WAITING"
	SingularityPendingDeployDeployStateOVERDUE               SingularityPendingDeployDeployState = "OVERDUE"
	SingularityPendingDeployDeployStateFAILED                SingularityPendingDeployDeployState = "FAILED"
	SingularityPendingDeployDeployStateCANCELED              SingularityPendingDeployDeployState = "CANCELED"
)

type SingularityPendingDeploy struct {
	DeployMarker           *SingularityDeployMarker             `json:"deployMarker,omitempty"`
	LastLoadBalancerUpdate *SingularityLoadBalancerUpdate       `json:"lastLoadBalancerUpdate,omitempty"`
	CurrentDeployState     *SingularityPendingDeployDeployState `json:"currentDeployState,omitempty"`
	DeployProgress         *SingularityDeployProgress           `json:"deployProgress,omitempty"`
	UpdatedRequest         *SingularityRequest                  `json:"updatedRequest,omitempty"`
}

func (self *SingularityPendingDeploy) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPendingDeploy) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingDeploy); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingDeploy cannot copy the values from %#v", other)
}

func (self *SingularityPendingDeploy) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPendingDeploy) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPendingDeploy) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingDeploy", name)

	case "deployMarker", "DeployMarker":
		v, ok := value.(*SingularityDeployMarker)
		if ok {
			self.DeployMarker = v
			return nil
		}
		return fmt.Errorf("Field deployMarker/DeployMarker: value %v(%T) couldn't be cast to type *SingularityDeployMarker", value, value)

	case "lastLoadBalancerUpdate", "LastLoadBalancerUpdate":
		v, ok := value.(*SingularityLoadBalancerUpdate)
		if ok {
			self.LastLoadBalancerUpdate = v
			return nil
		}
		return fmt.Errorf("Field lastLoadBalancerUpdate/LastLoadBalancerUpdate: value %v(%T) couldn't be cast to type *SingularityLoadBalancerUpdate", value, value)

	case "currentDeployState", "CurrentDeployState":
		v, ok := value.(SingularityPendingDeployDeployState)
		if ok {
			self.CurrentDeployState = &v
			return nil
		}
		return fmt.Errorf("Field currentDeployState/CurrentDeployState: value %v(%T) couldn't be cast to type SingularityPendingDeployDeployState", value, value)

	case "deployProgress", "DeployProgress":
		v, ok := value.(*SingularityDeployProgress)
		if ok {
			self.DeployProgress = v
			return nil
		}
		return fmt.Errorf("Field deployProgress/DeployProgress: value %v(%T) couldn't be cast to type *SingularityDeployProgress", value, value)

	case "updatedRequest", "UpdatedRequest":
		v, ok := value.(*SingularityRequest)
		if ok {
			self.UpdatedRequest = v
			return nil
		}
		return fmt.Errorf("Field updatedRequest/UpdatedRequest: value %v(%T) couldn't be cast to type *SingularityRequest", value, value)

	}
}

func (self *SingularityPendingDeploy) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPendingDeploy", name)

	case "deployMarker", "DeployMarker":
		return self.DeployMarker, nil
		return nil, fmt.Errorf("Field DeployMarker no set on DeployMarker %+v", self)

	case "lastLoadBalancerUpdate", "LastLoadBalancerUpdate":
		return self.LastLoadBalancerUpdate, nil
		return nil, fmt.Errorf("Field LastLoadBalancerUpdate no set on LastLoadBalancerUpdate %+v", self)

	case "currentDeployState", "CurrentDeployState":
		return *self.CurrentDeployState, nil
		return nil, fmt.Errorf("Field CurrentDeployState no set on CurrentDeployState %+v", self)

	case "deployProgress", "DeployProgress":
		return self.DeployProgress, nil
		return nil, fmt.Errorf("Field DeployProgress no set on DeployProgress %+v", self)

	case "updatedRequest", "UpdatedRequest":
		return self.UpdatedRequest, nil
		return nil, fmt.Errorf("Field UpdatedRequest no set on UpdatedRequest %+v", self)

	}
}

func (self *SingularityPendingDeploy) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingDeploy", name)

	case "deployMarker", "DeployMarker":
		self.DeployMarker = nil

	case "lastLoadBalancerUpdate", "LastLoadBalancerUpdate":
		self.LastLoadBalancerUpdate = nil

	case "currentDeployState", "CurrentDeployState":
		self.CurrentDeployState = nil

	case "deployProgress", "DeployProgress":
		self.DeployProgress = nil

	case "updatedRequest", "UpdatedRequest":
		self.UpdatedRequest = nil

	}

	return nil
}

func (self *SingularityPendingDeploy) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPendingDeployList []*SingularityPendingDeploy

func (self *SingularityPendingDeployList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingDeployList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingDeployList cannot copy the values from %#v", other)
}

func (list *SingularityPendingDeployList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingDeployList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingDeployList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
