# Go API client for gridly

Gridly API documentation

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.29.1
- Package version: 1.2.11
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.gridly.com](https://www.gridly.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import gridly "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), gridly.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), gridly.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), gridly.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), gridly.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.gridly.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BranchApi* | [**Create**](docs/BranchApi.md#create) | **Post** /v1/branches | create
*BranchApi* | [**CreateDiffCheck**](docs/BranchApi.md#creatediffcheck) | **Post** /v1/branches/diffcheck | createDiffCheck
*BranchApi* | [**Delete**](docs/BranchApi.md#delete) | **Delete** /v1/branches/{branchId} | delete
*BranchApi* | [**Get**](docs/BranchApi.md#get) | **Get** /v1/branches/{branchId} | get
*BranchApi* | [**GetDiffCheck**](docs/BranchApi.md#getdiffcheck) | **Get** /v1/branches/diffcheck/{taskId} | getDiffCheck
*BranchApi* | [**List**](docs/BranchApi.md#list) | **Get** /v1/branches | list
*BranchApi* | [**Merge**](docs/BranchApi.md#merge) | **Post** /v1/branches/{branchId}/merge | merge
*DatabaseApi* | [**Create**](docs/DatabaseApi.md#create) | **Post** /v1/databases | create
*DatabaseApi* | [**Delete**](docs/DatabaseApi.md#delete) | **Delete** /v1/databases/{dbId} | delete
*DatabaseApi* | [**Duplicate**](docs/DatabaseApi.md#duplicate) | **Post** /v1/databases/{dbId}/duplicate | duplicate
*DatabaseApi* | [**Get**](docs/DatabaseApi.md#get) | **Get** /v1/databases/{dbId} | get
*DatabaseApi* | [**List**](docs/DatabaseApi.md#list) | **Get** /v1/databases | list
*DatabaseApi* | [**Update**](docs/DatabaseApi.md#update) | **Put** /v1/databases/{dbId} | update
*GlossaryApi* | [**Create**](docs/GlossaryApi.md#create) | **Post** /v1/glossaries | Create a new glossary
*GlossaryApi* | [**Delete**](docs/GlossaryApi.md#delete) | **Delete** /v1/glossaries/{id} | Delete a glossary
*GlossaryApi* | [**ExportFile**](docs/GlossaryApi.md#exportfile) | **Get** /v1/glossaries/{id}/export | Export a glossary
*GlossaryApi* | [**Get**](docs/GlossaryApi.md#get) | **Get** /v1/glossaries/{id} | get glossary info
*GlossaryApi* | [**GetAll**](docs/GlossaryApi.md#getall) | **Get** /v1/glossaries | List all glossaries
*GlossaryApi* | [**ImportFile**](docs/GlossaryApi.md#importfile) | **Post** /v1/glossaries/{id}/import | Import a glossary from file
*GlossaryApi* | [**Update**](docs/GlossaryApi.md#update) | **Put** /v1/glossaries/{id} | Update glossary info
*GridApi* | [**Create**](docs/GridApi.md#create) | **Post** /v1/grids | create
*GridApi* | [**CreateCategory**](docs/GridApi.md#createcategory) | **Post** /v1/grids/{gridId}/settings/categories | createCategory
*GridApi* | [**Delete**](docs/GridApi.md#delete) | **Delete** /v1/grids/{gridId} | delete
*GridApi* | [**DeleteCategory**](docs/GridApi.md#deletecategory) | **Delete** /v1/grids/{gridId}/settings/categories/{categoryId} | deleteCategory
*GridApi* | [**DeleteFile**](docs/GridApi.md#deletefile) | **Delete** /v1/grids/{gridId}/settings/categories/{categoryId}/files/{fileId} | deleteFile
*GridApi* | [**Get**](docs/GridApi.md#get) | **Get** /v1/grids/{gridId} | get
*GridApi* | [**GetSetting**](docs/GridApi.md#getsetting) | **Get** /v1/grids/{gridId}/settings | getSetting
*GridApi* | [**List**](docs/GridApi.md#list) | **Get** /v1/grids | list
*GridApi* | [**ListFiles**](docs/GridApi.md#listfiles) | **Get** /v1/grids/{gridId}/settings/files | listFiles
*GridApi* | [**ListTemplateGrids**](docs/GridApi.md#listtemplategrids) | **Get** /v1/template-grids | listTemplateGrids
*GridApi* | [**Update**](docs/GridApi.md#update) | **Patch** /v1/grids/{gridId} | update
*GridApi* | [**UpdateCategory**](docs/GridApi.md#updatecategory) | **Put** /v1/grids/{gridId}/settings/categories/{categoryId} | updateCategory
*GridApi* | [**UpdateSetting**](docs/GridApi.md#updatesetting) | **Patch** /v1/grids/{gridId}/settings | updateSetting
*GridApi* | [**UploadSettingFile**](docs/GridApi.md#uploadsettingfile) | **Post** /v1/grids/{gridId}/settings/categories/{categoryId}/files | uploadSettingFile
*PathApi* | [**Create**](docs/PathApi.md#create) | **Post** /v1/views/{viewId}/paths | create
*PathApi* | [**Delete**](docs/PathApi.md#delete) | **Delete** /v1/views/{viewId}/paths | delete
*PathApi* | [**List**](docs/PathApi.md#list) | **Get** /v1/views/{viewId}/paths/tree | list
*PathApi* | [**Move**](docs/PathApi.md#move) | **Post** /v1/views/{viewId}/paths/move | move
*PathApi* | [**Update**](docs/PathApi.md#update) | **Put** /v1/views/{viewId}/paths/{path} | update
*ProjectApi* | [**Create**](docs/ProjectApi.md#create) | **Post** /v1/projects | create
*ProjectApi* | [**Delete**](docs/ProjectApi.md#delete) | **Delete** /v1/projects/{projectId} | delete
*ProjectApi* | [**FindOne**](docs/ProjectApi.md#findone) | **Get** /v1/projects/{projectId} | findOne
*ProjectApi* | [**List**](docs/ProjectApi.md#list) | **Get** /v1/projects | list
*ProjectApi* | [**Update**](docs/ProjectApi.md#update) | **Put** /v1/projects/{projectId} | update
*RecordApi* | [**Create**](docs/RecordApi.md#create) | **Post** /v1/views/{viewId}/records | create
*RecordApi* | [**Delete**](docs/RecordApi.md#delete) | **Delete** /v1/views/{viewId}/records | delete
*RecordApi* | [**Fetch**](docs/RecordApi.md#fetch) | **Get** /v1/views/{viewId}/records | fetch
*RecordApi* | [**FetchHistories**](docs/RecordApi.md#fetchhistories) | **Get** /v1/views/{viewId}/records/{recordId}/histories | fetchHistories
*RecordApi* | [**Update**](docs/RecordApi.md#update) | **Patch** /v1/views/{viewId}/records | update
*RecordApi* | [**UpdateRecord**](docs/RecordApi.md#updaterecord) | **Patch** /v1/views/{viewId}/records/{id} | updateRecord
*ShareViewApi* | [**Create**](docs/ShareViewApi.md#create) | **Put** /v1/views/{viewId}/shares | create
*ShareViewApi* | [**Delete**](docs/ShareViewApi.md#delete) | **Delete** /v1/views/{viewId}/shares | delete
*ShareViewApi* | [**Get**](docs/ShareViewApi.md#get) | **Get** /v1/views/{viewId}/shares | get
*TaskApi* | [**Get**](docs/TaskApi.md#get) | **Get** /v1/tasks/{taskId} | get
*TransmemApi* | [**Cleanup**](docs/TransmemApi.md#cleanup) | **Put** /v1/transmems/{tmId}/cleanup | Erases all the translation data of the provided tmId
*TransmemApi* | [**Create**](docs/TransmemApi.md#create) | **Post** /v1/transmems | Create a new translation memory
*TransmemApi* | [**CreateWithFile**](docs/TransmemApi.md#createwithfile) | **Post** /v1/transmems/upload | Create a new translation memory by uploading tmx file
*TransmemApi* | [**Delete**](docs/TransmemApi.md#delete) | **Delete** /v1/transmems/{tmId} | Delete a translation memory by id
*TransmemApi* | [**Export**](docs/TransmemApi.md#export) | **Get** /v1/transmems/{tmId}/export | Export translation memory tmx file
*TransmemApi* | [**Get**](docs/TransmemApi.md#get) | **Get** /v1/transmems/{tmId} | Get translation memory info by id
*TransmemApi* | [**ImportTmx**](docs/TransmemApi.md#importtmx) | **Post** /v1/transmems/{tmId}/import | Import a translation memory from tmx file
*TransmemApi* | [**ListTM**](docs/TransmemApi.md#listtm) | **Get** /v1/transmems | List all available translation memories or create default one if there is no translation memory
*TransmemApi* | [**Update**](docs/TransmemApi.md#update) | **Put** /v1/transmems/{tmId} | Update a translation memory
*ViewApi* | [**Create**](docs/ViewApi.md#create) | **Post** /v1/views | create
*ViewApi* | [**Export**](docs/ViewApi.md#export) | **Get** /v1/views/{viewId}/export | export
*ViewApi* | [**Get**](docs/ViewApi.md#get) | **Get** /v1/views/{viewId} | get
*ViewApi* | [**GetStatistic**](docs/ViewApi.md#getstatistic) | **Get** /v1/views/{viewId}/statistic | getStatistic
*ViewApi* | [**ImportView**](docs/ViewApi.md#importview) | **Post** /v1/views/{viewId}/import | importView
*ViewApi* | [**List**](docs/ViewApi.md#list) | **Get** /v1/views | list
*ViewApi* | [**Merge**](docs/ViewApi.md#merge) | **Post** /v1/views/{viewId}/merge | merge
*ViewColumnApi* | [**Add**](docs/ViewColumnApi.md#add) | **Post** /v1/views/{viewId}/columns/{columnId}/add | add
*ViewColumnApi* | [**BulkCreate**](docs/ViewColumnApi.md#bulkcreate) | **Post** /v1/views/{viewId}/columns/bulk | bulkCreate
*ViewColumnApi* | [**Create**](docs/ViewColumnApi.md#create) | **Post** /v1/views/{viewId}/columns | create
*ViewColumnApi* | [**Delete**](docs/ViewColumnApi.md#delete) | **Delete** /v1/views/{viewId}/columns/{columnId} | delete
*ViewColumnApi* | [**Get**](docs/ViewColumnApi.md#get) | **Get** /v1/views/{viewId}/columns/{columnId} | get
*ViewColumnApi* | [**Remove**](docs/ViewColumnApi.md#remove) | **Post** /v1/views/{viewId}/columns/{columnId}/remove | remove
*ViewColumnApi* | [**Update**](docs/ViewColumnApi.md#update) | **Patch** /v1/views/{viewId}/columns/{columnId} | update
*ViewDependencyApi* | [**Create**](docs/ViewDependencyApi.md#create) | **Post** /v1/views/{viewId}/dependencies | create
*ViewDependencyApi* | [**Delete**](docs/ViewDependencyApi.md#delete) | **Delete** /v1/views/{viewId}/dependencies | delete
*ViewDependencyApi* | [**DeleteById**](docs/ViewDependencyApi.md#deletebyid) | **Delete** /v1/views/{viewId}/dependencies/{dependencyId} | deleteById
*ViewDependencyApi* | [**Get**](docs/ViewDependencyApi.md#get) | **Get** /v1/views/{viewId}/dependencies/{dependencyId} | get
*ViewDependencyApi* | [**List**](docs/ViewDependencyApi.md#list) | **Get** /v1/views/{viewId}/dependencies | list
*ViewDependencyApi* | [**Update**](docs/ViewDependencyApi.md#update) | **Put** /v1/views/{viewId}/dependencies/{dependencyId} | update
*ViewFileApi* | [**Delete**](docs/ViewFileApi.md#delete) | **Delete** /v1/views/{viewId}/files | delete
*ViewFileApi* | [**Download**](docs/ViewFileApi.md#download) | **Get** /v1/views/{viewId}/files/{fileId} | download
*ViewFileApi* | [**Upload**](docs/ViewFileApi.md#upload) | **Post** /v1/views/{viewId}/files | upload
*ViewFileApi* | [**UploadZip**](docs/ViewFileApi.md#uploadzip) | **Post** /v1/views/{viewId}/files/zip | uploadZip


## Documentation For Models

 - [AddViewColumn](docs/AddViewColumn.md)
 - [Branch](docs/Branch.md)
 - [BranchDiffCell](docs/BranchDiffCell.md)
 - [BranchDiffRecord](docs/BranchDiffRecord.md)
 - [Cell](docs/Cell.md)
 - [CellHistory](docs/CellHistory.md)
 - [ColumnReference](docs/ColumnReference.md)
 - [ColumnStatistic](docs/ColumnStatistic.md)
 - [CreateBranch](docs/CreateBranch.md)
 - [CreateColumn](docs/CreateColumn.md)
 - [CreateDatabase](docs/CreateDatabase.md)
 - [CreateDependency](docs/CreateDependency.md)
 - [CreateFileCategory](docs/CreateFileCategory.md)
 - [CreateGlossary](docs/CreateGlossary.md)
 - [CreateGrid](docs/CreateGrid.md)
 - [CreatePath](docs/CreatePath.md)
 - [CreateProject](docs/CreateProject.md)
 - [CreateShareView](docs/CreateShareView.md)
 - [CreateTransMem](docs/CreateTransMem.md)
 - [CreateView](docs/CreateView.md)
 - [Database](docs/Database.md)
 - [DateTimeFormat](docs/DateTimeFormat.md)
 - [DeleteDependency](docs/DeleteDependency.md)
 - [DeleteFile](docs/DeleteFile.md)
 - [DeletePath](docs/DeletePath.md)
 - [DeleteRecord](docs/DeleteRecord.md)
 - [Dependency](docs/Dependency.md)
 - [ExportFileHeader](docs/ExportFileHeader.md)
 - [ExportFormat](docs/ExportFormat.md)
 - [FetchFileOption](docs/FetchFileOption.md)
 - [FetchRecordHistoryRequest](docs/FetchRecordHistoryRequest.md)
 - [FileCategory](docs/FileCategory.md)
 - [FilterField](docs/FilterField.md)
 - [Formula](docs/Formula.md)
 - [Glossary](docs/Glossary.md)
 - [GlossaryExportFormat](docs/GlossaryExportFormat.md)
 - [GlossaryProject](docs/GlossaryProject.md)
 - [Grid](docs/Grid.md)
 - [GridSetting](docs/GridSetting.md)
 - [Group](docs/Group.md)
 - [ImportOption](docs/ImportOption.md)
 - [MergeBranchRequest](docs/MergeBranchRequest.md)
 - [MergeCellConflict](docs/MergeCellConflict.md)
 - [MergeRecordConflict](docs/MergeRecordConflict.md)
 - [MovePath](docs/MovePath.md)
 - [NumberFormat](docs/NumberFormat.md)
 - [PathList](docs/PathList.md)
 - [PathNode](docs/PathNode.md)
 - [PathSingle](docs/PathSingle.md)
 - [Privilege](docs/Privilege.md)
 - [Project](docs/Project.md)
 - [ProjectDetail](docs/ProjectDetail.md)
 - [Record](docs/Record.md)
 - [RecordHistory](docs/RecordHistory.md)
 - [RecordIdentifierWrapper](docs/RecordIdentifierWrapper.md)
 - [Reference](docs/Reference.md)
 - [ReferencedColumn](docs/ReferencedColumn.md)
 - [ReferencedGrid](docs/ReferencedGrid.md)
 - [Role](docs/Role.md)
 - [SetCell](docs/SetCell.md)
 - [SetRecord](docs/SetRecord.md)
 - [SettingFile](docs/SettingFile.md)
 - [ShareView](docs/ShareView.md)
 - [Task](docs/Task.md)
 - [TransMem](docs/TransMem.md)
 - [TranslationCount](docs/TranslationCount.md)
 - [TranslationStatus](docs/TranslationStatus.md)
 - [UpdateCategory](docs/UpdateCategory.md)
 - [UpdateColumn](docs/UpdateColumn.md)
 - [UpdateDatabase](docs/UpdateDatabase.md)
 - [UpdateDependency](docs/UpdateDependency.md)
 - [UpdateGlossary](docs/UpdateGlossary.md)
 - [UpdateGrid](docs/UpdateGrid.md)
 - [UpdateGridSetting](docs/UpdateGridSetting.md)
 - [UpdatePath](docs/UpdatePath.md)
 - [UpdateProject](docs/UpdateProject.md)
 - [UpdateTransMem](docs/UpdateTransMem.md)
 - [UploadSettingFileRequest](docs/UploadSettingFileRequest.md)
 - [UploadZipRequest](docs/UploadZipRequest.md)
 - [UploadedFile](docs/UploadedFile.md)
 - [View](docs/View.md)
 - [ViewColumn](docs/ViewColumn.md)
 - [ViewStatistic](docs/ViewStatistic.md)


## Documentation For Authorization



### ApiKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@gridly.com

