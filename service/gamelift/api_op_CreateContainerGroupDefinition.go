// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is used with the Amazon GameLift containers feature, which is
// currently in public preview. Creates a ContainerGroupDefinition resource that
// describes a set of containers for hosting your game server with Amazon GameLift
// managed EC2 hosting. An Amazon GameLift container group is similar to a
// container "task" and "pod". Each container group can have one or more
// containers. Use container group definitions when you create a container fleet.
// Container group definitions determine how Amazon GameLift deploys your
// containers to each instance in a container fleet. You can create two types of
// container groups, based on scheduling strategy:
//   - A replica container group manages the containers that run your game server
//     application and supporting software. Replica container groups might be
//     replicated multiple times on each fleet instance, depending on instance
//     resources.
//   - A daemon container group manages containers that run other software, such
//     as background services, logging, or test processes. You might use a daemon
//     container group for processes that need to run only once per fleet instance, or
//     processes that need to persist independently of the replica container group.
//
// To create a container group definition, specify a group name, a list of
// container definitions, and maximum total CPU and memory requirements for the
// container group. Specify an operating system and scheduling strategy or use the
// default values. When using the Amazon Web Services CLI tool, you can pass in
// your container definitions as a JSON file. This operation requires Identity and
// Access Management (IAM) permissions to access container images in Amazon ECR
// repositories. See IAM permissions for Amazon GameLift (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-iam-policy-examples.html)
// for help setting the appropriate permissions. If successful, this operation
// creates a new ContainerGroupDefinition resource with an ARN value assigned. You
// can't change the properties of a container group definition. Instead, create a
// new one. Learn more
//   - Create a container group definition (https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-create-groups.html)
//   - Container fleet design guide (https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-design-fleet.html)
//   - Create a container definition as a JSON file (https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-definitions.html#containers-definitions-create)
func (c *Client) CreateContainerGroupDefinition(ctx context.Context, params *CreateContainerGroupDefinitionInput, optFns ...func(*Options)) (*CreateContainerGroupDefinitionOutput, error) {
	if params == nil {
		params = &CreateContainerGroupDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContainerGroupDefinition", params, optFns, c.addOperationCreateContainerGroupDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContainerGroupDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContainerGroupDefinitionInput struct {

	// Definitions for all containers in this group. Each container definition
	// identifies the container image and specifies configuration settings for the
	// container. See the Container fleet design guide (https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-design-fleet.html)
	// for container guidelines.
	//
	// This member is required.
	ContainerDefinitions []types.ContainerDefinitionInput

	// A descriptive identifier for the container group definition. The name value
	// must be unique in an Amazon Web Services Region.
	//
	// This member is required.
	Name *string

	// The platform that is used by containers in the container group definition. All
	// containers in a group must run on the same operating system.
	//
	// This member is required.
	OperatingSystem types.ContainerOperatingSystem

	// The maximum amount of CPU units to allocate to the container group. Set this
	// parameter to an integer value in CPU units (1 vCPU is equal to 1024 CPU units).
	// All containers in the group share this memory. If you specify CPU limits for
	// individual containers, set this parameter based on the following guidelines. The
	// value must be equal to or greater than the sum of the CPU limits for all
	// containers in the group.
	//
	// This member is required.
	TotalCpuLimit *int32

	// The maximum amount of memory (in MiB) to allocate to the container group. All
	// containers in the group share this memory. If you specify memory limits for
	// individual containers, set this parameter based on the following guidelines. The
	// value must be (1) greater than the sum of the soft memory limits for all
	// containers in the group, and (2) greater than any individual container's hard
	// memory limit.
	//
	// This member is required.
	TotalMemoryLimit *int32

	// The method for deploying the container group across fleet instances. A replica
	// container group might have multiple copies on each fleet instance. A daemon
	// container group has one copy per fleet instance. Default value is REPLICA .
	SchedulingStrategy types.ContainerSchedulingStrategy

	// A list of labels to assign to the container group definition resource. Tags are
	// developer-defined key-value pairs. Tagging Amazon Web Services resources are
	// useful for resource management, access management and cost allocation. For more
	// information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the Amazon Web Services General Reference.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateContainerGroupDefinitionOutput struct {

	// The properties of the newly created container group definition resource. You
	// use this resource to create a container fleet.
	ContainerGroupDefinition *types.ContainerGroupDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContainerGroupDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContainerGroupDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContainerGroupDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateContainerGroupDefinition"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateContainerGroupDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContainerGroupDefinition(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateContainerGroupDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateContainerGroupDefinition",
	}
}
