// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// The resource hierarchy is changing.
type ConcurrentModificationException struct {
	Message *string
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Another operation is in progress on the resource that conflicts with the current
// operation.
type ConflictingOperationException struct {
	Message *string
}

func (e *ConflictingOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictingOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictingOperationException) ErrorCode() string             { return "ConflictingOperationException" }
func (e *ConflictingOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The limit has been reached on the number of custom properties for the specified
// resource.
type CustomMetadataLimitExceededException struct {
	Message *string
}

func (e *CustomMetadataLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomMetadataLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomMetadataLimitExceededException) ErrorCode() string {
	return "CustomMetadataLimitExceededException"
}
func (e *CustomMetadataLimitExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The last user in the organization is being deactivated.
type DeactivatingLastSystemUserException struct {
	Message *string

	Code *string
}

func (e *DeactivatingLastSystemUserException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeactivatingLastSystemUserException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeactivatingLastSystemUserException) ErrorCode() string {
	return "DeactivatingLastSystemUserException"
}
func (e *DeactivatingLastSystemUserException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the document is locked for comments and user tries
// to create or delete a comment on that document.
type DocumentLockedForCommentsException struct {
	Message *string
}

func (e *DocumentLockedForCommentsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DocumentLockedForCommentsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DocumentLockedForCommentsException) ErrorCode() string {
	return "DocumentLockedForCommentsException"
}
func (e *DocumentLockedForCommentsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when a valid checkout ID is not presented on document
// version upload calls for a document that has been checked out from Web client.
type DraftUploadOutOfSyncException struct {
	Message *string
}

func (e *DraftUploadOutOfSyncException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DraftUploadOutOfSyncException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DraftUploadOutOfSyncException) ErrorCode() string             { return "DraftUploadOutOfSyncException" }
func (e *DraftUploadOutOfSyncException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource already exists.
type EntityAlreadyExistsException struct {
	Message *string
}

func (e *EntityAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityAlreadyExistsException) ErrorCode() string             { return "EntityAlreadyExistsException" }
func (e *EntityAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource does not exist.
type EntityNotExistsException struct {
	Message *string

	EntityIds []*string
}

func (e *EntityNotExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityNotExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityNotExistsException) ErrorCode() string             { return "EntityNotExistsException" }
func (e *EntityNotExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The AWS Directory Service cannot reach an on-premises instance. Or a dependency
// under the control of the organization is failing, such as a connected Active
// Directory.
type FailedDependencyException struct {
	Message *string
}

func (e *FailedDependencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FailedDependencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FailedDependencyException) ErrorCode() string             { return "FailedDependencyException" }
func (e *FailedDependencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The user is undergoing transfer of ownership.
type IllegalUserStateException struct {
	Message *string
}

func (e *IllegalUserStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IllegalUserStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IllegalUserStateException) ErrorCode() string             { return "IllegalUserStateException" }
func (e *IllegalUserStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The pagination marker or limit fields are not valid.
type InvalidArgumentException struct {
	Message *string
}

func (e *InvalidArgumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgumentException) ErrorCode() string             { return "InvalidArgumentException" }
func (e *InvalidArgumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested operation is not allowed on the specified comment object.
type InvalidCommentOperationException struct {
	Message *string
}

func (e *InvalidCommentOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCommentOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCommentOperationException) ErrorCode() string {
	return "InvalidCommentOperationException"
}
func (e *InvalidCommentOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation is invalid.
type InvalidOperationException struct {
	Message *string
}

func (e *InvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOperationException) ErrorCode() string             { return "InvalidOperationException" }
func (e *InvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The password is invalid.
type InvalidPasswordException struct {
	Message *string
}

func (e *InvalidPasswordException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPasswordException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPasswordException) ErrorCode() string             { return "InvalidPasswordException" }
func (e *InvalidPasswordException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum of 100,000 folders under the parent folder has been exceeded.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified document version is not in the INITIALIZED state.
type ProhibitedStateException struct {
	Message *string
}

func (e *ProhibitedStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProhibitedStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProhibitedStateException) ErrorCode() string             { return "ProhibitedStateException" }
func (e *ProhibitedStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The response is too large to return. The request must include a filter to reduce
// the size of the response.
type RequestedEntityTooLargeException struct {
	Message *string
}

func (e *RequestedEntityTooLargeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestedEntityTooLargeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestedEntityTooLargeException) ErrorCode() string {
	return "RequestedEntityTooLargeException"
}
func (e *RequestedEntityTooLargeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource is already checked out.
type ResourceAlreadyCheckedOutException struct {
	Message *string
}

func (e *ResourceAlreadyCheckedOutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyCheckedOutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyCheckedOutException) ErrorCode() string {
	return "ResourceAlreadyCheckedOutException"
}
func (e *ResourceAlreadyCheckedOutException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// One or more of the dependencies is unavailable.
type ServiceUnavailableException struct {
	Message *string
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string             { return "ServiceUnavailableException" }
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The storage limit has been exceeded.
type StorageLimitExceededException struct {
	Message *string
}

func (e *StorageLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageLimitExceededException) ErrorCode() string             { return "StorageLimitExceededException" }
func (e *StorageLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The storage limit will be exceeded.
type StorageLimitWillExceedException struct {
	Message *string
}

func (e *StorageLimitWillExceedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageLimitWillExceedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageLimitWillExceedException) ErrorCode() string {
	return "StorageLimitWillExceedException"
}
func (e *StorageLimitWillExceedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The limit has been reached on the number of labels for the specified resource.
type TooManyLabelsException struct {
	Message *string
}

func (e *TooManyLabelsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyLabelsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyLabelsException) ErrorCode() string             { return "TooManyLabelsException" }
func (e *TooManyLabelsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of subscriptions for the WorkDocs
// instance.
type TooManySubscriptionsException struct {
	Message *string
}

func (e *TooManySubscriptionsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManySubscriptionsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManySubscriptionsException) ErrorCode() string             { return "TooManySubscriptionsException" }
func (e *TooManySubscriptionsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation is not permitted.
type UnauthorizedOperationException struct {
	Message *string

	Code *string
}

func (e *UnauthorizedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedOperationException) ErrorCode() string             { return "UnauthorizedOperationException" }
func (e *UnauthorizedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The caller does not have access to perform the action on the resource.
type UnauthorizedResourceAccessException struct {
	Message *string
}

func (e *UnauthorizedResourceAccessException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedResourceAccessException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedResourceAccessException) ErrorCode() string {
	return "UnauthorizedResourceAccessException"
}
func (e *UnauthorizedResourceAccessException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
