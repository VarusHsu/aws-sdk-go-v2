// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// An attachment with the specified ID could not be found.
type AttachmentIdNotFound struct {
	Message *string
}

func (e *AttachmentIdNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentIdNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentIdNotFound) ErrorCode() string             { return "AttachmentIdNotFound" }
func (e *AttachmentIdNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The limit for the number of attachment sets created in a short period of time
// has been exceeded.
type AttachmentLimitExceeded struct {
	Message *string
}

func (e *AttachmentLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentLimitExceeded) ErrorCode() string             { return "AttachmentLimitExceeded" }
func (e *AttachmentLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The expiration time of the attachment set has passed. The set expires 1 hour
// after it is created.
type AttachmentSetExpired struct {
	Message *string
}

func (e *AttachmentSetExpired) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentSetExpired) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentSetExpired) ErrorCode() string             { return "AttachmentSetExpired" }
func (e *AttachmentSetExpired) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An attachment set with the specified ID could not be found.
type AttachmentSetIdNotFound struct {
	Message *string
}

func (e *AttachmentSetIdNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentSetIdNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentSetIdNotFound) ErrorCode() string             { return "AttachmentSetIdNotFound" }
func (e *AttachmentSetIdNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A limit for the size of an attachment set has been exceeded. The limits are
// three attachments and 5 MB per attachment.
type AttachmentSetSizeLimitExceeded struct {
	Message *string
}

func (e *AttachmentSetSizeLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentSetSizeLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentSetSizeLimitExceeded) ErrorCode() string             { return "AttachmentSetSizeLimitExceeded" }
func (e *AttachmentSetSizeLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The case creation limit for the account has been exceeded.
type CaseCreationLimitExceeded struct {
	Message *string
}

func (e *CaseCreationLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CaseCreationLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CaseCreationLimitExceeded) ErrorCode() string             { return "CaseCreationLimitExceeded" }
func (e *CaseCreationLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested caseId could not be located.
type CaseIdNotFound struct {
	Message *string
}

func (e *CaseIdNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CaseIdNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CaseIdNotFound) ErrorCode() string             { return "CaseIdNotFound" }
func (e *CaseIdNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The limit for the number of DescribeAttachment () requests in a short period of
// time has been exceeded.
type DescribeAttachmentLimitExceeded struct {
	Message *string
}

func (e *DescribeAttachmentLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DescribeAttachmentLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DescribeAttachmentLimitExceeded) ErrorCode() string {
	return "DescribeAttachmentLimitExceeded"
}
func (e *DescribeAttachmentLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal server error occurred.
type InternalServerError struct {
	Message *string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string             { return "InternalServerError" }
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
