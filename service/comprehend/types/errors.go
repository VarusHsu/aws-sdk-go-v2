// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// The number of documents in the request exceeds the limit of 25. Try your request
// again with fewer documents.
type BatchSizeLimitExceededException struct {
	Message *string
}

func (e *BatchSizeLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchSizeLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchSizeLimitExceededException) ErrorCode() string {
	return "BatchSizeLimitExceededException"
}
func (e *BatchSizeLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Concurrent modification of the tags associated with an Amazon Comprehend
// resource is not supported.
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

// An internal server error occurred. Retry your request.
type InternalServerException struct {
	Message *string
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string             { return "InternalServerException" }
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The filter specified for the operation is invalid. Specify a different filter.
type InvalidFilterException struct {
	Message *string
}

func (e *InvalidFilterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidFilterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidFilterException) ErrorCode() string             { return "InvalidFilterException" }
func (e *InvalidFilterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request is invalid.
type InvalidRequestException struct {
	Message *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified job was not found. Check the job ID and try again.
type JobNotFoundException struct {
	Message *string
}

func (e *JobNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *JobNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *JobNotFoundException) ErrorCode() string             { return "JobNotFoundException" }
func (e *JobNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The KMS customer managed key (CMK) entered cannot be validated. Verify the key
// and re-enter it.
type KmsKeyValidationException struct {
	Message *string
}

func (e *KmsKeyValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsKeyValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsKeyValidationException) ErrorCode() string             { return "KmsKeyValidationException" }
func (e *KmsKeyValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource name is already in use. Use a different name and try your
// request again.
type ResourceInUseException struct {
	Message *string
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum number of resources per account has been exceeded. Review the
// resources, and then try your request again.
type ResourceLimitExceededException struct {
	Message *string
}

func (e *ResourceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceededException) ErrorCode() string             { return "ResourceLimitExceededException" }
func (e *ResourceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource ARN was not found. Check the ARN and try your request
// again.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource is not available. Check the resource and try your request
// again.
type ResourceUnavailableException struct {
	Message *string
}

func (e *ResourceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceUnavailableException) ErrorCode() string             { return "ResourceUnavailableException" }
func (e *ResourceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The size of the input text exceeds the limit. Use a smaller document.
type TextSizeLimitExceededException struct {
	Message *string
}

func (e *TextSizeLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TextSizeLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TextSizeLimitExceededException) ErrorCode() string             { return "TextSizeLimitExceededException" }
func (e *TextSizeLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The number of requests exceeds the limit. Resubmit your request later.
type TooManyRequestsException struct {
	Message *string
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string             { return "TooManyRequestsException" }
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request contains more tag keys than can be associated with a resource (50
// tag keys per resource).
type TooManyTagKeysException struct {
	Message *string
}

func (e *TooManyTagKeysException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagKeysException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagKeysException) ErrorCode() string             { return "TooManyTagKeysException" }
func (e *TooManyTagKeysException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request contains more tags than can be associated with a resource (50 tags
// per resource). The maximum number of tags includes both existing tags and those
// included in your current request.
type TooManyTagsException struct {
	Message *string
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Comprehend can't process the language of the input text. For all custom
// entity recognition APIs (such as CreateEntityRecognizer), only English is
// accepted. For most other APIs, such as those for Custom Classification, Amazon
// Comprehend accepts text in all supported languages. For a list of supported
// languages, see supported-languages ().
type UnsupportedLanguageException struct {
	Message *string
}

func (e *UnsupportedLanguageException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedLanguageException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedLanguageException) ErrorCode() string             { return "UnsupportedLanguageException" }
func (e *UnsupportedLanguageException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
