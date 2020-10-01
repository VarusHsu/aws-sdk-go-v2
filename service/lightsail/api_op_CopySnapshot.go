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

// Copies a manual snapshot of an instance or disk as another manual snapshot, or
// copies an automatic snapshot of an instance or disk as a manual snapshot. This
// operation can also be used to copy a manual or automatic snapshot of an instance
// or a disk from one AWS Region to another in Amazon Lightsail. When copying a
// manual snapshot, be sure to define the source region, source snapshot name, and
// target snapshot name parameters. When copying an automatic snapshot, be sure to
// define the source region, source resource name, target snapshot name, and either
// the restore date or the use latest restorable auto snapshot parameters.
func (c *Client) CopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*Options)) (*CopySnapshotOutput, error) {
	stack := middleware.NewStack("CopySnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCopySnapshotMiddlewares(stack)
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
	addOpCopySnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopySnapshot(options.Region), middleware.Before)
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
			OperationName: "CopySnapshot",
			Err:           err,
		}
	}
	out := result.(*CopySnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopySnapshotInput struct {

	// The AWS Region where the source manual or automatic snapshot is located.
	//
	// This member is required.
	SourceRegion types.RegionName

	// The date of the source automatic snapshot to copy. Use the get auto snapshots
	// operation to identify the dates of the available automatic snapshots.
	// Constraints:
	//
	//     * Must be specified in YYYY-MM-DD format.
	//
	//     * This
	// parameter cannot be defined together with the use latest restorable auto
	// snapshot parameter. The restore date and use latest restorable auto snapshot
	// parameters are mutually exclusive.
	//
	//     * Define this parameter only when
	// copying an automatic snapshot as a manual snapshot. For more information, see
	// the Lightsail Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-keeping-automatic-snapshots).
	RestoreDate *string

	// A Boolean value to indicate whether to use the latest available automatic
	// snapshot of the specified source instance or disk. Constraints:
	//
	//     * This
	// parameter cannot be defined together with the restore date parameter. The use
	// latest restorable auto snapshot and restore date parameters are mutually
	// exclusive.
	//
	//     * Define this parameter only when copying an automatic snapshot
	// as a manual snapshot. For more information, see the Lightsail Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-keeping-automatic-snapshots).
	UseLatestRestorableAutoSnapshot *bool

	// The name of the new manual snapshot to be created as a copy.
	//
	// This member is required.
	TargetSnapshotName *string

	// The name of the source manual snapshot to copy. Constraint:
	//
	//     * Define this
	// parameter only when copying a manual snapshot as another manual snapshot.
	SourceSnapshotName *string

	// The name of the source instance or disk from which the source automatic snapshot
	// was created. Constraint:
	//
	//     * Define this parameter only when copying an
	// automatic snapshot as a manual snapshot. For more information, see the Lightsail
	// Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-keeping-automatic-snapshots).
	SourceResourceName *string
}

type CopySnapshotOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCopySnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCopySnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopySnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopySnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CopySnapshot",
	}
}
