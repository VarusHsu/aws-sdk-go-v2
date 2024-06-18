// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new hardware security module (HSM) in the specified AWS CloudHSM
// cluster.
func (c *Client) CreateHsm(ctx context.Context, params *CreateHsmInput, optFns ...func(*Options)) (*CreateHsmOutput, error) {
	if params == nil {
		params = &CreateHsmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHsm", params, optFns, c.addOperationCreateHsmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHsmInput struct {

	// The Availability Zone where you are creating the HSM. To find the cluster's
	// Availability Zones, use DescribeClusters.
	//
	// This member is required.
	AvailabilityZone *string

	// The identifier (ID) of the HSM's cluster. To find the cluster ID, use DescribeClusters.
	//
	// This member is required.
	ClusterId *string

	// The HSM's IP address. If you specify an IP address, use an available address
	// from the subnet that maps to the Availability Zone where you are creating the
	// HSM. If you don't specify an IP address, one is chosen for you from that subnet.
	IpAddress *string

	noSmithyDocumentSerde
}

type CreateHsmOutput struct {

	// Information about the HSM that was created.
	Hsm *types.Hsm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHsmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHsm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHsm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHsm"); err != nil {
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
	if err = addOpCreateHsmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHsm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHsm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHsm",
	}
}
