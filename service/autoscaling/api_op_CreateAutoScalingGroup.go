// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Auto Scaling group with the specified name and attributes. If you
// exceed your maximum limit of Auto Scaling groups, the call fails. To query this
// limit, call the DescribeAccountLimits () API. For information about updating
// this limit, see Amazon EC2 Auto Scaling Service Quotas
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-account-limits.html)
// in the Amazon EC2 Auto Scaling User Guide. For introductory exercises for
// creating an Auto Scaling group, see Getting Started with Amazon EC2 Auto Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/GettingStartedTutorial.html)
// and Tutorial: Set Up a Scaled and Load-Balanced Application
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-register-lbs-with-asg.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Auto
// Scaling Groups
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/AutoScalingGroup.html) in
// the Amazon EC2 Auto Scaling User Guide. Every Auto Scaling group has three size
// parameters (DesiredCapacity, MaxSize, and MinSize). Usually, you set these sizes
// based on a specific number of instances. However, if you configure a mixed
// instances policy that defines weights for the instance types, you must specify
// these sizes with the same units that you use for weighting instances.
func (c *Client) CreateAutoScalingGroup(ctx context.Context, params *CreateAutoScalingGroupInput, optFns ...func(*Options)) (*CreateAutoScalingGroupOutput, error) {
	stack := middleware.NewStack("CreateAutoScalingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateAutoScalingGroupMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateAutoScalingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoScalingGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateAutoScalingGroup",
			Err:           err,
		}
	}
	out := result.(*CreateAutoScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoScalingGroupInput struct {

	// The minimum size of the group.
	//
	// This member is required.
	MinSize *int32

	// The name of the placement group into which to launch your instances, if any. A
	// placement group is a logical grouping of instances within a single Availability
	// Zone. You cannot specify multiple Availability Zones and a placement group. For
	// more information, see Placement Groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in
	// the Amazon EC2 User Guide for Linux Instances.
	PlacementGroup *string

	// Parameters used to specify the launch template and version to use when an
	// instance is launched. For more information, see LaunchTemplateSpecification
	// (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LaunchTemplateSpecification.html)
	// in the Amazon EC2 Auto Scaling API Reference. You can alternatively associate a
	// launch template to the Auto Scaling group by using the MixedInstancesPolicy
	// parameter. You must specify one of the following parameters in your request:
	// LaunchConfigurationName, LaunchTemplate, InstanceId, or MixedInstancesPolicy.
	LaunchTemplate *types.LaunchTemplateSpecification

	// One or more tags. You can tag your Auto Scaling group and propagate the tags to
	// the Amazon EC2 instances it launches. Tags are not propagated to Amazon EBS
	// volumes. To add tags to Amazon EBS volumes, specify the tags in a launch
	// template but use caution. If the launch template specifies an instance tag with
	// a key that is also specified for the Auto Scaling group, Amazon EC2 Auto Scaling
	// overrides the value of that instance tag with the value specified by the Auto
	// Scaling group. For more information, see Tagging Auto Scaling Groups and
	// Instances
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-tagging.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	Tags []*types.Tag

	// One or more termination policies used to select the instance to terminate. These
	// policies are executed in the order that they are listed. For more information,
	// see Controlling Which Instances Auto Scaling Terminates During Scale In
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	TerminationPolicies []*string

	// The service to use for the health checks. The valid values are EC2 and ELB. The
	// default value is EC2. If you configure an Auto Scaling group to use ELB health
	// checks, it considers the instance unhealthy if it fails either the EC2 status
	// checks or the load balancer health checks. For more information, see Health
	// Checks for Auto Scaling Instances
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html) in the
	// Amazon EC2 Auto Scaling User Guide.
	HealthCheckType *string

	// The amount of time, in seconds, after a scaling activity completes before
	// another scaling activity can start. The default value is 300. This setting
	// applies when using simple scaling policies, but not when using other scaling
	// policies or scheduled scaling. For more information, see Scaling Cooldowns for
	// Amazon EC2 Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the
	// Amazon EC2 Auto Scaling User Guide.
	DefaultCooldown *int32

	// The desired capacity is the initial capacity of the Auto Scaling group at the
	// time of its creation and the capacity it attempts to maintain. It can scale
	// beyond this capacity if you configure automatic scaling. This number must be
	// greater than or equal to the minimum size of the group and less than or equal to
	// the maximum size of the group. If you do not specify a desired capacity, the
	// default is the minimum size of the group.
	DesiredCapacity *int32

	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before
	// checking the health status of an EC2 instance that has come into service. During
	// this time, any health check failures for the instance are ignored. The default
	// value is 0. For more information, see Health Check Grace Period
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html#health-check-grace-period)
	// in the Amazon EC2 Auto Scaling User Guide. Required if you are adding an ELB
	// health check.
	HealthCheckGracePeriod *int32

	// Indicates whether newly launched instances are protected from termination by
	// Amazon EC2 Auto Scaling when scaling in. For more information about preventing
	// instances from terminating on scale in, see Instance Protection
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection)
	// in the Amazon EC2 Auto Scaling User Guide.
	NewInstancesProtectedFromScaleIn *bool

	// The name of the Auto Scaling group. This name must be unique per Region per
	// account.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling
	// group uses to call other AWS services on your behalf. By default, Amazon EC2
	// Auto Scaling uses a service-linked role named AWSServiceRoleForAutoScaling,
	// which it creates if it does not exist. For more information, see Service-Linked
	// Roles
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	ServiceLinkedRoleARN *string

	// One or more Availability Zones for the group. This parameter is optional if you
	// specify one or more subnets for VPCZoneIdentifier. Conditional: If your account
	// supports EC2-Classic and VPC, this parameter is required to launch instances
	// into EC2-Classic.
	AvailabilityZones []*string

	// The name of the launch configuration to use when an instance is launched. To get
	// the launch configuration name, use the DescribeLaunchConfigurations () API
	// operation. New launch configurations can be created with the
	// CreateLaunchConfiguration () API. You must specify one of the following
	// parameters in your request: LaunchConfigurationName, LaunchTemplate, InstanceId,
	// or MixedInstancesPolicy.
	LaunchConfigurationName *string

	// A list of Classic Load Balancers associated with this Auto Scaling group. For
	// Application Load Balancers and Network Load Balancers, specify a list of target
	// groups using the TargetGroupARNs property instead. For more information, see
	// Using a Load Balancer with an Auto Scaling Group
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	LoadBalancerNames []*string

	// One or more lifecycle hooks.
	LifecycleHookSpecificationList []*types.LifecycleHookSpecification

	// The maximum size of the group. With a mixed instances policy that uses instance
	// weighting, Amazon EC2 Auto Scaling may need to go above MaxSize to meet your
	// capacity requirements. In this event, Amazon EC2 Auto Scaling will never go
	// above MaxSize by more than your largest instance weight (weights that define how
	// many units each instance contributes to the desired capacity of the group).
	//
	// This member is required.
	MaxSize *int32

	// The maximum amount of time, in seconds, that an instance can be in service. The
	// default is null. This parameter is optional, but if you specify a value for it,
	// you must specify a value of at least 604,800 seconds (7 days). To clear a
	// previously set value, specify a new value of 0. For more information, see
	// Replacing Auto Scaling Instances Based on Maximum Instance Lifetime
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html)
	// in the Amazon EC2 Auto Scaling User Guide. Valid Range: Minimum value of 0.
	MaxInstanceLifetime *int32

	// An embedded object that specifies a mixed instances policy. The required
	// parameters must be specified. If optional parameters are unspecified, their
	// default values are used. The policy includes parameters that not only define the
	// distribution of On-Demand Instances and Spot Instances, the maximum price to pay
	// for Spot Instances, and how the Auto Scaling group allocates instance types to
	// fulfill On-Demand and Spot capacity, but also the parameters that specify the
	// instance configuration information—the launch template and instance types.
	// <p>For more information, see <a
	// href="https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_MixedInstancesPolicy.html">MixedInstancesPolicy</a>
	// in the <i>Amazon EC2 Auto Scaling API Reference</i> and <a
	// href="https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-purchase-options.html">Auto
	// Scaling Groups with Multiple Instance Types and Purchase Options</a> in the
	// <i>Amazon EC2 Auto Scaling User Guide</i>.</p> <p>You must specify one of the
	// following parameters in your request: <code>LaunchConfigurationName</code>,
	// <code>LaunchTemplate</code>, <code>InstanceId</code>, or
	// <code>MixedInstancesPolicy</code>.</p>
	MixedInstancesPolicy *types.MixedInstancesPolicy

	// A comma-separated list of subnet IDs for your virtual private cloud (VPC). If
	// you specify VPCZoneIdentifier with AvailabilityZones, the subnets that you
	// specify for this parameter must reside in those Availability Zones. Conditional:
	// If your account supports EC2-Classic and VPC, this parameter is required to
	// launch instances into a VPC.
	VPCZoneIdentifier *string

	// The ID of the instance used to create a launch configuration for the group. To
	// get the instance ID, use the Amazon EC2 DescribeInstances
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html)
	// API operation. When you specify an ID of an instance, Amazon EC2 Auto Scaling
	// creates a new launch configuration and associates it with the group. This launch
	// configuration derives its attributes from the specified instance, except for the
	// block device mapping. You must specify one of the following parameters in your
	// request: LaunchConfigurationName, LaunchTemplate, InstanceId, or
	// MixedInstancesPolicy.
	InstanceId *string

	// The Amazon Resource Names (ARN) of the target groups to associate with the Auto
	// Scaling group. Instances are registered as targets in a target group, and
	// traffic is routed to the target group. For more information, see Using a Load
	// Balancer with an Auto Scaling Group
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	TargetGroupARNs []*string
}

type CreateAutoScalingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateAutoScalingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateAutoScalingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateAutoScalingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAutoScalingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "CreateAutoScalingGroup",
	}
}
