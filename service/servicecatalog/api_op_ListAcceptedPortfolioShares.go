// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all portfolios for which sharing was accepted by this account.
func (c *Client) ListAcceptedPortfolioShares(ctx context.Context, params *ListAcceptedPortfolioSharesInput, optFns ...func(*Options)) (*ListAcceptedPortfolioSharesOutput, error) {
	stack := middleware.NewStack("ListAcceptedPortfolioShares", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAcceptedPortfolioSharesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAcceptedPortfolioShares(options.Region), middleware.Before)
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
			OperationName: "ListAcceptedPortfolioShares",
			Err:           err,
		}
	}
	out := result.(*ListAcceptedPortfolioSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAcceptedPortfolioSharesInput struct {

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize *int32

	// The type of shared portfolios to list. The default is to list imported
	// portfolios.
	//
	//     * AWS_ORGANIZATIONS - List portfolios shared by the master
	// account of your organization
	//
	//     * AWS_SERVICECATALOG - List default
	// portfolios
	//
	//     * IMPORTED - List imported portfolios
	PortfolioShareType types.PortfolioShareType

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
}

type ListAcceptedPortfolioSharesOutput struct {

	// Information about the portfolios.
	PortfolioDetails []*types.PortfolioDetail

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAcceptedPortfolioSharesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAcceptedPortfolioShares{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAcceptedPortfolioShares{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAcceptedPortfolioShares(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListAcceptedPortfolioShares",
	}
}
