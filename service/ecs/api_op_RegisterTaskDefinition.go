// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a new task definition from the supplied family and
// containerDefinitions . Optionally, you can add data volumes to your containers
// with the volumes parameter. For more information about task definition
// parameters and defaults, see [Amazon ECS Task Definitions]in the Amazon Elastic Container Service Developer
// Guide.
//
// You can specify a role for your task with the taskRoleArn parameter. When you
// specify a role for a task, its containers can then use the latest versions of
// the CLI or SDKs to make API requests to the Amazon Web Services services that
// are specified in the policy that's associated with the role. For more
// information, see [IAM Roles for Tasks]in the Amazon Elastic Container Service Developer Guide.
//
// You can specify a Docker networking mode for the containers in your task
// definition with the networkMode parameter. The available network modes
// correspond to those described in [Network settings]in the Docker run reference. If you specify
// the awsvpc network mode, the task is allocated an elastic network interface,
// and you must specify a NetworkConfigurationwhen you create a service or run a task with the task
// definition. For more information, see [Task Networking]in the Amazon Elastic Container Service
// Developer Guide.
//
// [Amazon ECS Task Definitions]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html
// [Task Networking]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html
// [IAM Roles for Tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
// [Network settings]: https://docs.docker.com/engine/reference/run/#/network-settings
func (c *Client) RegisterTaskDefinition(ctx context.Context, params *RegisterTaskDefinitionInput, optFns ...func(*Options)) (*RegisterTaskDefinitionOutput, error) {
	if params == nil {
		params = &RegisterTaskDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterTaskDefinition", params, optFns, c.addOperationRegisterTaskDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTaskDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTaskDefinitionInput struct {

	// A list of container definitions in JSON format that describe the different
	// containers that make up your task.
	//
	// This member is required.
	ContainerDefinitions []types.ContainerDefinition

	// You must specify a family for a task definition. You can use it track multiple
	// versions of the same task definition. The family is used as a name for your
	// task definition. Up to 255 letters (uppercase and lowercase), numbers,
	// underscores, and hyphens are allowed.
	//
	// This member is required.
	Family *string

	// The number of CPU units used by the task. It can be expressed as an integer
	// using CPU units (for example, 1024 ) or as a string using vCPUs (for example, 1
	// vCPU or 1 vcpu ) in a task definition. String values are converted to an integer
	// indicating the CPU units when the task definition is registered.
	//
	// Task-level CPU and memory parameters are ignored for Windows containers. We
	// recommend specifying container-level resources for Windows containers.
	//
	// If you're using the EC2 launch type, this field is optional. Supported values
	// are between 128 CPU units ( 0.125 vCPUs) and 10240 CPU units ( 10 vCPUs). If
	// you do not specify a value, the parameter is ignored.
	//
	// If you're using the Fargate launch type, this field is required and you must
	// use one of the following values, which determines your range of supported values
	// for the memory parameter:
	//
	// The CPU units cannot be less than 1 vCPU when you use Windows containers on
	// Fargate.
	//
	//   - 256 (.25 vCPU) - Available memory values: 512 (0.5 GB), 1024 (1 GB), 2048 (2
	//   GB)
	//
	//   - 512 (.5 vCPU) - Available memory values: 1024 (1 GB), 2048 (2 GB), 3072 (3
	//   GB), 4096 (4 GB)
	//
	//   - 1024 (1 vCPU) - Available memory values: 2048 (2 GB), 3072 (3 GB), 4096 (4
	//   GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	//
	//   - 2048 (2 vCPU) - Available memory values: 4096 (4 GB) and 16384 (16 GB) in
	//   increments of 1024 (1 GB)
	//
	//   - 4096 (4 vCPU) - Available memory values: 8192 (8 GB) and 30720 (30 GB) in
	//   increments of 1024 (1 GB)
	//
	//   - 8192 (8 vCPU) - Available memory values: 16 GB and 60 GB in 4 GB increments
	//
	// This option requires Linux platform 1.4.0 or later.
	//
	//   - 16384 (16vCPU) - Available memory values: 32GB and 120 GB in 8 GB increments
	//
	// This option requires Linux platform 1.4.0 or later.
	Cpu *string

	// The amount of ephemeral storage to allocate for the task. This parameter is
	// used to expand the total amount of ephemeral storage available, beyond the
	// default amount, for tasks hosted on Fargate. For more information, see [Using data volumes in tasks]in the
	// Amazon ECS Developer Guide.
	//
	// For tasks using the Fargate launch type, the task requires the following
	// platforms:
	//
	//   - Linux platform version 1.4.0 or later.
	//
	//   - Windows platform version 1.0.0 or later.
	//
	// [Using data volumes in tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html
	EphemeralStorage *types.EphemeralStorage

	// The Amazon Resource Name (ARN) of the task execution role that grants the
	// Amazon ECS container agent permission to make Amazon Web Services API calls on
	// your behalf. The task execution IAM role is required depending on the
	// requirements of your task. For more information, see [Amazon ECS task execution IAM role]in the Amazon Elastic
	// Container Service Developer Guide.
	//
	// [Amazon ECS task execution IAM role]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html
	ExecutionRoleArn *string

	// The Elastic Inference accelerators to use for the containers in the task.
	InferenceAccelerators []types.InferenceAccelerator

	// The IPC resource namespace to use for the containers in the task. The valid
	// values are host , task , or none . If host is specified, then all containers
	// within the tasks that specified the host IPC mode on the same container
	// instance share the same IPC resources with the host Amazon EC2 instance. If task
	// is specified, all containers within the specified task share the same IPC
	// resources. If none is specified, then IPC resources within the containers of a
	// task are private and not shared with other containers in a task or on the
	// container instance. If no value is specified, then the IPC resource namespace
	// sharing depends on the Docker daemon setting on the container instance. For more
	// information, see [IPC settings]in the Docker run reference.
	//
	// If the host IPC mode is used, be aware that there is a heightened risk of
	// undesired IPC namespace expose. For more information, see [Docker security].
	//
	// If you are setting namespaced kernel parameters using systemControls for the
	// containers in the task, the following will apply to your IPC resource namespace.
	// For more information, see [System Controls]in the Amazon Elastic Container Service Developer
	// Guide.
	//
	//   - For tasks that use the host IPC mode, IPC namespace related systemControls
	//   are not supported.
	//
	//   - For tasks that use the task IPC mode, IPC namespace related systemControls
	//   will apply to all containers within a task.
	//
	// This parameter is not supported for Windows containers or tasks run on Fargate.
	//
	// [System Controls]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html
	// [Docker security]: https://docs.docker.com/engine/security/security/
	// [IPC settings]: https://docs.docker.com/engine/reference/run/#ipc-settings---ipc
	IpcMode types.IpcMode

	// The amount of memory (in MiB) used by the task. It can be expressed as an
	// integer using MiB (for example , 1024 ) or as a string using GB (for example,
	// 1GB or 1 GB ) in a task definition. String values are converted to an integer
	// indicating the MiB when the task definition is registered.
	//
	// Task-level CPU and memory parameters are ignored for Windows containers. We
	// recommend specifying container-level resources for Windows containers.
	//
	// If using the EC2 launch type, this field is optional.
	//
	// If using the Fargate launch type, this field is required and you must use one
	// of the following values. This determines your range of supported values for the
	// cpu parameter.
	//
	// The CPU units cannot be less than 1 vCPU when you use Windows containers on
	// Fargate.
	//
	//   - 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	//   - 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values:
	//   512 (.5 vCPU)
	//
	//   - 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7
	//   GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	//   - Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) -
	//   Available cpu values: 2048 (2 vCPU)
	//
	//   - Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) -
	//   Available cpu values: 4096 (4 vCPU)
	//
	//   - Between 16 GB and 60 GB in 4 GB increments - Available cpu values: 8192 (8
	//   vCPU)
	//
	// This option requires Linux platform 1.4.0 or later.
	//
	//   - Between 32GB and 120 GB in 8 GB increments - Available cpu values: 16384 (16
	//   vCPU)
	//
	// This option requires Linux platform 1.4.0 or later.
	Memory *string

	// The Docker networking mode to use for the containers in the task. The valid
	// values are none , bridge , awsvpc , and host . If no network mode is specified,
	// the default is bridge .
	//
	// For Amazon ECS tasks on Fargate, the awsvpc network mode is required. For
	// Amazon ECS tasks on Amazon EC2 Linux instances, any network mode can be used.
	// For Amazon ECS tasks on Amazon EC2 Windows instances, or awsvpc can be used. If
	// the network mode is set to none , you cannot specify port mappings in your
	// container definitions, and the tasks containers do not have external
	// connectivity. The host and awsvpc network modes offer the highest networking
	// performance for containers because they use the EC2 network stack instead of the
	// virtualized network stack provided by the bridge mode.
	//
	// With the host and awsvpc network modes, exposed container ports are mapped
	// directly to the corresponding host port (for the host network mode) or the
	// attached elastic network interface port (for the awsvpc network mode), so you
	// cannot take advantage of dynamic host port mappings.
	//
	// When using the host network mode, you should not run containers using the root
	// user (UID 0). It is considered best practice to use a non-root user.
	//
	// If the network mode is awsvpc , the task is allocated an elastic network
	// interface, and you must specify a NetworkConfigurationvalue when you create a service or run a task
	// with the task definition. For more information, see [Task Networking]in the Amazon Elastic
	// Container Service Developer Guide.
	//
	// If the network mode is host , you cannot run multiple instantiations of the same
	// task on a single container instance when port mappings are used.
	//
	// For more information, see [Network settings] in the Docker run reference.
	//
	// [Task Networking]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html
	// [Network settings]: https://docs.docker.com/engine/reference/run/#network-settings
	NetworkMode types.NetworkMode

	// The process namespace to use for the containers in the task. The valid values
	// are host or task . On Fargate for Linux containers, the only valid value is task
	// . For example, monitoring sidecars might need pidMode to access information
	// about other containers running in the same task.
	//
	// If host is specified, all containers within the tasks that specified the host
	// PID mode on the same container instance share the same process namespace with
	// the host Amazon EC2 instance.
	//
	// If task is specified, all containers within the specified task share the same
	// process namespace.
	//
	// If no value is specified, the default is a private namespace for each
	// container. For more information, see [PID settings]in the Docker run reference.
	//
	// If the host PID mode is used, there's a heightened risk of undesired process
	// namespace exposure. For more information, see [Docker security].
	//
	// This parameter is not supported for Windows containers.
	//
	// This parameter is only supported for tasks that are hosted on Fargate if the
	// tasks are using platform version 1.4.0 or later (Linux). This isn't supported
	// for Windows containers on Fargate.
	//
	// [PID settings]: https://docs.docker.com/engine/reference/run/#pid-settings---pid
	// [Docker security]: https://docs.docker.com/engine/security/security/
	PidMode types.PidMode

	// An array of placement constraint objects to use for the task. You can specify a
	// maximum of 10 constraints for each task. This limit includes constraints in the
	// task definition and those specified at runtime.
	PlacementConstraints []types.TaskDefinitionPlacementConstraint

	// The configuration details for the App Mesh proxy.
	//
	// For tasks hosted on Amazon EC2 instances, the container instances require at
	// least version 1.26.0 of the container agent and at least version 1.26.0-1 of
	// the ecs-init package to use a proxy configuration. If your container instances
	// are launched from the Amazon ECS-optimized AMI version 20190301 or later, then
	// they contain the required versions of the container agent and ecs-init . For
	// more information, see [Amazon ECS-optimized AMI versions]in the Amazon Elastic Container Service Developer Guide.
	//
	// [Amazon ECS-optimized AMI versions]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-ami-versions.html
	ProxyConfiguration *types.ProxyConfiguration

	// The task launch type that Amazon ECS validates the task definition against. A
	// client exception is returned if the task definition doesn't validate against the
	// compatibilities specified. If no value is specified, the parameter is omitted
	// from the response.
	RequiresCompatibilities []types.Compatibility

	// The operating system that your tasks definitions run on. A platform family is
	// specified only for tasks using the Fargate launch type.
	RuntimePlatform *types.RuntimePlatform

	// The metadata that you apply to the task definition to help you categorize and
	// organize them. Each tag consists of a key and an optional value. You define both
	// of them.
	//
	// The following basic restrictions apply to tags:
	//
	//   - Maximum number of tags per resource - 50
	//
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//
	//   - Maximum key length - 128 Unicode characters in UTF-8
	//
	//   - Maximum value length - 256 Unicode characters in UTF-8
	//
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//
	//   - Tag keys and values are case-sensitive.
	//
	//   - Do not use aws: , AWS: , or any upper or lowercase combination of such as a
	//   prefix for either keys or values as it is reserved for Amazon Web Services use.
	//   You cannot edit or delete tag keys or values with this prefix. Tags with this
	//   prefix do not count against your tags per resource limit.
	Tags []types.Tag

	// The short name or full Amazon Resource Name (ARN) of the IAM role that
	// containers in this task can assume. All containers in this task are granted the
	// permissions that are specified in this role. For more information, see [IAM Roles for Tasks]in the
	// Amazon Elastic Container Service Developer Guide.
	//
	// [IAM Roles for Tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
	TaskRoleArn *string

	// A list of volume definitions in JSON format that containers in your task might
	// use.
	Volumes []types.Volume

	noSmithyDocumentSerde
}

type RegisterTaskDefinitionOutput struct {

	// The list of tags associated with the task definition.
	Tags []types.Tag

	// The full description of the registered task definition.
	TaskDefinition *types.TaskDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterTaskDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterTaskDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterTaskDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterTaskDefinition"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterTaskDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterTaskDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterTaskDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterTaskDefinition",
	}
}
