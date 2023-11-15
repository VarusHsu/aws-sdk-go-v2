// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Route 53 Resolver on an Outpost.
func (c *Client) CreateOutpostResolver(ctx context.Context, params *CreateOutpostResolverInput, optFns ...func(*Options)) (*CreateOutpostResolverOutput, error) {
	if params == nil {
		params = &CreateOutpostResolverInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOutpostResolver", params, optFns, c.addOperationCreateOutpostResolverMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOutpostResolverOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOutpostResolverInput struct {

	// A unique string that identifies the request and that allows failed requests to
	// be retried without the risk of running the operation twice. CreatorRequestId
	// can be any unique string, for example, a date/time stamp.
	//
	// This member is required.
	CreatorRequestId *string

	// A friendly name that lets you easily find a configuration in the Resolver
	// dashboard in the Route 53 console.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the Outpost. If you specify this, you must
	// also specify a value for the PreferredInstanceType .
	//
	// This member is required.
	OutpostArn *string

	// The Amazon EC2 instance type. If you specify this, you must also specify a
	// value for the OutpostArn .
	//
	// This member is required.
	PreferredInstanceType *string

	// Number of Amazon EC2 instances for the Resolver on Outpost. The default and
	// minimal value is 4.
	InstanceCount *int32

	// A string that helps identify the Route 53 Resolvers on Outpost.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateOutpostResolverOutput struct {

	// Information about the CreateOutpostResolver request, including the status of
	// the request.
	OutpostResolver *types.OutpostResolver

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOutpostResolverMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateOutpostResolver{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateOutpostResolver{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateOutpostResolver"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpCreateOutpostResolverValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOutpostResolver(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opCreateOutpostResolver(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateOutpostResolver",
	}
}
