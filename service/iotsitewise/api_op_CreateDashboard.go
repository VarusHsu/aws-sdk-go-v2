// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a dashboard in an AWS IoT SiteWise Monitor project.
func (c *Client) CreateDashboard(ctx context.Context, params *CreateDashboardInput, optFns ...func(*Options)) (*CreateDashboardOutput, error) {
	stack := middleware.NewStack("CreateDashboard", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDashboardMiddlewares(stack)
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
	addIdempotencyToken_opCreateDashboardMiddleware(stack, options)
	addOpCreateDashboardValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDashboard(options.Region), middleware.Before)
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
			OperationName: "CreateDashboard",
			Err:           err,
		}
	}
	out := result.(*CreateDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDashboardInput struct {

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	// The dashboard definition specified in a JSON literal. For detailed information,
	// see Creating Dashboards (CLI)
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// This member is required.
	DashboardDefinition *string

	// A list of key-value pairs that contain metadata for the dashboard. For more
	// information, see Tagging your AWS IoT SiteWise resources
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]*string

	// A friendly name for the dashboard.
	//
	// This member is required.
	DashboardName *string

	// The ID of the project in which to create the dashboard.
	//
	// This member is required.
	ProjectId *string

	// A description for the dashboard.
	DashboardDescription *string
}

type CreateDashboardOutput struct {

	// The ID of the dashboard.
	//
	// This member is required.
	DashboardId *string

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the dashboard, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:dashboard/${DashboardId}
	//
	// This member is required.
	DashboardArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDashboardMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDashboard{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDashboard{}, middleware.After)
}

type idempotencyToken_initializeOpCreateDashboard struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDashboard) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDashboard) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDashboardInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDashboardInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDashboardMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateDashboard{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDashboard(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "CreateDashboard",
	}
}
