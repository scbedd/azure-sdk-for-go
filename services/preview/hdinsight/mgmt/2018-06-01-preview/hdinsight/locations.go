package hdinsight

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LocationsClient is the hDInsight Management Client
type LocationsClient struct {
	BaseClient
}

// NewLocationsClient creates an instance of the LocationsClient client.
func NewLocationsClient(subscriptionID string) LocationsClient {
	return NewLocationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationsClientWithBaseURI creates an instance of the LocationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return LocationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetCapabilities gets the capabilities for the specified location.
// Parameters:
// location - the Azure location (region) for which to make the request.
func (client LocationsClient) GetCapabilities(ctx context.Context, location string) (result CapabilitiesResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.GetCapabilities")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetCapabilitiesPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "GetCapabilities", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCapabilitiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "GetCapabilities", resp, "Failure sending request")
		return
	}

	result, err = client.GetCapabilitiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "GetCapabilities", resp, "Failure responding to request")
	}

	return
}

// GetCapabilitiesPreparer prepares the GetCapabilities request.
func (client LocationsClient) GetCapabilitiesPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/capabilities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCapabilitiesSender sends the GetCapabilities request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) GetCapabilitiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetCapabilitiesResponder handles the response to the GetCapabilities request. The method always
// closes the http.Response Body.
func (client LocationsClient) GetCapabilitiesResponder(resp *http.Response) (result CapabilitiesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBillingSpecs lists the billingSpecs for the specified subscription and location.
// Parameters:
// location - the Azure location (region) for which to make the request.
func (client LocationsClient) ListBillingSpecs(ctx context.Context, location string) (result BillingResponseListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.ListBillingSpecs")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListBillingSpecsPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "ListBillingSpecs", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBillingSpecsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "ListBillingSpecs", resp, "Failure sending request")
		return
	}

	result, err = client.ListBillingSpecsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "ListBillingSpecs", resp, "Failure responding to request")
	}

	return
}

// ListBillingSpecsPreparer prepares the ListBillingSpecs request.
func (client LocationsClient) ListBillingSpecsPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/billingSpecs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBillingSpecsSender sends the ListBillingSpecs request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) ListBillingSpecsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBillingSpecsResponder handles the response to the ListBillingSpecs request. The method always
// closes the http.Response Body.
func (client LocationsClient) ListBillingSpecsResponder(resp *http.Response) (result BillingResponseListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListUsages lists the usages for the specified location.
// Parameters:
// location - the Azure location (region) for which to make the request.
func (client LocationsClient) ListUsages(ctx context.Context, location string) (result UsagesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.ListUsages")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListUsagesPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "ListUsages", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListUsagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "ListUsages", resp, "Failure sending request")
		return
	}

	result, err = client.ListUsagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.LocationsClient", "ListUsages", resp, "Failure responding to request")
	}

	return
}

// ListUsagesPreparer prepares the ListUsages request.
func (client LocationsClient) ListUsagesPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListUsagesSender sends the ListUsages request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) ListUsagesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListUsagesResponder handles the response to the ListUsages request. The method always
// closes the http.Response Body.
func (client LocationsClient) ListUsagesResponder(resp *http.Response) (result UsagesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
