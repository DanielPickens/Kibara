package singularity

import "github.com/danielpickens/kibara/dtos"

func (client *Client) DTORequest(endpoint string, response interface{}, method string, path string, pathParams map[string]interface{}, queryParams map[string]interface{}, body interface{}) (err error) {
	// Implementation of the DTORequest method goes here

	DTORequest(endpoint, response, method, path, pathParams, queryParams, body)
	if err != nil {
		return err
	}
	for _, interceptor := range client.interceptors {
		err = interceptor.OnRequest(endpoint, method, path, pathParams, queryParams, body)
		if err != nil {
			return err
		}
	}
	return nil

}

func (client *Client) Deploy(body *dtos.SingularityDeployRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-deploy", response, "POST", "/api/deploys", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetPendingDeploys() (response dtos.SingularityPendingDeployList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityPendingDeployList, 0)
	err = client.DTORequest("singularity-getpendingdeploys", &response, "GET", "/api/deploys/pending", pathParamMap, queryParamMap)

	return
}
func (client *Client) CancelDeploy(requestId string, deployId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-canceldeploy", response, "DELETE", "/api/deploys/deploy/{deployId}/request/{requestId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) UpdatePendingDeploy(body *dtos.SingularityUpdatePendingDeployRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-updatependingdeploy", response, "POST", "/api/deploys/update", pathParamMap, queryParamMap, body)

	return
}
