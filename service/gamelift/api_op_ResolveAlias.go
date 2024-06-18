// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attempts to retrieve a fleet ID that is associated with an alias. Specify a
// unique alias identifier.
//
// If the alias has a SIMPLE routing strategy, Amazon GameLift returns a fleet ID.
// If the alias has a TERMINAL routing strategy, the result is a
// TerminalRoutingStrategyException .
//
// # Related actions
//
// [All APIs by task]
//
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
func (c *Client) ResolveAlias(ctx context.Context, params *ResolveAliasInput, optFns ...func(*Options)) (*ResolveAliasOutput, error) {
	if params == nil {
		params = &ResolveAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResolveAlias", params, optFns, c.addOperationResolveAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResolveAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResolveAliasInput struct {

	// The unique identifier of the alias that you want to retrieve a fleet ID for.
	// You can use either the alias ID or ARN value.
	//
	// This member is required.
	AliasId *string

	noSmithyDocumentSerde
}

type ResolveAliasOutput struct {

	//  The Amazon Resource Name ([ARN] ) associated with the GameLift fleet resource that
	// this alias points to.
	//
	// [ARN]: https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html
	FleetArn *string

	// The fleet identifier that the alias is pointing to.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResolveAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResolveAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResolveAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ResolveAlias"); err != nil {
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
	if err = addOpResolveAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResolveAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResolveAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ResolveAlias",
	}
}
