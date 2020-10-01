// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the export snapshot record created as a result of the export snapshot
// operation. An export snapshot record can be used to create a new Amazon EC2
// instance and its related resources with the create cloud formation stack
// operation.
func (c *Client) GetExportSnapshotRecords(ctx context.Context, params *GetExportSnapshotRecordsInput, optFns ...func(*Options)) (*GetExportSnapshotRecordsOutput, error) {
	stack := middleware.NewStack("GetExportSnapshotRecords", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetExportSnapshotRecordsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetExportSnapshotRecords(options.Region), middleware.Before)
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
			OperationName: "GetExportSnapshotRecords",
			Err:           err,
		}
	}
	out := result.(*GetExportSnapshotRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExportSnapshotRecordsInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetExportSnapshotRecords request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string
}

type GetExportSnapshotRecordsOutput struct {

	// A list of objects describing the export snapshot records.
	ExportSnapshotRecords []*types.ExportSnapshotRecord

	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetExportSnapshotRecords request and specify
	// the next page token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetExportSnapshotRecordsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetExportSnapshotRecords{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetExportSnapshotRecords{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetExportSnapshotRecords(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetExportSnapshotRecords",
	}
}
