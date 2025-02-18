// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UniverseDefinitionTaskParams universe definition task params
//
// swagger:model UniverseDefinitionTaskParams
type UniverseDefinitionTaskParams struct {

	// allow insecure
	AllowInsecure bool `json:"allowInsecure,omitempty"`

	// backup in progress
	BackupInProgress bool `json:"backupInProgress,omitempty"`

	// capability
	// Enum: [READ_ONLY EDITS_ALLOWED]
	Capability string `json:"capability,omitempty"`

	// client root c a
	// Format: uuid
	ClientRootCA strfmt.UUID `json:"clientRootCA,omitempty"`

	// clusters
	// Required: true
	Clusters []*Cluster `json:"clusters"`

	// Amazon Resource Name (ARN) of the CMK
	CmkArn string `json:"cmkArn,omitempty"`

	// Communication ports
	CommunicationPorts *CommunicationPorts `json:"communicationPorts,omitempty"`

	// current cluster type
	// Enum: [PRIMARY ASYNC]
	CurrentClusterType string `json:"currentClusterType,omitempty"`

	// Device information
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`

	// Encryption at rest configation
	EncryptionAtRestConfig *EncryptionAtRestConfig `json:"encryptionAtRestConfig,omitempty"`

	// Error message
	ErrorString string `json:"errorString,omitempty"`

	// Expected universe version
	ExpectedUniverseVersion int32 `json:"expectedUniverseVersion,omitempty"`

	// Extra dependencies
	ExtraDependencies *ExtraDependencies `json:"extraDependencies,omitempty"`

	// Whether this task has been tried before
	FirstTry bool `json:"firstTry,omitempty"`

	// imported state
	// Enum: [NONE STARTED MASTERS_ADDED TSERVERS_ADDED IMPORTED]
	ImportedState string `json:"importedState,omitempty"`

	// itest s3 package path
	ItestS3PackagePath string `json:"itestS3PackagePath,omitempty"`

	// next cluster index
	NextClusterIndex int32 `json:"nextClusterIndex,omitempty"`

	// Node details
	// Unique: true
	NodeDetailsSet []*NodeDetails `json:"nodeDetailsSet"`

	// Node exporter user
	NodeExporterUser string `json:"nodeExporterUser,omitempty"`

	// node prefix
	NodePrefix string `json:"nodePrefix,omitempty"`

	// remote package path
	RemotePackagePath string `json:"remotePackagePath,omitempty"`

	// reset a z config
	ResetAZConfig bool `json:"resetAZConfig,omitempty"`

	// root and client root c a same
	RootAndClientRootCASame bool `json:"rootAndClientRootCASame,omitempty"`

	// root c a
	// Format: uuid
	RootCA strfmt.UUID `json:"rootCA,omitempty"`

	// set txn table wait count flag
	SetTxnTableWaitCountFlag bool `json:"setTxnTableWaitCountFlag,omitempty"`

	// The source universe's sync replication relationships
	// Read Only: true
	SourceAsyncReplicationRelationships []*AsyncReplicationConfig `json:"sourceAsyncReplicationRelationships"`

	// The target universe's async replication relationships
	// Read Only: true
	TargetAsyncReplicationRelationships []*AsyncReplicationConfig `json:"targetAsyncReplicationRelationships"`

	// universe paused
	UniversePaused bool `json:"universePaused,omitempty"`

	// Associated universe UUID
	// Format: uuid
	UniverseUUID strfmt.UUID `json:"universeUUID,omitempty"`

	// update in progress
	UpdateInProgress bool `json:"updateInProgress,omitempty"`

	// update succeeded
	UpdateSucceeded bool `json:"updateSucceeded,omitempty"`

	// user a z selected
	UserAZSelected bool `json:"userAZSelected,omitempty"`

	// Previous software version
	YbPrevSoftwareVersion string `json:"ybPrevSoftwareVersion,omitempty"`
}

// Validate validates this universe definition task params
func (m *UniverseDefinitionTaskParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientRootCA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommunicationPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentClusterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionAtRestConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportedState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeDetailsSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceAsyncReplicationRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetAsyncReplicationRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniverseUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var universeDefinitionTaskParamsTypeCapabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["READ_ONLY","EDITS_ALLOWED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		universeDefinitionTaskParamsTypeCapabilityPropEnum = append(universeDefinitionTaskParamsTypeCapabilityPropEnum, v)
	}
}

const (

	// UniverseDefinitionTaskParamsCapabilityREADONLY captures enum value "READ_ONLY"
	UniverseDefinitionTaskParamsCapabilityREADONLY string = "READ_ONLY"

	// UniverseDefinitionTaskParamsCapabilityEDITSALLOWED captures enum value "EDITS_ALLOWED"
	UniverseDefinitionTaskParamsCapabilityEDITSALLOWED string = "EDITS_ALLOWED"
)

// prop value enum
func (m *UniverseDefinitionTaskParams) validateCapabilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, universeDefinitionTaskParamsTypeCapabilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UniverseDefinitionTaskParams) validateCapability(formats strfmt.Registry) error {
	if swag.IsZero(m.Capability) { // not required
		return nil
	}

	// value enum
	if err := m.validateCapabilityEnum("capability", "body", m.Capability); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateClientRootCA(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientRootCA) { // not required
		return nil
	}

	if err := validate.FormatOf("clientRootCA", "body", "uuid", m.ClientRootCA.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateClusters(formats strfmt.Registry) error {

	if err := validate.Required("clusters", "body", m.Clusters); err != nil {
		return err
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateCommunicationPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.CommunicationPorts) { // not required
		return nil
	}

	if m.CommunicationPorts != nil {
		if err := m.CommunicationPorts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("communicationPorts")
			}
			return err
		}
	}

	return nil
}

var universeDefinitionTaskParamsTypeCurrentClusterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRIMARY","ASYNC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		universeDefinitionTaskParamsTypeCurrentClusterTypePropEnum = append(universeDefinitionTaskParamsTypeCurrentClusterTypePropEnum, v)
	}
}

const (

	// UniverseDefinitionTaskParamsCurrentClusterTypePRIMARY captures enum value "PRIMARY"
	UniverseDefinitionTaskParamsCurrentClusterTypePRIMARY string = "PRIMARY"

	// UniverseDefinitionTaskParamsCurrentClusterTypeASYNC captures enum value "ASYNC"
	UniverseDefinitionTaskParamsCurrentClusterTypeASYNC string = "ASYNC"
)

// prop value enum
func (m *UniverseDefinitionTaskParams) validateCurrentClusterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, universeDefinitionTaskParamsTypeCurrentClusterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UniverseDefinitionTaskParams) validateCurrentClusterType(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentClusterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrentClusterTypeEnum("currentClusterType", "body", m.CurrentClusterType); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateDeviceInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceInfo) { // not required
		return nil
	}

	if m.DeviceInfo != nil {
		if err := m.DeviceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceInfo")
			}
			return err
		}
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateEncryptionAtRestConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionAtRestConfig) { // not required
		return nil
	}

	if m.EncryptionAtRestConfig != nil {
		if err := m.EncryptionAtRestConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionAtRestConfig")
			}
			return err
		}
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateExtraDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraDependencies) { // not required
		return nil
	}

	if m.ExtraDependencies != nil {
		if err := m.ExtraDependencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extraDependencies")
			}
			return err
		}
	}

	return nil
}

var universeDefinitionTaskParamsTypeImportedStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","STARTED","MASTERS_ADDED","TSERVERS_ADDED","IMPORTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		universeDefinitionTaskParamsTypeImportedStatePropEnum = append(universeDefinitionTaskParamsTypeImportedStatePropEnum, v)
	}
}

const (

	// UniverseDefinitionTaskParamsImportedStateNONE captures enum value "NONE"
	UniverseDefinitionTaskParamsImportedStateNONE string = "NONE"

	// UniverseDefinitionTaskParamsImportedStateSTARTED captures enum value "STARTED"
	UniverseDefinitionTaskParamsImportedStateSTARTED string = "STARTED"

	// UniverseDefinitionTaskParamsImportedStateMASTERSADDED captures enum value "MASTERS_ADDED"
	UniverseDefinitionTaskParamsImportedStateMASTERSADDED string = "MASTERS_ADDED"

	// UniverseDefinitionTaskParamsImportedStateTSERVERSADDED captures enum value "TSERVERS_ADDED"
	UniverseDefinitionTaskParamsImportedStateTSERVERSADDED string = "TSERVERS_ADDED"

	// UniverseDefinitionTaskParamsImportedStateIMPORTED captures enum value "IMPORTED"
	UniverseDefinitionTaskParamsImportedStateIMPORTED string = "IMPORTED"
)

// prop value enum
func (m *UniverseDefinitionTaskParams) validateImportedStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, universeDefinitionTaskParamsTypeImportedStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UniverseDefinitionTaskParams) validateImportedState(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportedState) { // not required
		return nil
	}

	// value enum
	if err := m.validateImportedStateEnum("importedState", "body", m.ImportedState); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateNodeDetailsSet(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeDetailsSet) { // not required
		return nil
	}

	if err := validate.UniqueItems("nodeDetailsSet", "body", m.NodeDetailsSet); err != nil {
		return err
	}

	for i := 0; i < len(m.NodeDetailsSet); i++ {
		if swag.IsZero(m.NodeDetailsSet[i]) { // not required
			continue
		}

		if m.NodeDetailsSet[i] != nil {
			if err := m.NodeDetailsSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeDetailsSet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateRootCA(formats strfmt.Registry) error {
	if swag.IsZero(m.RootCA) { // not required
		return nil
	}

	if err := validate.FormatOf("rootCA", "body", "uuid", m.RootCA.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateSourceAsyncReplicationRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceAsyncReplicationRelationships) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceAsyncReplicationRelationships); i++ {
		if swag.IsZero(m.SourceAsyncReplicationRelationships[i]) { // not required
			continue
		}

		if m.SourceAsyncReplicationRelationships[i] != nil {
			if err := m.SourceAsyncReplicationRelationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateTargetAsyncReplicationRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetAsyncReplicationRelationships) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetAsyncReplicationRelationships); i++ {
		if swag.IsZero(m.TargetAsyncReplicationRelationships[i]) { // not required
			continue
		}

		if m.TargetAsyncReplicationRelationships[i] != nil {
			if err := m.TargetAsyncReplicationRelationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UniverseDefinitionTaskParams) validateUniverseUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UniverseUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("universeUUID", "body", "uuid", m.UniverseUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this universe definition task params based on the context it is used
func (m *UniverseDefinitionTaskParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommunicationPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryptionAtRestConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtraDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeDetailsSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceAsyncReplicationRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetAsyncReplicationRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UniverseDefinitionTaskParams) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UniverseDefinitionTaskParams) contextValidateCommunicationPorts(ctx context.Context, formats strfmt.Registry) error {

	if m.CommunicationPorts != nil {
		if err := m.CommunicationPorts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("communicationPorts")
			}
			return err
		}
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) contextValidateDeviceInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceInfo != nil {
		if err := m.DeviceInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceInfo")
			}
			return err
		}
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) contextValidateEncryptionAtRestConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionAtRestConfig != nil {
		if err := m.EncryptionAtRestConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionAtRestConfig")
			}
			return err
		}
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) contextValidateExtraDependencies(ctx context.Context, formats strfmt.Registry) error {

	if m.ExtraDependencies != nil {
		if err := m.ExtraDependencies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extraDependencies")
			}
			return err
		}
	}

	return nil
}

func (m *UniverseDefinitionTaskParams) contextValidateNodeDetailsSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeDetailsSet); i++ {

		if m.NodeDetailsSet[i] != nil {
			if err := m.NodeDetailsSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeDetailsSet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UniverseDefinitionTaskParams) contextValidateSourceAsyncReplicationRelationships(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceAsyncReplicationRelationships", "body", []*AsyncReplicationConfig(m.SourceAsyncReplicationRelationships)); err != nil {
		return err
	}

	for i := 0; i < len(m.SourceAsyncReplicationRelationships); i++ {

		if m.SourceAsyncReplicationRelationships[i] != nil {
			if err := m.SourceAsyncReplicationRelationships[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UniverseDefinitionTaskParams) contextValidateTargetAsyncReplicationRelationships(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "targetAsyncReplicationRelationships", "body", []*AsyncReplicationConfig(m.TargetAsyncReplicationRelationships)); err != nil {
		return err
	}

	for i := 0; i < len(m.TargetAsyncReplicationRelationships); i++ {

		if m.TargetAsyncReplicationRelationships[i] != nil {
			if err := m.TargetAsyncReplicationRelationships[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UniverseDefinitionTaskParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UniverseDefinitionTaskParams) UnmarshalBinary(b []byte) error {
	var res UniverseDefinitionTaskParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
