# SetCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnId** | Pointer to **string** |  | [optional] 
**ReferencedIds** | Pointer to **[]string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSetCell

`func NewSetCell() *SetCell`

NewSetCell instantiates a new SetCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCellWithDefaults

`func NewSetCellWithDefaults() *SetCell`

NewSetCellWithDefaults instantiates a new SetCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnId

`func (o *SetCell) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *SetCell) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *SetCell) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *SetCell) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetReferencedIds

`func (o *SetCell) GetReferencedIds() []string`

GetReferencedIds returns the ReferencedIds field if non-nil, zero value otherwise.

### GetReferencedIdsOk

`func (o *SetCell) GetReferencedIdsOk() (*[]string, bool)`

GetReferencedIdsOk returns a tuple with the ReferencedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedIds

`func (o *SetCell) SetReferencedIds(v []string)`

SetReferencedIds sets ReferencedIds field to given value.

### HasReferencedIds

`func (o *SetCell) HasReferencedIds() bool`

HasReferencedIds returns a boolean if a field has been set.

### GetValue

`func (o *SetCell) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetCell) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SetCell) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SetCell) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

