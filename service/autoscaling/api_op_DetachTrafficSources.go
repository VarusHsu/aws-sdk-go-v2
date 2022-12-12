// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Reserved for use with Amazon VPC Lattice, which is in preview and subject to
// change. Do not use this API for production workloads. This API is also subject
// to change. Detaches one or more traffic sources from the specified Auto Scaling
// group.
func (c *Client) DetachTrafficSources(ctx context.Context, params *DetachTrafficSourcesInput, optFns ...func(*Options)) (*DetachTrafficSourcesOutput, error) {
	if params == nil {
		params = &DetachTrafficSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachTrafficSources", params, optFns, c.addOperationDetachTrafficSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachTrafficSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachTrafficSourcesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The unique identifiers of one or more traffic sources you are detaching. You can
	// specify up to 10 traffic sources. Currently, you must specify an Amazon Resource
	// Name (ARN) for an existing VPC Lattice target group. When you detach a target
	// group, it enters the Removing state while deregistering the instances in the
	// group. When all instances are deregistered, then you can no longer describe the
	// target group using the DescribeTrafficSources API call. The instances continue
	// to run.
	//
	// This member is required.
	TrafficSources []types.TrafficSourceIdentifier

	noSmithyDocumentSerde
}

type DetachTrafficSourcesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachTrafficSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDetachTrafficSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDetachTrafficSources{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDetachTrafficSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachTrafficSources(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDetachTrafficSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DetachTrafficSources",
	}
}