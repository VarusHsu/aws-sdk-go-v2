// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an environment template for Proton. For more information, see [Environment Templates] in the
// Proton User Guide.
//
// You can create an environment template in one of the two following ways:
//
//   - Register and publish a standard environment template that instructs Proton
//     to deploy and manage environment infrastructure.
//
//   - Register and publish a customer managed environment template that connects
//     Proton to your existing provisioned infrastructure that you manage. Proton
//     doesn't manage your existing provisioned infrastructure. To create an
//     environment template for customer provisioned and managed infrastructure,
//     include the provisioning parameter and set the value to CUSTOMER_MANAGED . For
//     more information, see [Register and publish an environment template]in the Proton User Guide.
//
// [Register and publish an environment template]: https://docs.aws.amazon.com/proton/latest/userguide/template-create.html
// [Environment Templates]: https://docs.aws.amazon.com/proton/latest/userguide/ag-templates.html
func (c *Client) CreateEnvironmentTemplate(ctx context.Context, params *CreateEnvironmentTemplateInput, optFns ...func(*Options)) (*CreateEnvironmentTemplateOutput, error) {
	if params == nil {
		params = &CreateEnvironmentTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentTemplate", params, optFns, c.addOperationCreateEnvironmentTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentTemplateInput struct {

	// The name of the environment template.
	//
	// This member is required.
	Name *string

	// A description of the environment template.
	Description *string

	// The environment template name as displayed in the developer interface.
	DisplayName *string

	// A customer provided encryption key that Proton uses to encrypt data.
	EncryptionKey *string

	// When included, indicates that the environment template is for customer
	// provisioned and managed infrastructure.
	Provisioning types.Provisioning

	// An optional list of metadata items that you can associate with the Proton
	// environment template. A tag is a key-value pair.
	//
	// For more information, see [Proton resources and tagging] in the Proton User Guide.
	//
	// [Proton resources and tagging]: https://docs.aws.amazon.com/proton/latest/userguide/resources.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEnvironmentTemplateOutput struct {

	// The environment template detail data that's returned by Proton.
	//
	// This member is required.
	EnvironmentTemplate *types.EnvironmentTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateEnvironmentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateEnvironmentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEnvironmentTemplate"); err != nil {
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
	if err = addOpCreateEnvironmentTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEnvironmentTemplate",
	}
}
