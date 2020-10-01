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

// Exports an Amazon Lightsail instance or block storage disk snapshot to Amazon
// Elastic Compute Cloud (Amazon EC2). This operation results in an export snapshot
// record that can be used with the create cloud formation stack operation to
// create new Amazon EC2 instances. Exported instance snapshots appear in Amazon
// EC2 as Amazon Machine Images (AMIs), and the instance system disk appears as an
// Amazon Elastic Block Store (Amazon EBS) volume. Exported disk snapshots appear
// in Amazon EC2 as Amazon EBS volumes. Snapshots are exported to the same Amazon
// Web Services Region in Amazon EC2 as the source Lightsail snapshot. The export
// snapshot operation supports tag-based access control via resource tags applied
// to the resource identified by source snapshot name. For more information, see
// the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
// Use the get instance snapshots or get disk snapshots operations to get a list of
// snapshots that you can export to Amazon EC2.
func (c *Client) ExportSnapshot(ctx context.Context, params *ExportSnapshotInput, optFns ...func(*Options)) (*ExportSnapshotOutput, error) {
	stack := middleware.NewStack("ExportSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpExportSnapshotMiddlewares(stack)
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
	addOpExportSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportSnapshot(options.Region), middleware.Before)
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
			OperationName: "ExportSnapshot",
			Err:           err,
		}
	}
	out := result.(*ExportSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportSnapshotInput struct {

	// The name of the instance or disk snapshot to be exported to Amazon EC2.
	//
	// This member is required.
	SourceSnapshotName *string
}

type ExportSnapshotOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpExportSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpExportSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpExportSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opExportSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "ExportSnapshot",
	}
}
