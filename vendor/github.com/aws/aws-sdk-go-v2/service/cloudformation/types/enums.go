// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountGateStatus string

// Enum values for AccountGateStatus
const (
	AccountGateStatusSucceeded AccountGateStatus = "SUCCEEDED"
	AccountGateStatusFailed    AccountGateStatus = "FAILED"
	AccountGateStatusSkipped   AccountGateStatus = "SKIPPED"
)

// Values returns all known values for AccountGateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccountGateStatus) Values() []AccountGateStatus {
	return []AccountGateStatus{
		"SUCCEEDED",
		"FAILED",
		"SKIPPED",
	}
}

type CallAs string

// Enum values for CallAs
const (
	CallAsSelf           CallAs = "SELF"
	CallAsDelegatedAdmin CallAs = "DELEGATED_ADMIN"
)

// Values returns all known values for CallAs. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (CallAs) Values() []CallAs {
	return []CallAs{
		"SELF",
		"DELEGATED_ADMIN",
	}
}

type Capability string

// Enum values for Capability
const (
	CapabilityCapabilityIam        Capability = "CAPABILITY_IAM"
	CapabilityCapabilityNamedIam   Capability = "CAPABILITY_NAMED_IAM"
	CapabilityCapabilityAutoExpand Capability = "CAPABILITY_AUTO_EXPAND"
)

// Values returns all known values for Capability. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Capability) Values() []Capability {
	return []Capability{
		"CAPABILITY_IAM",
		"CAPABILITY_NAMED_IAM",
		"CAPABILITY_AUTO_EXPAND",
	}
}

type Category string

// Enum values for Category
const (
	CategoryRegistered Category = "REGISTERED"
	CategoryActivated  Category = "ACTIVATED"
	CategoryThirdParty Category = "THIRD_PARTY"
	CategoryAwsTypes   Category = "AWS_TYPES"
)

// Values returns all known values for Category. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Category) Values() []Category {
	return []Category{
		"REGISTERED",
		"ACTIVATED",
		"THIRD_PARTY",
		"AWS_TYPES",
	}
}

type ChangeAction string

// Enum values for ChangeAction
const (
	ChangeActionAdd     ChangeAction = "Add"
	ChangeActionModify  ChangeAction = "Modify"
	ChangeActionRemove  ChangeAction = "Remove"
	ChangeActionImport  ChangeAction = "Import"
	ChangeActionDynamic ChangeAction = "Dynamic"
)

// Values returns all known values for ChangeAction. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeAction) Values() []ChangeAction {
	return []ChangeAction{
		"Add",
		"Modify",
		"Remove",
		"Import",
		"Dynamic",
	}
}

type ChangeSetStatus string

// Enum values for ChangeSetStatus
const (
	ChangeSetStatusCreatePending    ChangeSetStatus = "CREATE_PENDING"
	ChangeSetStatusCreateInProgress ChangeSetStatus = "CREATE_IN_PROGRESS"
	ChangeSetStatusCreateComplete   ChangeSetStatus = "CREATE_COMPLETE"
	ChangeSetStatusDeletePending    ChangeSetStatus = "DELETE_PENDING"
	ChangeSetStatusDeleteInProgress ChangeSetStatus = "DELETE_IN_PROGRESS"
	ChangeSetStatusDeleteComplete   ChangeSetStatus = "DELETE_COMPLETE"
	ChangeSetStatusDeleteFailed     ChangeSetStatus = "DELETE_FAILED"
	ChangeSetStatusFailed           ChangeSetStatus = "FAILED"
)

// Values returns all known values for ChangeSetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChangeSetStatus) Values() []ChangeSetStatus {
	return []ChangeSetStatus{
		"CREATE_PENDING",
		"CREATE_IN_PROGRESS",
		"CREATE_COMPLETE",
		"DELETE_PENDING",
		"DELETE_IN_PROGRESS",
		"DELETE_COMPLETE",
		"DELETE_FAILED",
		"FAILED",
	}
}

type ChangeSetType string

// Enum values for ChangeSetType
const (
	ChangeSetTypeCreate ChangeSetType = "CREATE"
	ChangeSetTypeUpdate ChangeSetType = "UPDATE"
	ChangeSetTypeImport ChangeSetType = "IMPORT"
)

// Values returns all known values for ChangeSetType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChangeSetType) Values() []ChangeSetType {
	return []ChangeSetType{
		"CREATE",
		"UPDATE",
		"IMPORT",
	}
}

type ChangeSource string

// Enum values for ChangeSource
const (
	ChangeSourceResourceReference  ChangeSource = "ResourceReference"
	ChangeSourceParameterReference ChangeSource = "ParameterReference"
	ChangeSourceResourceAttribute  ChangeSource = "ResourceAttribute"
	ChangeSourceDirectModification ChangeSource = "DirectModification"
	ChangeSourceAutomatic          ChangeSource = "Automatic"
)

// Values returns all known values for ChangeSource. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeSource) Values() []ChangeSource {
	return []ChangeSource{
		"ResourceReference",
		"ParameterReference",
		"ResourceAttribute",
		"DirectModification",
		"Automatic",
	}
}

type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypeResource ChangeType = "Resource"
)

// Values returns all known values for ChangeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeType) Values() []ChangeType {
	return []ChangeType{
		"Resource",
	}
}

type DeprecatedStatus string

// Enum values for DeprecatedStatus
const (
	DeprecatedStatusLive       DeprecatedStatus = "LIVE"
	DeprecatedStatusDeprecated DeprecatedStatus = "DEPRECATED"
)

// Values returns all known values for DeprecatedStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeprecatedStatus) Values() []DeprecatedStatus {
	return []DeprecatedStatus{
		"LIVE",
		"DEPRECATED",
	}
}

type DifferenceType string

// Enum values for DifferenceType
const (
	DifferenceTypeAdd      DifferenceType = "ADD"
	DifferenceTypeRemove   DifferenceType = "REMOVE"
	DifferenceTypeNotEqual DifferenceType = "NOT_EQUAL"
)

// Values returns all known values for DifferenceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DifferenceType) Values() []DifferenceType {
	return []DifferenceType{
		"ADD",
		"REMOVE",
		"NOT_EQUAL",
	}
}

type EvaluationType string

// Enum values for EvaluationType
const (
	EvaluationTypeStatic  EvaluationType = "Static"
	EvaluationTypeDynamic EvaluationType = "Dynamic"
)

// Values returns all known values for EvaluationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EvaluationType) Values() []EvaluationType {
	return []EvaluationType{
		"Static",
		"Dynamic",
	}
}

type ExecutionStatus string

// Enum values for ExecutionStatus
const (
	ExecutionStatusUnavailable       ExecutionStatus = "UNAVAILABLE"
	ExecutionStatusAvailable         ExecutionStatus = "AVAILABLE"
	ExecutionStatusExecuteInProgress ExecutionStatus = "EXECUTE_IN_PROGRESS"
	ExecutionStatusExecuteComplete   ExecutionStatus = "EXECUTE_COMPLETE"
	ExecutionStatusExecuteFailed     ExecutionStatus = "EXECUTE_FAILED"
	ExecutionStatusObsolete          ExecutionStatus = "OBSOLETE"
)

// Values returns all known values for ExecutionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExecutionStatus) Values() []ExecutionStatus {
	return []ExecutionStatus{
		"UNAVAILABLE",
		"AVAILABLE",
		"EXECUTE_IN_PROGRESS",
		"EXECUTE_COMPLETE",
		"EXECUTE_FAILED",
		"OBSOLETE",
	}
}

type HandlerErrorCode string

// Enum values for HandlerErrorCode
const (
	HandlerErrorCodeNotUpdatable             HandlerErrorCode = "NotUpdatable"
	HandlerErrorCodeInvalidRequest           HandlerErrorCode = "InvalidRequest"
	HandlerErrorCodeAccessDenied             HandlerErrorCode = "AccessDenied"
	HandlerErrorCodeInvalidCredentials       HandlerErrorCode = "InvalidCredentials"
	HandlerErrorCodeAlreadyExists            HandlerErrorCode = "AlreadyExists"
	HandlerErrorCodeNotFound                 HandlerErrorCode = "NotFound"
	HandlerErrorCodeResourceConflict         HandlerErrorCode = "ResourceConflict"
	HandlerErrorCodeThrottling               HandlerErrorCode = "Throttling"
	HandlerErrorCodeServiceLimitExceeded     HandlerErrorCode = "ServiceLimitExceeded"
	HandlerErrorCodeServiceTimeout           HandlerErrorCode = "NotStabilized"
	HandlerErrorCodeGeneralServiceException  HandlerErrorCode = "GeneralServiceException"
	HandlerErrorCodeServiceInternalError     HandlerErrorCode = "ServiceInternalError"
	HandlerErrorCodeNetworkFailure           HandlerErrorCode = "NetworkFailure"
	HandlerErrorCodeInternalFailure          HandlerErrorCode = "InternalFailure"
	HandlerErrorCodeInvalidTypeConfiguration HandlerErrorCode = "InvalidTypeConfiguration"
)

// Values returns all known values for HandlerErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HandlerErrorCode) Values() []HandlerErrorCode {
	return []HandlerErrorCode{
		"NotUpdatable",
		"InvalidRequest",
		"AccessDenied",
		"InvalidCredentials",
		"AlreadyExists",
		"NotFound",
		"ResourceConflict",
		"Throttling",
		"ServiceLimitExceeded",
		"NotStabilized",
		"GeneralServiceException",
		"ServiceInternalError",
		"NetworkFailure",
		"InternalFailure",
		"InvalidTypeConfiguration",
	}
}

type IdentityProvider string

// Enum values for IdentityProvider
const (
	IdentityProviderAwsMarketplace IdentityProvider = "AWS_Marketplace"
	IdentityProviderGitHub         IdentityProvider = "GitHub"
	IdentityProviderBitbucket      IdentityProvider = "Bitbucket"
)

// Values returns all known values for IdentityProvider. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IdentityProvider) Values() []IdentityProvider {
	return []IdentityProvider{
		"AWS_Marketplace",
		"GitHub",
		"Bitbucket",
	}
}

type OnFailure string

// Enum values for OnFailure
const (
	OnFailureDoNothing OnFailure = "DO_NOTHING"
	OnFailureRollback  OnFailure = "ROLLBACK"
	OnFailureDelete    OnFailure = "DELETE"
)

// Values returns all known values for OnFailure. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OnFailure) Values() []OnFailure {
	return []OnFailure{
		"DO_NOTHING",
		"ROLLBACK",
		"DELETE",
	}
}

type Op string

// Enum values for Op
const (
	OpAdd     Op = "add"
	OpRemove  Op = "remove"
	OpReplace Op = "replace"
	OpTest    Op = "test"
)

// Values returns all known values for Op. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Op) Values() []Op {
	return []Op{
		"add",
		"remove",
		"replace",
		"test",
	}
}

type Operation string

// Enum values for Operation
const (
	OperationCreate Operation = "CREATE"
	OperationUpdate Operation = "UPDATE"
	OperationDelete Operation = "DELETE"
)

// Values returns all known values for Operation. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Operation) Values() []Operation {
	return []Operation{
		"CREATE",
		"UPDATE",
		"DELETE",
	}
}

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusPending          OperationStatus = "PENDING"
	OperationStatusInProgress       OperationStatus = "IN_PROGRESS"
	OperationStatusSuccess          OperationStatus = "SUCCESS"
	OperationStatusFailed           OperationStatus = "FAILED"
	OperationStatusCancelInProgress OperationStatus = "CANCEL_IN_PROGRESS"
	OperationStatusCancelComplete   OperationStatus = "CANCEL_COMPLETE"
)

// Values returns all known values for OperationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationStatus) Values() []OperationStatus {
	return []OperationStatus{
		"PENDING",
		"IN_PROGRESS",
		"SUCCESS",
		"FAILED",
		"CANCEL_IN_PROGRESS",
		"CANCEL_COMPLETE",
	}
}

type PermissionModels string

// Enum values for PermissionModels
const (
	PermissionModelsServiceManaged PermissionModels = "SERVICE_MANAGED"
	PermissionModelsSelfManaged    PermissionModels = "SELF_MANAGED"
)

// Values returns all known values for PermissionModels. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PermissionModels) Values() []PermissionModels {
	return []PermissionModels{
		"SERVICE_MANAGED",
		"SELF_MANAGED",
	}
}

type ProvisioningType string

// Enum values for ProvisioningType
const (
	ProvisioningTypeNonProvisionable ProvisioningType = "NON_PROVISIONABLE"
	ProvisioningTypeImmutable        ProvisioningType = "IMMUTABLE"
	ProvisioningTypeFullyMutable     ProvisioningType = "FULLY_MUTABLE"
)

// Values returns all known values for ProvisioningType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProvisioningType) Values() []ProvisioningType {
	return []ProvisioningType{
		"NON_PROVISIONABLE",
		"IMMUTABLE",
		"FULLY_MUTABLE",
	}
}

type PublisherStatus string

// Enum values for PublisherStatus
const (
	PublisherStatusVerified   PublisherStatus = "VERIFIED"
	PublisherStatusUnverified PublisherStatus = "UNVERIFIED"
)

// Values returns all known values for PublisherStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PublisherStatus) Values() []PublisherStatus {
	return []PublisherStatus{
		"VERIFIED",
		"UNVERIFIED",
	}
}

type RegionConcurrencyType string

// Enum values for RegionConcurrencyType
const (
	RegionConcurrencyTypeSequential RegionConcurrencyType = "SEQUENTIAL"
	RegionConcurrencyTypeParallel   RegionConcurrencyType = "PARALLEL"
)

// Values returns all known values for RegionConcurrencyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RegionConcurrencyType) Values() []RegionConcurrencyType {
	return []RegionConcurrencyType{
		"SEQUENTIAL",
		"PARALLEL",
	}
}

type RegistrationStatus string

// Enum values for RegistrationStatus
const (
	RegistrationStatusComplete   RegistrationStatus = "COMPLETE"
	RegistrationStatusInProgress RegistrationStatus = "IN_PROGRESS"
	RegistrationStatusFailed     RegistrationStatus = "FAILED"
)

// Values returns all known values for RegistrationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RegistrationStatus) Values() []RegistrationStatus {
	return []RegistrationStatus{
		"COMPLETE",
		"IN_PROGRESS",
		"FAILED",
	}
}

type RegistryType string

// Enum values for RegistryType
const (
	RegistryTypeResource RegistryType = "RESOURCE"
	RegistryTypeModule   RegistryType = "MODULE"
)

// Values returns all known values for RegistryType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RegistryType) Values() []RegistryType {
	return []RegistryType{
		"RESOURCE",
		"MODULE",
	}
}

type Replacement string

// Enum values for Replacement
const (
	ReplacementTrue        Replacement = "True"
	ReplacementFalse       Replacement = "False"
	ReplacementConditional Replacement = "Conditional"
)

// Values returns all known values for Replacement. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Replacement) Values() []Replacement {
	return []Replacement{
		"True",
		"False",
		"Conditional",
	}
}

type RequiresRecreation string

// Enum values for RequiresRecreation
const (
	RequiresRecreationNever         RequiresRecreation = "Never"
	RequiresRecreationConditionally RequiresRecreation = "Conditionally"
	RequiresRecreationAlways        RequiresRecreation = "Always"
)

// Values returns all known values for RequiresRecreation. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RequiresRecreation) Values() []RequiresRecreation {
	return []RequiresRecreation{
		"Never",
		"Conditionally",
		"Always",
	}
}

type ResourceAttribute string

// Enum values for ResourceAttribute
const (
	ResourceAttributeProperties     ResourceAttribute = "Properties"
	ResourceAttributeMetadata       ResourceAttribute = "Metadata"
	ResourceAttributeCreationPolicy ResourceAttribute = "CreationPolicy"
	ResourceAttributeUpdatePolicy   ResourceAttribute = "UpdatePolicy"
	ResourceAttributeDeletionPolicy ResourceAttribute = "DeletionPolicy"
	ResourceAttributeTags           ResourceAttribute = "Tags"
)

// Values returns all known values for ResourceAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceAttribute) Values() []ResourceAttribute {
	return []ResourceAttribute{
		"Properties",
		"Metadata",
		"CreationPolicy",
		"UpdatePolicy",
		"DeletionPolicy",
		"Tags",
	}
}

type ResourceSignalStatus string

// Enum values for ResourceSignalStatus
const (
	ResourceSignalStatusSuccess ResourceSignalStatus = "SUCCESS"
	ResourceSignalStatusFailure ResourceSignalStatus = "FAILURE"
)

// Values returns all known values for ResourceSignalStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceSignalStatus) Values() []ResourceSignalStatus {
	return []ResourceSignalStatus{
		"SUCCESS",
		"FAILURE",
	}
}

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusCreateInProgress         ResourceStatus = "CREATE_IN_PROGRESS"
	ResourceStatusCreateFailed             ResourceStatus = "CREATE_FAILED"
	ResourceStatusCreateComplete           ResourceStatus = "CREATE_COMPLETE"
	ResourceStatusDeleteInProgress         ResourceStatus = "DELETE_IN_PROGRESS"
	ResourceStatusDeleteFailed             ResourceStatus = "DELETE_FAILED"
	ResourceStatusDeleteComplete           ResourceStatus = "DELETE_COMPLETE"
	ResourceStatusDeleteSkipped            ResourceStatus = "DELETE_SKIPPED"
	ResourceStatusUpdateInProgress         ResourceStatus = "UPDATE_IN_PROGRESS"
	ResourceStatusUpdateFailed             ResourceStatus = "UPDATE_FAILED"
	ResourceStatusUpdateComplete           ResourceStatus = "UPDATE_COMPLETE"
	ResourceStatusImportFailed             ResourceStatus = "IMPORT_FAILED"
	ResourceStatusImportComplete           ResourceStatus = "IMPORT_COMPLETE"
	ResourceStatusImportInProgress         ResourceStatus = "IMPORT_IN_PROGRESS"
	ResourceStatusImportRollbackInProgress ResourceStatus = "IMPORT_ROLLBACK_IN_PROGRESS"
	ResourceStatusImportRollbackFailed     ResourceStatus = "IMPORT_ROLLBACK_FAILED"
	ResourceStatusImportRollbackComplete   ResourceStatus = "IMPORT_ROLLBACK_COMPLETE"
)

// Values returns all known values for ResourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceStatus) Values() []ResourceStatus {
	return []ResourceStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"CREATE_COMPLETE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETE_COMPLETE",
		"DELETE_SKIPPED",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED",
		"UPDATE_COMPLETE",
		"IMPORT_FAILED",
		"IMPORT_COMPLETE",
		"IMPORT_IN_PROGRESS",
		"IMPORT_ROLLBACK_IN_PROGRESS",
		"IMPORT_ROLLBACK_FAILED",
		"IMPORT_ROLLBACK_COMPLETE",
	}
}

type StackDriftDetectionStatus string

// Enum values for StackDriftDetectionStatus
const (
	StackDriftDetectionStatusDetectionInProgress StackDriftDetectionStatus = "DETECTION_IN_PROGRESS"
	StackDriftDetectionStatusDetectionFailed     StackDriftDetectionStatus = "DETECTION_FAILED"
	StackDriftDetectionStatusDetectionComplete   StackDriftDetectionStatus = "DETECTION_COMPLETE"
)

// Values returns all known values for StackDriftDetectionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (StackDriftDetectionStatus) Values() []StackDriftDetectionStatus {
	return []StackDriftDetectionStatus{
		"DETECTION_IN_PROGRESS",
		"DETECTION_FAILED",
		"DETECTION_COMPLETE",
	}
}

type StackDriftStatus string

// Enum values for StackDriftStatus
const (
	StackDriftStatusDrifted    StackDriftStatus = "DRIFTED"
	StackDriftStatusInSync     StackDriftStatus = "IN_SYNC"
	StackDriftStatusUnknown    StackDriftStatus = "UNKNOWN"
	StackDriftStatusNotChecked StackDriftStatus = "NOT_CHECKED"
)

// Values returns all known values for StackDriftStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackDriftStatus) Values() []StackDriftStatus {
	return []StackDriftStatus{
		"DRIFTED",
		"IN_SYNC",
		"UNKNOWN",
		"NOT_CHECKED",
	}
}

type StackInstanceDetailedStatus string

// Enum values for StackInstanceDetailedStatus
const (
	StackInstanceDetailedStatusPending    StackInstanceDetailedStatus = "PENDING"
	StackInstanceDetailedStatusRunning    StackInstanceDetailedStatus = "RUNNING"
	StackInstanceDetailedStatusSucceeded  StackInstanceDetailedStatus = "SUCCEEDED"
	StackInstanceDetailedStatusFailed     StackInstanceDetailedStatus = "FAILED"
	StackInstanceDetailedStatusCancelled  StackInstanceDetailedStatus = "CANCELLED"
	StackInstanceDetailedStatusInoperable StackInstanceDetailedStatus = "INOPERABLE"
)

// Values returns all known values for StackInstanceDetailedStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (StackInstanceDetailedStatus) Values() []StackInstanceDetailedStatus {
	return []StackInstanceDetailedStatus{
		"PENDING",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
		"INOPERABLE",
	}
}

type StackInstanceFilterName string

// Enum values for StackInstanceFilterName
const (
	StackInstanceFilterNameDetailedStatus StackInstanceFilterName = "DETAILED_STATUS"
)

// Values returns all known values for StackInstanceFilterName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackInstanceFilterName) Values() []StackInstanceFilterName {
	return []StackInstanceFilterName{
		"DETAILED_STATUS",
	}
}

type StackInstanceStatus string

// Enum values for StackInstanceStatus
const (
	StackInstanceStatusCurrent    StackInstanceStatus = "CURRENT"
	StackInstanceStatusOutdated   StackInstanceStatus = "OUTDATED"
	StackInstanceStatusInoperable StackInstanceStatus = "INOPERABLE"
)

// Values returns all known values for StackInstanceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackInstanceStatus) Values() []StackInstanceStatus {
	return []StackInstanceStatus{
		"CURRENT",
		"OUTDATED",
		"INOPERABLE",
	}
}

type StackResourceDriftStatus string

// Enum values for StackResourceDriftStatus
const (
	StackResourceDriftStatusInSync     StackResourceDriftStatus = "IN_SYNC"
	StackResourceDriftStatusModified   StackResourceDriftStatus = "MODIFIED"
	StackResourceDriftStatusDeleted    StackResourceDriftStatus = "DELETED"
	StackResourceDriftStatusNotChecked StackResourceDriftStatus = "NOT_CHECKED"
)

// Values returns all known values for StackResourceDriftStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackResourceDriftStatus) Values() []StackResourceDriftStatus {
	return []StackResourceDriftStatus{
		"IN_SYNC",
		"MODIFIED",
		"DELETED",
		"NOT_CHECKED",
	}
}

type StackSetDriftDetectionStatus string

// Enum values for StackSetDriftDetectionStatus
const (
	StackSetDriftDetectionStatusCompleted      StackSetDriftDetectionStatus = "COMPLETED"
	StackSetDriftDetectionStatusFailed         StackSetDriftDetectionStatus = "FAILED"
	StackSetDriftDetectionStatusPartialSuccess StackSetDriftDetectionStatus = "PARTIAL_SUCCESS"
	StackSetDriftDetectionStatusInProgress     StackSetDriftDetectionStatus = "IN_PROGRESS"
	StackSetDriftDetectionStatusStopped        StackSetDriftDetectionStatus = "STOPPED"
)

// Values returns all known values for StackSetDriftDetectionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (StackSetDriftDetectionStatus) Values() []StackSetDriftDetectionStatus {
	return []StackSetDriftDetectionStatus{
		"COMPLETED",
		"FAILED",
		"PARTIAL_SUCCESS",
		"IN_PROGRESS",
		"STOPPED",
	}
}

type StackSetDriftStatus string

// Enum values for StackSetDriftStatus
const (
	StackSetDriftStatusDrifted    StackSetDriftStatus = "DRIFTED"
	StackSetDriftStatusInSync     StackSetDriftStatus = "IN_SYNC"
	StackSetDriftStatusNotChecked StackSetDriftStatus = "NOT_CHECKED"
)

// Values returns all known values for StackSetDriftStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackSetDriftStatus) Values() []StackSetDriftStatus {
	return []StackSetDriftStatus{
		"DRIFTED",
		"IN_SYNC",
		"NOT_CHECKED",
	}
}

type StackSetOperationAction string

// Enum values for StackSetOperationAction
const (
	StackSetOperationActionCreate      StackSetOperationAction = "CREATE"
	StackSetOperationActionUpdate      StackSetOperationAction = "UPDATE"
	StackSetOperationActionDelete      StackSetOperationAction = "DELETE"
	StackSetOperationActionDetectDrift StackSetOperationAction = "DETECT_DRIFT"
)

// Values returns all known values for StackSetOperationAction. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackSetOperationAction) Values() []StackSetOperationAction {
	return []StackSetOperationAction{
		"CREATE",
		"UPDATE",
		"DELETE",
		"DETECT_DRIFT",
	}
}

type StackSetOperationResultStatus string

// Enum values for StackSetOperationResultStatus
const (
	StackSetOperationResultStatusPending   StackSetOperationResultStatus = "PENDING"
	StackSetOperationResultStatusRunning   StackSetOperationResultStatus = "RUNNING"
	StackSetOperationResultStatusSucceeded StackSetOperationResultStatus = "SUCCEEDED"
	StackSetOperationResultStatusFailed    StackSetOperationResultStatus = "FAILED"
	StackSetOperationResultStatusCancelled StackSetOperationResultStatus = "CANCELLED"
)

// Values returns all known values for StackSetOperationResultStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (StackSetOperationResultStatus) Values() []StackSetOperationResultStatus {
	return []StackSetOperationResultStatus{
		"PENDING",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
	}
}

type StackSetOperationStatus string

// Enum values for StackSetOperationStatus
const (
	StackSetOperationStatusRunning   StackSetOperationStatus = "RUNNING"
	StackSetOperationStatusSucceeded StackSetOperationStatus = "SUCCEEDED"
	StackSetOperationStatusFailed    StackSetOperationStatus = "FAILED"
	StackSetOperationStatusStopping  StackSetOperationStatus = "STOPPING"
	StackSetOperationStatusStopped   StackSetOperationStatus = "STOPPED"
	StackSetOperationStatusQueued    StackSetOperationStatus = "QUEUED"
)

// Values returns all known values for StackSetOperationStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackSetOperationStatus) Values() []StackSetOperationStatus {
	return []StackSetOperationStatus{
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"STOPPING",
		"STOPPED",
		"QUEUED",
	}
}

type StackSetStatus string

// Enum values for StackSetStatus
const (
	StackSetStatusActive  StackSetStatus = "ACTIVE"
	StackSetStatusDeleted StackSetStatus = "DELETED"
)

// Values returns all known values for StackSetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackSetStatus) Values() []StackSetStatus {
	return []StackSetStatus{
		"ACTIVE",
		"DELETED",
	}
}

type StackStatus string

// Enum values for StackStatus
const (
	StackStatusCreateInProgress                        StackStatus = "CREATE_IN_PROGRESS"
	StackStatusCreateFailed                            StackStatus = "CREATE_FAILED"
	StackStatusCreateComplete                          StackStatus = "CREATE_COMPLETE"
	StackStatusRollbackInProgress                      StackStatus = "ROLLBACK_IN_PROGRESS"
	StackStatusRollbackFailed                          StackStatus = "ROLLBACK_FAILED"
	StackStatusRollbackComplete                        StackStatus = "ROLLBACK_COMPLETE"
	StackStatusDeleteInProgress                        StackStatus = "DELETE_IN_PROGRESS"
	StackStatusDeleteFailed                            StackStatus = "DELETE_FAILED"
	StackStatusDeleteComplete                          StackStatus = "DELETE_COMPLETE"
	StackStatusUpdateInProgress                        StackStatus = "UPDATE_IN_PROGRESS"
	StackStatusUpdateCompleteCleanupInProgress         StackStatus = "UPDATE_COMPLETE_CLEANUP_IN_PROGRESS"
	StackStatusUpdateComplete                          StackStatus = "UPDATE_COMPLETE"
	StackStatusUpdateRollbackInProgress                StackStatus = "UPDATE_ROLLBACK_IN_PROGRESS"
	StackStatusUpdateRollbackFailed                    StackStatus = "UPDATE_ROLLBACK_FAILED"
	StackStatusUpdateRollbackCompleteCleanupInProgress StackStatus = "UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS"
	StackStatusUpdateRollbackComplete                  StackStatus = "UPDATE_ROLLBACK_COMPLETE"
	StackStatusReviewInProgress                        StackStatus = "REVIEW_IN_PROGRESS"
	StackStatusImportInProgress                        StackStatus = "IMPORT_IN_PROGRESS"
	StackStatusImportComplete                          StackStatus = "IMPORT_COMPLETE"
	StackStatusImportRollbackInProgress                StackStatus = "IMPORT_ROLLBACK_IN_PROGRESS"
	StackStatusImportRollbackFailed                    StackStatus = "IMPORT_ROLLBACK_FAILED"
	StackStatusImportRollbackComplete                  StackStatus = "IMPORT_ROLLBACK_COMPLETE"
)

// Values returns all known values for StackStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StackStatus) Values() []StackStatus {
	return []StackStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"CREATE_COMPLETE",
		"ROLLBACK_IN_PROGRESS",
		"ROLLBACK_FAILED",
		"ROLLBACK_COMPLETE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETE_COMPLETE",
		"UPDATE_IN_PROGRESS",
		"UPDATE_COMPLETE_CLEANUP_IN_PROGRESS",
		"UPDATE_COMPLETE",
		"UPDATE_ROLLBACK_IN_PROGRESS",
		"UPDATE_ROLLBACK_FAILED",
		"UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS",
		"UPDATE_ROLLBACK_COMPLETE",
		"REVIEW_IN_PROGRESS",
		"IMPORT_IN_PROGRESS",
		"IMPORT_COMPLETE",
		"IMPORT_ROLLBACK_IN_PROGRESS",
		"IMPORT_ROLLBACK_FAILED",
		"IMPORT_ROLLBACK_COMPLETE",
	}
}

type TemplateStage string

// Enum values for TemplateStage
const (
	TemplateStageOriginal  TemplateStage = "Original"
	TemplateStageProcessed TemplateStage = "Processed"
)

// Values returns all known values for TemplateStage. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TemplateStage) Values() []TemplateStage {
	return []TemplateStage{
		"Original",
		"Processed",
	}
}

type ThirdPartyType string

// Enum values for ThirdPartyType
const (
	ThirdPartyTypeResource ThirdPartyType = "RESOURCE"
	ThirdPartyTypeModule   ThirdPartyType = "MODULE"
)

// Values returns all known values for ThirdPartyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThirdPartyType) Values() []ThirdPartyType {
	return []ThirdPartyType{
		"RESOURCE",
		"MODULE",
	}
}

type TypeTestsStatus string

// Enum values for TypeTestsStatus
const (
	TypeTestsStatusPassed     TypeTestsStatus = "PASSED"
	TypeTestsStatusFailed     TypeTestsStatus = "FAILED"
	TypeTestsStatusInProgress TypeTestsStatus = "IN_PROGRESS"
	TypeTestsStatusNotTested  TypeTestsStatus = "NOT_TESTED"
)

// Values returns all known values for TypeTestsStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TypeTestsStatus) Values() []TypeTestsStatus {
	return []TypeTestsStatus{
		"PASSED",
		"FAILED",
		"IN_PROGRESS",
		"NOT_TESTED",
	}
}

type VersionBump string

// Enum values for VersionBump
const (
	VersionBumpMajor VersionBump = "MAJOR"
	VersionBumpMinor VersionBump = "MINOR"
)

// Values returns all known values for VersionBump. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (VersionBump) Values() []VersionBump {
	return []VersionBump{
		"MAJOR",
		"MINOR",
	}
}

type Visibility string

// Enum values for Visibility
const (
	VisibilityPublic  Visibility = "PUBLIC"
	VisibilityPrivate Visibility = "PRIVATE"
)

// Values returns all known values for Visibility. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Visibility) Values() []Visibility {
	return []Visibility{
		"PUBLIC",
		"PRIVATE",
	}
}
