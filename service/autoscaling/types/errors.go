// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// The request failed because an active instance refresh for the specified Auto
// Scaling group was not found.
type ActiveInstanceRefreshNotFoundFault struct {
	Message *string
}

func (e *ActiveInstanceRefreshNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActiveInstanceRefreshNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActiveInstanceRefreshNotFoundFault) ErrorCode() string {
	return "ActiveInstanceRefreshNotFoundFault"
}
func (e *ActiveInstanceRefreshNotFoundFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// You already have an Auto Scaling group or launch configuration with this name.
type AlreadyExistsFault struct {
	Message *string
}

func (e *AlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsFault) ErrorCode() string             { return "AlreadyExistsFault" }
func (e *AlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because an active instance refresh operation already exists
// for the specified Auto Scaling group.
type InstanceRefreshInProgressFault struct {
	Message *string
}

func (e *InstanceRefreshInProgressFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InstanceRefreshInProgressFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InstanceRefreshInProgressFault) ErrorCode() string             { return "InstanceRefreshInProgressFault" }
func (e *InstanceRefreshInProgressFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The NextToken value is not valid.
type InvalidNextToken struct {
	Message *string
}

func (e *InvalidNextToken) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextToken) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextToken) ErrorCode() string             { return "InvalidNextToken" }
func (e *InvalidNextToken) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have already reached a limit for your Amazon EC2 Auto Scaling resources (for
// example, Auto Scaling groups, launch configurations, or lifecycle hooks). For
// more information, see DescribeAccountLimits
// (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeAccountLimits.html)
// in the Amazon EC2 Auto Scaling API Reference.
type LimitExceededFault struct {
	Message *string
}

func (e *LimitExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededFault) ErrorCode() string             { return "LimitExceededFault" }
func (e *LimitExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a pending update to an Amazon EC2 Auto Scaling resource (for
// example, an Auto Scaling group, instance, or load balancer).
type ResourceContentionFault struct {
	Message *string
}

func (e *ResourceContentionFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceContentionFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceContentionFault) ErrorCode() string             { return "ResourceContentionFault" }
func (e *ResourceContentionFault) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The operation can't be performed because the resource is in use.
type ResourceInUseFault struct {
	Message *string
}

func (e *ResourceInUseFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseFault) ErrorCode() string             { return "ResourceInUseFault" }
func (e *ResourceInUseFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation can't be performed because there are scaling activities in
// progress.
type ScalingActivityInProgressFault struct {
	Message *string
}

func (e *ScalingActivityInProgressFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ScalingActivityInProgressFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ScalingActivityInProgressFault) ErrorCode() string             { return "ScalingActivityInProgressFault" }
func (e *ScalingActivityInProgressFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service-linked role is not yet ready for use.
type ServiceLinkedRoleFailure struct {
	Message *string
}

func (e *ServiceLinkedRoleFailure) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLinkedRoleFailure) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLinkedRoleFailure) ErrorCode() string             { return "ServiceLinkedRoleFailure" }
func (e *ServiceLinkedRoleFailure) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
