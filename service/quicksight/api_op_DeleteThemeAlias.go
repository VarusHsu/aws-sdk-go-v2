// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the version of the theme that the specified theme alias points to. If
// you provide a specific alias, you delete the version of the theme that the alias
// points to.
func (c *Client) DeleteThemeAlias(ctx context.Context, params *DeleteThemeAliasInput, optFns ...func(*Options)) (*DeleteThemeAliasOutput, error) {
	stack := middleware.NewStack("DeleteThemeAlias", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteThemeAliasMiddlewares(stack)
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
	addOpDeleteThemeAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteThemeAlias(options.Region), middleware.Before)
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
			OperationName: "DeleteThemeAlias",
			Err:           err,
		}
	}
	out := result.(*DeleteThemeAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteThemeAliasInput struct {

	// The ID of the AWS account that contains the theme alias to delete.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the theme that the specified alias is for.
	//
	// This member is required.
	ThemeId *string

	// The unique name for the theme alias to delete.
	//
	// This member is required.
	AliasName *string
}

type DeleteThemeAliasOutput struct {

	// An ID for the theme associated with the deletion.
	ThemeId *string

	// The Amazon Resource Name (ARN) of the theme resource using the deleted alias.
	Arn *string

	// The AWS request ID for this operation.
	RequestId *string

	// The name for the theme alias.
	AliasName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteThemeAliasMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteThemeAlias{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteThemeAlias{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteThemeAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteThemeAlias",
	}
}
