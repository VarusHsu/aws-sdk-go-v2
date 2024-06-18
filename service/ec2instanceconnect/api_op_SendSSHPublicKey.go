// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2instanceconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Pushes an SSH public key to the specified EC2 instance for use by the specified
// user. The key remains for 60 seconds. For more information, see [Connect to your Linux instance using EC2 Instance Connect]in the Amazon
// EC2 User Guide.
//
// [Connect to your Linux instance using EC2 Instance Connect]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Connect-using-EC2-Instance-Connect.html
func (c *Client) SendSSHPublicKey(ctx context.Context, params *SendSSHPublicKeyInput, optFns ...func(*Options)) (*SendSSHPublicKeyOutput, error) {
	if params == nil {
		params = &SendSSHPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendSSHPublicKey", params, optFns, c.addOperationSendSSHPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendSSHPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendSSHPublicKeyInput struct {

	// The ID of the EC2 instance.
	//
	// This member is required.
	InstanceId *string

	// The OS user on the EC2 instance for whom the key can be used to authenticate.
	//
	// This member is required.
	InstanceOSUser *string

	// The public key material. To use the public key, you must have the matching
	// private key.
	//
	// This member is required.
	SSHPublicKey *string

	// The Availability Zone in which the EC2 instance was launched.
	AvailabilityZone *string

	noSmithyDocumentSerde
}

type SendSSHPublicKeyOutput struct {

	// The ID of the request. Please provide this ID when contacting AWS Support for
	// assistance.
	RequestId *string

	// Is true if the request succeeds and an error otherwise.
	Success bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendSSHPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendSSHPublicKey"); err != nil {
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
	if err = addOpSendSSHPublicKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendSSHPublicKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendSSHPublicKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendSSHPublicKey",
	}
}
