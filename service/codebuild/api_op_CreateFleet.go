// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a compute fleet.
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	if params == nil {
		params = &CreateFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleet", params, optFns, c.addOperationCreateFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetInput struct {

	// The initial number of machines allocated to the ﬂeet, which deﬁnes the number
	// of builds that can run in parallel.
	//
	// This member is required.
	BaseCapacity *int32

	// Information about the compute resources the compute fleet uses. Available
	// values include:
	//
	//   - BUILD_GENERAL1_SMALL : Use up to 3 GB memory and 2 vCPUs for builds.
	//
	//   - BUILD_GENERAL1_MEDIUM : Use up to 7 GB memory and 4 vCPUs for builds.
	//
	//   - BUILD_GENERAL1_LARGE : Use up to 16 GB memory and 8 vCPUs for builds,
	//   depending on your environment type.
	//
	//   - BUILD_GENERAL1_XLARGE : Use up to 70 GB memory and 36 vCPUs for builds,
	//   depending on your environment type.
	//
	//   - BUILD_GENERAL1_2XLARGE : Use up to 145 GB memory, 72 vCPUs, and 824 GB of
	//   SSD storage for builds. This compute type supports Docker images up to 100 GB
	//   uncompressed.
	//
	// If you use BUILD_GENERAL1_SMALL :
	//
	//   - For environment type LINUX_CONTAINER , you can use up to 3 GB memory and 2
	//   vCPUs for builds.
	//
	//   - For environment type LINUX_GPU_CONTAINER , you can use up to 16 GB memory, 4
	//   vCPUs, and 1 NVIDIA A10G Tensor Core GPU for builds.
	//
	//   - For environment type ARM_CONTAINER , you can use up to 4 GB memory and 2
	//   vCPUs on ARM-based processors for builds.
	//
	// If you use BUILD_GENERAL1_LARGE :
	//
	//   - For environment type LINUX_CONTAINER , you can use up to 15 GB memory and 8
	//   vCPUs for builds.
	//
	//   - For environment type LINUX_GPU_CONTAINER , you can use up to 255 GB memory,
	//   32 vCPUs, and 4 NVIDIA Tesla V100 GPUs for builds.
	//
	//   - For environment type ARM_CONTAINER , you can use up to 16 GB memory and 8
	//   vCPUs on ARM-based processors for builds.
	//
	// For more information, see [Build environment compute types] in the CodeBuild User Guide.
	//
	// [Build environment compute types]: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html
	//
	// This member is required.
	ComputeType types.ComputeType

	// The environment type of the compute fleet.
	//
	//   - The environment type ARM_CONTAINER is available only in regions US East (N.
	//   Virginia), US East (Ohio), US West (Oregon), EU (Ireland), Asia Pacific
	//   (Mumbai), Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney),
	//   EU (Frankfurt), and South America (São Paulo).
	//
	//   - The environment type LINUX_CONTAINER is available only in regions US East
	//   (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt),
	//   Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South
	//   America (São Paulo), and Asia Pacific (Mumbai).
	//
	//   - The environment type LINUX_GPU_CONTAINER is available only in regions US
	//   East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU
	//   (Frankfurt), Asia Pacific (Tokyo), and Asia Pacific (Sydney).
	//
	//   - The environment type WINDOWS_SERVER_2019_CONTAINER is available only in
	//   regions US East (N. Virginia), US East (Ohio), US West (Oregon), Asia Pacific
	//   (Sydney), Asia Pacific (Tokyo), Asia Pacific (Mumbai) and EU (Ireland).
	//
	//   - The environment type WINDOWS_SERVER_2022_CONTAINER is available only in
	//   regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland),
	//   EU (Frankfurt), Asia Pacific (Sydney), Asia Pacific (Singapore), Asia Pacific
	//   (Tokyo), South America (São Paulo) and Asia Pacific (Mumbai).
	//
	// For more information, see [Build environment compute types] in the CodeBuild user guide.
	//
	// [Build environment compute types]: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html
	//
	// This member is required.
	EnvironmentType types.EnvironmentType

	// The name of the compute fleet.
	//
	// This member is required.
	Name *string

	// The service role associated with the compute fleet. For more information, see [Allow a user to add a permission policy for a fleet service role]
	// in the CodeBuild User Guide.
	//
	// [Allow a user to add a permission policy for a fleet service role]: https://docs.aws.amazon.com/codebuild/latest/userguide/auth-and-access-control-iam-identity-based-access-control.html#customer-managed-policies-example-permission-policy-fleet-service-role.html
	FleetServiceRole *string

	// The compute fleet overflow behavior.
	//
	//   - For overflow behavior QUEUE , your overflow builds need to wait on the
	//   existing fleet instance to become available.
	//
	//   - For overflow behavior ON_DEMAND , your overflow builds run on CodeBuild
	//   on-demand.
	//
	// If you choose to set your overflow behavior to on-demand while creating a
	//   VPC-connected fleet, make sure that you add the required VPC permissions to your
	//   project service role. For more information, see [Example policy statement to allow CodeBuild access to Amazon Web Services services required to create a VPC network interface].
	//
	// [Example policy statement to allow CodeBuild access to Amazon Web Services services required to create a VPC network interface]: https://docs.aws.amazon.com/codebuild/latest/userguide/auth-and-access-control-iam-identity-based-access-control.html#customer-managed-policies-example-create-vpc-network-interface
	OverflowBehavior types.FleetOverflowBehavior

	// The scaling configuration of the compute fleet.
	ScalingConfiguration *types.ScalingConfigurationInput

	// A list of tag key and value pairs associated with this compute fleet.
	//
	// These tags are available for use by Amazon Web Services services that support
	// CodeBuild build project tags.
	Tags []types.Tag

	// Information about the VPC configuration that CodeBuild accesses.
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateFleetOutput struct {

	// Information about the compute fleet
	Fleet *types.Fleet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFleet"); err != nil {
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
	if err = addOpCreateFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFleet",
	}
}
