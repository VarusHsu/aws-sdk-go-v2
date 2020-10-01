// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// There is an error in the call or in a SQL statement.
type BadRequestException struct {
	Message *string
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There are insufficient privileges to make the call.
type ForbiddenException struct {
	Message *string
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string             { return "ForbiddenException" }
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal error occurred.
type InternalServerErrorException struct {
	Message *string
}

func (e *InternalServerErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerErrorException) ErrorCode() string             { return "InternalServerErrorException" }
func (e *InternalServerErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The resourceArn, secretArn, or transactionId value can't be found.
type NotFoundException struct {
	Message *string
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service specified by the resourceArn parameter is not available.
type ServiceUnavailableError struct {
	Message *string
}

func (e *ServiceUnavailableError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableError) ErrorCode() string             { return "ServiceUnavailableError" }
func (e *ServiceUnavailableError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The execution of the SQL statement timed out.
type StatementTimeoutException struct {
	Message *string

	DbConnectionId *int64
}

func (e *StatementTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StatementTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StatementTimeoutException) ErrorCode() string             { return "StatementTimeoutException" }
func (e *StatementTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
