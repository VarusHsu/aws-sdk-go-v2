// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an import volume task using metadata from the specified disk image.For
// more information, see Importing Disks to Amazon EBS
// (https://docs.aws.amazon.com/AWSEC2/latest/CommandLineReference/importing-your-volumes-into-amazon-ebs.html).
// For information about the import manifest referenced by this API action, see VM
// Import Manifest
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
func (c *Client) ImportVolume(ctx context.Context, params *ImportVolumeInput, optFns ...func(*Options)) (*ImportVolumeOutput, error) {
	stack := middleware.NewStack("ImportVolume", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpImportVolumeMiddlewares(stack)
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
	addOpImportVolumeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportVolume(options.Region), middleware.Before)
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
			OperationName: "ImportVolume",
			Err:           err,
		}
	}
	out := result.(*ImportVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportVolumeInput struct {

	// The Availability Zone for the resulting EBS volume.
	//
	// This member is required.
	AvailabilityZone *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// A description of the volume.
	Description *string

	// The volume size.
	//
	// This member is required.
	Volume *types.VolumeDetail

	// The disk image.
	//
	// This member is required.
	Image *types.DiskImageDetail
}

type ImportVolumeOutput struct {

	// Information about the conversion task.
	ConversionTask *types.ConversionTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpImportVolumeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpImportVolume{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpImportVolume{}, middleware.After)
}

func newServiceMetadataMiddleware_opImportVolume(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ImportVolume",
	}
}
