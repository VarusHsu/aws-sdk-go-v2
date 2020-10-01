// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a report group. A report group contains a collection of reports.
func (c *Client) CreateReportGroup(ctx context.Context, params *CreateReportGroupInput, optFns ...func(*Options)) (*CreateReportGroupOutput, error) {
	stack := middleware.NewStack("CreateReportGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateReportGroupMiddlewares(stack)
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
	addOpCreateReportGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReportGroup(options.Region), middleware.Before)
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
			OperationName: "CreateReportGroup",
			Err:           err,
		}
	}
	out := result.(*CreateReportGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReportGroupInput struct {

	// The type of report group.
	//
	// This member is required.
	Type types.ReportType

	// The name of the report group.
	//
	// This member is required.
	Name *string

	// A list of tag key and value pairs associated with this report group. These tags
	// are available for use by AWS services that support AWS CodeBuild report group
	// tags.
	Tags []*types.Tag

	// A ReportExportConfig object that contains information about where the report
	// group test results are exported.
	//
	// This member is required.
	ExportConfig *types.ReportExportConfig
}

type CreateReportGroupOutput struct {

	// Information about the report group that was created.
	ReportGroup *types.ReportGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateReportGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReportGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReportGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateReportGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "CreateReportGroup",
	}
}
