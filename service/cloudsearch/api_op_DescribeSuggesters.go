// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the suggesters configured for a domain. A suggester enables you to display
// possible matches before users finish typing their queries. Can be limited to
// specific suggesters by name. By default, shows all suggesters and includes any
// pending changes to the configuration. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see [Getting Search Suggestions]
// in the Amazon CloudSearch Developer Guide.
//
// [Getting Search Suggestions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html
func (c *Client) DescribeSuggesters(ctx context.Context, params *DescribeSuggestersInput, optFns ...func(*Options)) (*DescribeSuggestersOutput, error) {
	if params == nil {
		params = &DescribeSuggestersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSuggesters", params, optFns, c.addOperationDescribeSuggestersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSuggestersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeSuggester operation. Specifies the name of the
// domain you want to describe. To restrict the response to particular suggesters,
// specify the names of the suggesters you want to describe. To show the active
// configuration and exclude any pending changes, set the Deployed option to true .
type DescribeSuggestersInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// Whether to display the deployed configuration ( true ) or include any pending
	// changes ( false ). Defaults to false .
	Deployed *bool

	// The suggesters you want to describe.
	SuggesterNames []string

	noSmithyDocumentSerde
}

// The result of a DescribeSuggesters request.
type DescribeSuggestersOutput struct {

	// The suggesters configured for the domain specified in the request.
	//
	// This member is required.
	Suggesters []types.SuggesterStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSuggestersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSuggesters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSuggesters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSuggesters"); err != nil {
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
	if err = addOpDescribeSuggestersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSuggesters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSuggesters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSuggesters",
	}
}
