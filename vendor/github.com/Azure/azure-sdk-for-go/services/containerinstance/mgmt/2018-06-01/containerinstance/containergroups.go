package containerinstance

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ContainerGroupsClient is the client for the ContainerGroups methods of the Containerinstance service.
type ContainerGroupsClient struct {
	BaseClient
}

// NewContainerGroupsClient creates an instance of the ContainerGroupsClient client.
func NewContainerGroupsClient(subscriptionID string) ContainerGroupsClient {
	return NewContainerGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContainerGroupsClientWithBaseURI creates an instance of the ContainerGroupsClient client.
func NewContainerGroupsClientWithBaseURI(baseURI string, subscriptionID string) ContainerGroupsClient {
	return ContainerGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update container groups with specified configurations.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
// containerGroup - the properties of the container group to be created or updated.
func (client ContainerGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup ContainerGroup) (result ContainerGroupsCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: containerGroup,
			Constraints: []validation.Constraint{{Target: "containerGroup.ContainerGroupProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "containerGroup.ContainerGroupProperties.Containers", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "containerGroup.ContainerGroupProperties.IPAddress", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "containerGroup.ContainerGroupProperties.IPAddress.Ports", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "containerGroup.ContainerGroupProperties.IPAddress.Type", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "containerGroup.ContainerGroupProperties.Diagnostics", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "containerGroup.ContainerGroupProperties.Diagnostics.LogAnalytics", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "containerGroup.ContainerGroupProperties.Diagnostics.LogAnalytics.WorkspaceID", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "containerGroup.ContainerGroupProperties.Diagnostics.LogAnalytics.WorkspaceKey", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("containerinstance.ContainerGroupsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, containerGroupName, containerGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ContainerGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup ContainerGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", pathParameters),
		autorest.WithJSON(containerGroup),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerGroupsClient) CreateOrUpdateSender(req *http.Request) (future ContainerGroupsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ContainerGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result ContainerGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the specified container group in the specified subscription and resource group. The operation does not
// delete other resources provided by the user, such as volumes.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
func (client ContainerGroupsClient) Delete(ctx context.Context, resourceGroupName string, containerGroupName string) (result ContainerGroup, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, containerGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ContainerGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, containerGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ContainerGroupsClient) DeleteResponder(resp *http.Response) (result ContainerGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the properties of the specified container group in the specified subscription and resource group. The
// operation returns the properties of each container group including containers, image registry credentials, restart
// policy, IP address type, OS type, state, and volumes.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
func (client ContainerGroupsClient) Get(ctx context.Context, resourceGroupName string, containerGroupName string) (result ContainerGroup, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, containerGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ContainerGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, containerGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ContainerGroupsClient) GetResponder(resp *http.Response) (result ContainerGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get a list of container groups in the specified subscription. This operation returns properties of each
// container group including containers, image registry credentials, restart policy, IP address type, OS type, state,
// and volumes.
func (client ContainerGroupsClient) List(ctx context.Context) (result ContainerGroupListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.cglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ContainerGroupsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/containerGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerGroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ContainerGroupsClient) ListResponder(resp *http.Response) (result ContainerGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ContainerGroupsClient) listNextResults(lastResults ContainerGroupListResult) (result ContainerGroupListResult, err error) {
	req, err := lastResults.containerGroupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ContainerGroupsClient) ListComplete(ctx context.Context) (result ContainerGroupListResultIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup get a list of container groups in a specified subscription and resource group. This operation
// returns properties of each container group including containers, image registry credentials, restart policy, IP
// address type, OS type, state, and volumes.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client ContainerGroupsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ContainerGroupListResultPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.cglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.cglr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ContainerGroupsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerGroupsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ContainerGroupsClient) ListByResourceGroupResponder(resp *http.Response) (result ContainerGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ContainerGroupsClient) listByResourceGroupNextResults(lastResults ContainerGroupListResult) (result ContainerGroupListResult, err error) {
	req, err := lastResults.containerGroupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ContainerGroupsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ContainerGroupListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Restart restarts all containers in a contaienr group in place. If container image has updates, new image will be
// downloaded.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
func (client ContainerGroupsClient) Restart(ctx context.Context, resourceGroupName string, containerGroupName string) (result ContainerGroupsRestartFuture, err error) {
	req, err := client.RestartPreparer(ctx, resourceGroupName, containerGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Restart", nil, "Failure preparing request")
		return
	}

	result, err = client.RestartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Restart", result.Response(), "Failure sending request")
		return
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client ContainerGroupsClient) RestartPreparer(ctx context.Context, resourceGroupName string, containerGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerGroupsClient) RestartSender(req *http.Request) (future ContainerGroupsRestartFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client ContainerGroupsClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops all containers in a contaienr group. Compute resources will be deallocated and billing will stop.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
func (client ContainerGroupsClient) Stop(ctx context.Context, resourceGroupName string, containerGroupName string) (result autorest.Response, err error) {
	req, err := client.StopPreparer(ctx, resourceGroupName, containerGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Stop", resp, "Failure sending request")
		return
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Stop", resp, "Failure responding to request")
	}

	return
}

// StopPreparer prepares the Stop request.
func (client ContainerGroupsClient) StopPreparer(ctx context.Context, resourceGroupName string, containerGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerGroupsClient) StopSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client ContainerGroupsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates container group tags with specified values.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
// resource - the container group resource with just the tags to be updated.
func (client ContainerGroupsClient) Update(ctx context.Context, resourceGroupName string, containerGroupName string, resource Resource) (result ContainerGroup, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, containerGroupName, resource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ContainerGroupsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, containerGroupName string, resource Resource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", pathParameters),
		autorest.WithJSON(resource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ContainerGroupsClient) UpdateResponder(resp *http.Response) (result ContainerGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
