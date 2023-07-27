/*
Gridly API

Gridly API documentation

API version: 4.15.1
Contact: support@gridly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// RecordApiService RecordApi service
type RecordApiService service

type RecordApiCreateRequest struct {
	ctx context.Context
	ApiService *RecordApiService
	viewId string
	createRecords *[]SetRecord
}

// createRecords
func (r RecordApiCreateRequest) CreateRecords(createRecords []SetRecord) RecordApiCreateRequest {
	r.createRecords = &createRecords
	return r
}

func (r RecordApiCreateRequest) Execute() ([]Record, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create create

create

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return RecordApiCreateRequest
*/
func (a *RecordApiService) Create(ctx context.Context, viewId string) RecordApiCreateRequest {
	return RecordApiCreateRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return []Record
func (a *RecordApiService) CreateExecute(r RecordApiCreateRequest) ([]Record, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Record
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordApiService.Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", url.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createRecords == nil {
		return localVarReturnValue, nil, reportError("createRecords is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createRecords
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordApiDeleteRequest struct {
	ctx context.Context
	ApiService *RecordApiService
	viewId string
	deleteRecord *DeleteRecord
}

func (r RecordApiDeleteRequest) DeleteRecord(deleteRecord DeleteRecord) RecordApiDeleteRequest {
	r.deleteRecord = &deleteRecord
	return r
}

func (r RecordApiDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete delete

delete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return RecordApiDeleteRequest
*/
func (a *RecordApiService) Delete(ctx context.Context, viewId string) RecordApiDeleteRequest {
	return RecordApiDeleteRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
func (a *RecordApiService) DeleteExecute(r RecordApiDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordApiService.Delete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", url.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deleteRecord == nil {
		return nil, reportError("deleteRecord is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deleteRecord
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RecordApiFetchRequest struct {
	ctx context.Context
	ApiService *RecordApiService
	viewId string
	columnIds *[]string
	page *string
	query *string
	sort *string
	fetchFileOption *FetchFileOption
	afterRecordId *string
	beforeRecordId *string
}

// columnIds
func (r RecordApiFetchRequest) ColumnIds(columnIds []string) RecordApiFetchRequest {
	r.columnIds = &columnIds
	return r
}

// page
func (r RecordApiFetchRequest) Page(page string) RecordApiFetchRequest {
	r.page = &page
	return r
}

// query
func (r RecordApiFetchRequest) Query(query string) RecordApiFetchRequest {
	r.query = &query
	return r
}

// sort
func (r RecordApiFetchRequest) Sort(sort string) RecordApiFetchRequest {
	r.sort = &sort
	return r
}

// fetchFileOption
func (r RecordApiFetchRequest) FetchFileOption(fetchFileOption FetchFileOption) RecordApiFetchRequest {
	r.fetchFileOption = &fetchFileOption
	return r
}

// afterRecordId
func (r RecordApiFetchRequest) AfterRecordId(afterRecordId string) RecordApiFetchRequest {
	r.afterRecordId = &afterRecordId
	return r
}

// beforeRecordId
func (r RecordApiFetchRequest) BeforeRecordId(beforeRecordId string) RecordApiFetchRequest {
	r.beforeRecordId = &beforeRecordId
	return r
}

func (r RecordApiFetchRequest) Execute() ([]Record, *http.Response, error) {
	return r.ApiService.FetchExecute(r)
}

/*
Fetch fetch

fetch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return RecordApiFetchRequest
*/
func (a *RecordApiService) Fetch(ctx context.Context, viewId string) RecordApiFetchRequest {
	return RecordApiFetchRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return []Record
func (a *RecordApiService) FetchExecute(r RecordApiFetchRequest) ([]Record, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Record
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordApiService.Fetch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", url.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.columnIds != nil {
		t := *r.columnIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("columnIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("columnIds", parameterToString(t, "multi"))
		}
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.query != nil {
		localVarQueryParams.Add("query", parameterToString(*r.query, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.fetchFileOption != nil {
		localVarQueryParams.Add("fetchFileOption", parameterToString(*r.fetchFileOption, ""))
	}
	if r.afterRecordId != nil {
		localVarQueryParams.Add("afterRecordId", parameterToString(*r.afterRecordId, ""))
	}
	if r.beforeRecordId != nil {
		localVarQueryParams.Add("beforeRecordId", parameterToString(*r.beforeRecordId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordApiFetchHistoriesRequest struct {
	ctx context.Context
	ApiService *RecordApiService
	viewId string
	recordId string
	page *string
}

// page
func (r RecordApiFetchHistoriesRequest) Page(page string) RecordApiFetchHistoriesRequest {
	r.page = &page
	return r
}

func (r RecordApiFetchHistoriesRequest) Execute() ([]RecordHistory, *http.Response, error) {
	return r.ApiService.FetchHistoriesExecute(r)
}

/*
FetchHistories fetchHistories

fetchHistories

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @param recordId recordId
 @return RecordApiFetchHistoriesRequest
*/
func (a *RecordApiService) FetchHistories(ctx context.Context, viewId string, recordId string) RecordApiFetchHistoriesRequest {
	return RecordApiFetchHistoriesRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
		recordId: recordId,
	}
}

// Execute executes the request
//  @return []RecordHistory
func (a *RecordApiService) FetchHistoriesExecute(r RecordApiFetchHistoriesRequest) ([]RecordHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RecordHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordApiService.FetchHistories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/records/{recordId}/histories"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", url.PathEscape(parameterToString(r.viewId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", url.PathEscape(parameterToString(r.recordId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordApiUpdateRequest struct {
	ctx context.Context
	ApiService *RecordApiService
	viewId string
	setRecord *[]SetRecord
}

func (r RecordApiUpdateRequest) SetRecord(setRecord []SetRecord) RecordApiUpdateRequest {
	r.setRecord = &setRecord
	return r
}

func (r RecordApiUpdateRequest) Execute() ([]Record, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update update

update

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return RecordApiUpdateRequest
*/
func (a *RecordApiService) Update(ctx context.Context, viewId string) RecordApiUpdateRequest {
	return RecordApiUpdateRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return []Record
func (a *RecordApiService) UpdateExecute(r RecordApiUpdateRequest) ([]Record, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Record
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordApiService.Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", url.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setRecord == nil {
		return localVarReturnValue, nil, reportError("setRecord is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.setRecord
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordApiUpdateRecordRequest struct {
	ctx context.Context
	ApiService *RecordApiService
	id string
	viewId string
	setRecord *SetRecord
	path *string
}

func (r RecordApiUpdateRecordRequest) SetRecord(setRecord SetRecord) RecordApiUpdateRecordRequest {
	r.setRecord = &setRecord
	return r
}

// path
func (r RecordApiUpdateRecordRequest) Path(path string) RecordApiUpdateRecordRequest {
	r.path = &path
	return r
}

func (r RecordApiUpdateRecordRequest) Execute() (*Record, *http.Response, error) {
	return r.ApiService.UpdateRecordExecute(r)
}

/*
UpdateRecord updateRecord

updateRecord

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id
 @param viewId viewId
 @return RecordApiUpdateRecordRequest
*/
func (a *RecordApiService) UpdateRecord(ctx context.Context, id string, viewId string) RecordApiUpdateRecordRequest {
	return RecordApiUpdateRecordRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return Record
func (a *RecordApiService) UpdateRecordExecute(r RecordApiUpdateRecordRequest) (*Record, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Record
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordApiService.UpdateRecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/records/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", url.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setRecord == nil {
		return localVarReturnValue, nil, reportError("setRecord is required and must be specified")
	}

	if r.path != nil {
		localVarQueryParams.Add("path", parameterToString(*r.path, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.setRecord
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
