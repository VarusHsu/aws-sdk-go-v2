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

// Returns the CloudFormation stack record created as a result of the create cloud
// formation stack operation. An AWS CloudFormation stack is used to create a new
// Amazon EC2 instance from an exported Lightsail snapshot.
func (c *Client) GetCloudFormationStackRecords(ctx context.Context, params *GetCloudFormationStackRecordsInput, optFns ...func(*Options)) (*GetCloudFormationStackRecordsOutput, error) {
	stack := middleware.NewStack("GetCloudFormationStackRecords", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCloudFormationStackRecordsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCloudFormationStackRecords(options.Region), middleware.Before)
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
			OperationName: "GetCloudFormationStackRecords",
			Err:           err,
		}
	}
	out := result.(*GetCloudFormationStackRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCloudFormationStackRecordsInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetClouFormationStackRecords request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.
	PageToken *string
}

type GetCloudFormationStackRecordsOutput struct {

	// A list of objects describing the CloudFormation stack records.
	CloudFormationStackRecords []*types.CloudFormationStackRecord

	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetCloudFormationStackRecords request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCloudFormationStackRecordsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCloudFormationStackRecords{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCloudFormationStackRecords{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCloudFormationStackRecords(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetCloudFormationStackRecords",
	}
}
