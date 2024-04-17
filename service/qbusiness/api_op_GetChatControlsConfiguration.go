// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about an chat controls configured for an existing Amazon Q
// Business application.
func (c *Client) GetChatControlsConfiguration(ctx context.Context, params *GetChatControlsConfigurationInput, optFns ...func(*Options)) (*GetChatControlsConfigurationOutput, error) {
	if params == nil {
		params = &GetChatControlsConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChatControlsConfiguration", params, optFns, c.addOperationGetChatControlsConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChatControlsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChatControlsConfigurationInput struct {

	// The identifier of the application for which the chat controls are configured.
	//
	// This member is required.
	ApplicationId *string

	// The maximum number of configured chat controls to return.
	MaxResults *int32

	// If the maxResults response was incomplete because there is more data to
	// retrieve, Amazon Q Business returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of Amazon Q Business chat
	// controls configured.
	NextToken *string

	noSmithyDocumentSerde
}

type GetChatControlsConfigurationOutput struct {

	// The phrases blocked from chat by your chat control configuration.
	BlockedPhrases *types.BlockedPhrasesConfiguration

	// The configuration details for CREATOR_MODE .
	CreatorModeConfiguration *types.AppliedCreatorModeConfiguration

	// If the maxResults response was incomplete because there is more data to
	// retrieve, Amazon Q Business returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of Amazon Q Business chat
	// controls configured.
	NextToken *string

	// The response scope configured for a Amazon Q Business application. This
	// determines whether your application uses its retrieval augmented generation
	// (RAG) system to generate answers only from your enterprise data, or also uses
	// the large language models (LLM) knowledge to respons to end user questions in
	// chat.
	ResponseScope types.ResponseScope

	// The topic specific controls configured for a Amazon Q Business application.
	TopicConfigurations []types.TopicConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetChatControlsConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetChatControlsConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetChatControlsConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetChatControlsConfiguration"); err != nil {
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
	if err = addOpGetChatControlsConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChatControlsConfiguration(options.Region), middleware.Before); err != nil {
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

// GetChatControlsConfigurationAPIClient is a client that implements the
// GetChatControlsConfiguration operation.
type GetChatControlsConfigurationAPIClient interface {
	GetChatControlsConfiguration(context.Context, *GetChatControlsConfigurationInput, ...func(*Options)) (*GetChatControlsConfigurationOutput, error)
}

var _ GetChatControlsConfigurationAPIClient = (*Client)(nil)

// GetChatControlsConfigurationPaginatorOptions is the paginator options for
// GetChatControlsConfiguration
type GetChatControlsConfigurationPaginatorOptions struct {
	// The maximum number of configured chat controls to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetChatControlsConfigurationPaginator is a paginator for
// GetChatControlsConfiguration
type GetChatControlsConfigurationPaginator struct {
	options   GetChatControlsConfigurationPaginatorOptions
	client    GetChatControlsConfigurationAPIClient
	params    *GetChatControlsConfigurationInput
	nextToken *string
	firstPage bool
}

// NewGetChatControlsConfigurationPaginator returns a new
// GetChatControlsConfigurationPaginator
func NewGetChatControlsConfigurationPaginator(client GetChatControlsConfigurationAPIClient, params *GetChatControlsConfigurationInput, optFns ...func(*GetChatControlsConfigurationPaginatorOptions)) *GetChatControlsConfigurationPaginator {
	if params == nil {
		params = &GetChatControlsConfigurationInput{}
	}

	options := GetChatControlsConfigurationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetChatControlsConfigurationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetChatControlsConfigurationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetChatControlsConfiguration page.
func (p *GetChatControlsConfigurationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetChatControlsConfigurationOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetChatControlsConfiguration(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetChatControlsConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetChatControlsConfiguration",
	}
}
