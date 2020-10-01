// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an image builder. An image builder is a virtual machine that is used to
// create an image. The initial state of the builder is PENDING. When it is ready,
// the state is RUNNING.
func (c *Client) CreateImageBuilder(ctx context.Context, params *CreateImageBuilderInput, optFns ...func(*Options)) (*CreateImageBuilderOutput, error) {
	stack := middleware.NewStack("CreateImageBuilder", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateImageBuilderMiddlewares(stack)
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
	addOpCreateImageBuilderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateImageBuilder(options.Region), middleware.Before)
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
			OperationName: "CreateImageBuilder",
			Err:           err,
		}
	}
	out := result.(*CreateImageBuilderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateImageBuilderInput struct {

	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool

	// The instance type to use when launching the image builder. The following
	// instance types are available:
	//
	//     * stream.standard.medium
	//
	//     *
	// stream.standard.large
	//
	//     * stream.compute.large
	//
	//     * stream.compute.xlarge
	//
	//
	// * stream.compute.2xlarge
	//
	//     * stream.compute.4xlarge
	//
	//     *
	// stream.compute.8xlarge
	//
	//     * stream.memory.large
	//
	//     * stream.memory.xlarge
	//
	//
	// * stream.memory.2xlarge
	//
	//     * stream.memory.4xlarge
	//
	//     *
	// stream.memory.8xlarge
	//
	//     * stream.graphics-design.large
	//
	//     *
	// stream.graphics-design.xlarge
	//
	//     * stream.graphics-design.2xlarge
	//
	//     *
	// stream.graphics-design.4xlarge
	//
	//     * stream.graphics-desktop.2xlarge
	//
	//     *
	// stream.graphics-pro.4xlarge
	//
	//     * stream.graphics-pro.8xlarge
	//
	//     *
	// stream.graphics-pro.16xlarge
	//
	// This member is required.
	InstanceType *string

	// The VPC configuration for the image builder. You can specify only one subnet.
	VpcConfig *types.VpcConfig

	// The Amazon Resource Name (ARN) of the IAM role to apply to the image builder. To
	// assume a role, the image builder calls the AWS Security Token Service (STS)
	// AssumeRole API operation and passes the ARN of the role to use. The operation
	// creates a new session with temporary credentials. AppStream 2.0 retrieves the
	// temporary credentials and creates the AppStream_Machine_Role credential profile
	// on the instance.  <p>For more information, see <a
	// href="https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html">Using
	// an IAM Role to Grant Permissions to Applications and Scripts Running on
	// AppStream 2.0 Streaming Instances</a> in the <i>Amazon AppStream 2.0
	// Administration Guide</i>.</p>
	IamRoleArn *string

	// The description to display.
	Description *string

	// The image builder name to display.
	DisplayName *string

	// The tags to associate with the image builder. A tag is a key-value pair, and the
	// value is optional. For example, Environment=Test. If you do not specify a value,
	// Environment=.  <p>Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following special characters: </p> <p>_ . : / =
	// + \ - @</p> <p>If you do not specify a value, the value is set to an empty
	// string.</p> <p>For more information about tags, see <a
	// href="https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html">Tagging
	// Your Resources</a> in the <i>Amazon AppStream 2.0 Administration Guide</i>.</p>
	Tags map[string]*string

	// The name of the image used to create the image builder.
	ImageName *string

	// The list of interface VPC endpoint (interface endpoint) objects. Administrators
	// can connect to the image builder only through the specified endpoints.
	AccessEndpoints []*types.AccessEndpoint

	// The ARN of the public, private, or shared image to use.
	ImageArn *string

	// The version of the AppStream 2.0 agent to use for this image builder. To use the
	// latest version of the AppStream 2.0 agent, specify [LATEST].
	AppstreamAgentVersion *string

	// The name of the directory and organizational unit (OU) to use to join the image
	// builder to a Microsoft Active Directory domain.
	DomainJoinInfo *types.DomainJoinInfo

	// A unique name for the image builder.
	//
	// This member is required.
	Name *string
}

type CreateImageBuilderOutput struct {

	// Information about the image builder.
	ImageBuilder *types.ImageBuilder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateImageBuilderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateImageBuilder{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateImageBuilder{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateImageBuilder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateImageBuilder",
	}
}
