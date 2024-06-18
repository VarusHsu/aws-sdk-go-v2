// Code generated by smithy-go-codegen DO NOT EDIT.

package codeconnections

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	internalauthsmithy "github.com/aws/aws-sdk-go-v2/internal/auth/smithy"
	internalConfig "github.com/aws/aws-sdk-go-v2/internal/configsources"
	internalmiddleware "github.com/aws/aws-sdk-go-v2/internal/middleware"
	smithy "github.com/aws/smithy-go"
	smithyauth "github.com/aws/smithy-go/auth"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

const ServiceID = "CodeConnections"
const ServiceAPIVersion = "2023-12-01"

// Client provides the API client to make operations call for AWS CodeConnections.
type Client struct {
	options Options

	// Difference between the time reported by the server and the client
	timeOffset *atomic.Int64
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveDefaultLogger(&options)

	setResolvedDefaultsMode(&options)

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveEndpointResolverV2(&options)

	resolveAuthSchemeResolver(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	finalizeRetryMaxAttempts(&options)

	ignoreAnonymousAuth(&options)

	wrapWithAnonymousAuth(&options)

	resolveAuthSchemes(&options)

	client := &Client{
		options: options,
	}

	initializeTimeOffsetResolver(client)

	return client
}

// Options returns a copy of the client configuration.
//
// Callers SHOULD NOT perform mutations on any inner structures within client
// config. Config overrides should instead be made on a per-operation basis through
// functional options.
func (c *Client) Options() Options {
	return c.options.Copy()
}

func (c *Client) invokeOperation(ctx context.Context, opID string, params interface{}, optFns []func(*Options), stackFns ...func(*middleware.Stack, Options) error) (result interface{}, metadata middleware.Metadata, err error) {
	ctx = middleware.ClearStackValues(ctx)
	stack := middleware.NewStack(opID, smithyhttp.NewStackRequest)
	options := c.options.Copy()

	for _, fn := range optFns {
		fn(&options)
	}

	finalizeOperationRetryMaxAttempts(&options, *c)

	finalizeClientEndpointResolverOptions(&options)

	for _, fn := range stackFns {
		if err := fn(stack, options); err != nil {
			return nil, metadata, err
		}
	}

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, metadata, err
		}
	}

	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err = handler.Handle(ctx, params)
	if err != nil {
		err = &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: opID,
			Err:           err,
		}
	}
	return result, metadata, err
}

type operationInputKey struct{}

func setOperationInput(ctx context.Context, input interface{}) context.Context {
	return middleware.WithStackValue(ctx, operationInputKey{}, input)
}

func getOperationInput(ctx context.Context) interface{} {
	return middleware.GetStackValue(ctx, operationInputKey{})
}

type setOperationInputMiddleware struct {
}

func (*setOperationInputMiddleware) ID() string {
	return "setOperationInput"
}

func (m *setOperationInputMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	ctx = setOperationInput(ctx, in.Parameters)
	return next.HandleSerialize(ctx, in)
}

func addProtocolFinalizerMiddlewares(stack *middleware.Stack, options Options, operation string) error {
	if err := stack.Finalize.Add(&resolveAuthSchemeMiddleware{operation: operation, options: options}, middleware.Before); err != nil {
		return fmt.Errorf("add ResolveAuthScheme: %w", err)
	}
	if err := stack.Finalize.Insert(&getIdentityMiddleware{options: options}, "ResolveAuthScheme", middleware.After); err != nil {
		return fmt.Errorf("add GetIdentity: %v", err)
	}
	if err := stack.Finalize.Insert(&resolveEndpointV2Middleware{options: options}, "GetIdentity", middleware.After); err != nil {
		return fmt.Errorf("add ResolveEndpointV2: %v", err)
	}
	if err := stack.Finalize.Insert(&signRequestMiddleware{}, "ResolveEndpointV2", middleware.After); err != nil {
		return fmt.Errorf("add Signing: %w", err)
	}
	return nil
}
func resolveAuthSchemeResolver(options *Options) {
	if options.AuthSchemeResolver == nil {
		options.AuthSchemeResolver = &defaultAuthSchemeResolver{}
	}
}

func resolveAuthSchemes(options *Options) {
	if options.AuthSchemes == nil {
		options.AuthSchemes = []smithyhttp.AuthScheme{
			internalauth.NewHTTPAuthScheme("aws.auth#sigv4", &internalauthsmithy.V4SignerAdapter{
				Signer:     options.HTTPSignerV4,
				Logger:     options.Logger,
				LogSigning: options.ClientLogMode.IsSigning(),
			}),
		}
	}
}

type noSmithyDocumentSerde = smithydocument.NoSerde

type legacyEndpointContextSetter struct {
	LegacyResolver EndpointResolver
}

func (*legacyEndpointContextSetter) ID() string {
	return "legacyEndpointContextSetter"
}

func (m *legacyEndpointContextSetter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.LegacyResolver != nil {
		ctx = awsmiddleware.SetRequiresLegacyEndpoints(ctx, true)
	}

	return next.HandleInitialize(ctx, in)

}
func addlegacyEndpointContextSetter(stack *middleware.Stack, o Options) error {
	return stack.Initialize.Add(&legacyEndpointContextSetter{
		LegacyResolver: o.EndpointResolver,
	}, middleware.Before)
}

func resolveDefaultLogger(o *Options) {
	if o.Logger != nil {
		return
	}
	o.Logger = logging.Nop{}
}

func addSetLoggerMiddleware(stack *middleware.Stack, o Options) error {
	return middleware.AddSetLoggerMiddleware(stack, o.Logger)
}

func setResolvedDefaultsMode(o *Options) {
	if len(o.resolvedDefaultsMode) > 0 {
		return
	}

	var mode aws.DefaultsMode
	mode.SetFromString(string(o.DefaultsMode))

	if mode == aws.DefaultsModeAuto {
		mode = defaults.ResolveDefaultsModeAuto(o.Region, o.RuntimeEnvironment)
	}

	o.resolvedDefaultsMode = mode
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:                cfg.Region,
		DefaultsMode:          cfg.DefaultsMode,
		RuntimeEnvironment:    cfg.RuntimeEnvironment,
		HTTPClient:            cfg.HTTPClient,
		Credentials:           cfg.Credentials,
		APIOptions:            cfg.APIOptions,
		Logger:                cfg.Logger,
		ClientLogMode:         cfg.ClientLogMode,
		AppID:                 cfg.AppID,
		AccountIDEndpointMode: cfg.AccountIDEndpointMode,
	}
	resolveAWSRetryerProvider(cfg, &opts)
	resolveAWSRetryMaxAttempts(cfg, &opts)
	resolveAWSRetryMode(cfg, &opts)
	resolveAWSEndpointResolver(cfg, &opts)
	resolveUseDualStackEndpoint(cfg, &opts)
	resolveUseFIPSEndpoint(cfg, &opts)
	resolveBaseEndpoint(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	var buildable *awshttp.BuildableClient

	if o.HTTPClient != nil {
		var ok bool
		buildable, ok = o.HTTPClient.(*awshttp.BuildableClient)
		if !ok {
			return
		}
	} else {
		buildable = awshttp.NewBuildableClient()
	}

	modeConfig, err := defaults.GetModeConfiguration(o.resolvedDefaultsMode)
	if err == nil {
		buildable = buildable.WithDialerOptions(func(dialer *net.Dialer) {
			if dialerTimeout, ok := modeConfig.GetConnectTimeout(); ok {
				dialer.Timeout = dialerTimeout
			}
		})

		buildable = buildable.WithTransportOptions(func(transport *http.Transport) {
			if tlsHandshakeTimeout, ok := modeConfig.GetTLSNegotiationTimeout(); ok {
				transport.TLSHandshakeTimeout = tlsHandshakeTimeout
			}
		})
	}

	o.HTTPClient = buildable
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}

	if len(o.RetryMode) == 0 {
		modeConfig, err := defaults.GetModeConfiguration(o.resolvedDefaultsMode)
		if err == nil {
			o.RetryMode = modeConfig.RetryMode
		}
	}
	if len(o.RetryMode) == 0 {
		o.RetryMode = aws.RetryModeStandard
	}

	var standardOptions []func(*retry.StandardOptions)
	if v := o.RetryMaxAttempts; v != 0 {
		standardOptions = append(standardOptions, func(so *retry.StandardOptions) {
			so.MaxAttempts = v
		})
	}

	switch o.RetryMode {
	case aws.RetryModeAdaptive:
		var adaptiveOptions []func(*retry.AdaptiveModeOptions)
		if len(standardOptions) != 0 {
			adaptiveOptions = append(adaptiveOptions, func(ao *retry.AdaptiveModeOptions) {
				ao.StandardOptions = append(ao.StandardOptions, standardOptions...)
			})
		}
		o.Retryer = retry.NewAdaptiveMode(adaptiveOptions...)

	default:
		o.Retryer = retry.NewStandard(standardOptions...)
	}
}

func resolveAWSRetryerProvider(cfg aws.Config, o *Options) {
	if cfg.Retryer == nil {
		return
	}
	o.Retryer = cfg.Retryer()
}

func resolveAWSRetryMode(cfg aws.Config, o *Options) {
	if len(cfg.RetryMode) == 0 {
		return
	}
	o.RetryMode = cfg.RetryMode
}
func resolveAWSRetryMaxAttempts(cfg aws.Config, o *Options) {
	if cfg.RetryMaxAttempts == 0 {
		return
	}
	o.RetryMaxAttempts = cfg.RetryMaxAttempts
}

func finalizeRetryMaxAttempts(o *Options) {
	if o.RetryMaxAttempts == 0 {
		return
	}

	o.Retryer = retry.AddWithMaxAttempts(o.Retryer, o.RetryMaxAttempts)
}

func finalizeOperationRetryMaxAttempts(o *Options, client Client) {
	if v := o.RetryMaxAttempts; v == 0 || v == client.options.RetryMaxAttempts {
		return
	}

	o.Retryer = retry.AddWithMaxAttempts(o.Retryer, o.RetryMaxAttempts)
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil && cfg.EndpointResolverWithOptions == nil {
		return
	}
	o.EndpointResolver = withEndpointResolver(cfg.EndpointResolver, cfg.EndpointResolverWithOptions)
}

func addClientUserAgent(stack *middleware.Stack, options Options) error {
	ua, err := getOrAddRequestUserAgent(stack)
	if err != nil {
		return err
	}

	ua.AddSDKAgentKeyValue(awsmiddleware.APIMetadata, "codeconnections", goModuleVersion)
	if len(options.AppID) > 0 {
		ua.AddSDKAgentKey(awsmiddleware.ApplicationIdentifier, options.AppID)
	}

	return nil
}

func getOrAddRequestUserAgent(stack *middleware.Stack) (*awsmiddleware.RequestUserAgent, error) {
	id := (*awsmiddleware.RequestUserAgent)(nil).ID()
	mw, ok := stack.Build.Get(id)
	if !ok {
		mw = awsmiddleware.NewRequestUserAgent()
		if err := stack.Build.Add(mw, middleware.After); err != nil {
			return nil, err
		}
	}

	ua, ok := mw.(*awsmiddleware.RequestUserAgent)
	if !ok {
		return nil, fmt.Errorf("%T for %s middleware did not match expected type", mw, id)
	}

	return ua, nil
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time, optFns ...func(*v4.SignerOptions)) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = newDefaultV4Signer(*o)
}

func newDefaultV4Signer(o Options) *v4.Signer {
	return v4.NewSigner(func(so *v4.SignerOptions) {
		so.Logger = o.Logger
		so.LogSigning = o.ClientLogMode.IsSigning()
	})
}

func addClientRequestID(stack *middleware.Stack) error {
	return stack.Build.Add(&awsmiddleware.ClientRequestID{}, middleware.After)
}

func addComputeContentLength(stack *middleware.Stack) error {
	return stack.Build.Add(&smithyhttp.ComputeContentLength{}, middleware.After)
}

func addRawResponseToMetadata(stack *middleware.Stack) error {
	return stack.Deserialize.Add(&awsmiddleware.AddRawResponse{}, middleware.Before)
}

func addRecordResponseTiming(stack *middleware.Stack) error {
	return stack.Deserialize.Add(&awsmiddleware.RecordResponseTiming{}, middleware.After)
}
func addStreamingEventsPayload(stack *middleware.Stack) error {
	return stack.Finalize.Add(&v4.StreamingEventsPayload{}, middleware.Before)
}

func addUnsignedPayload(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&v4.UnsignedPayload{}, "ResolveEndpointV2", middleware.After)
}

func addComputePayloadSHA256(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&v4.ComputePayloadSHA256{}, "ResolveEndpointV2", middleware.After)
}

func addContentSHA256Header(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&v4.ContentSHA256Header{}, (*v4.ComputePayloadSHA256)(nil).ID(), middleware.After)
}

func addIsWaiterUserAgent(o *Options) {
	o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
		ua, err := getOrAddRequestUserAgent(stack)
		if err != nil {
			return err
		}

		ua.AddUserAgentFeature(awsmiddleware.UserAgentFeatureWaiter)
		return nil
	})
}

func addIsPaginatorUserAgent(o *Options) {
	o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
		ua, err := getOrAddRequestUserAgent(stack)
		if err != nil {
			return err
		}

		ua.AddUserAgentFeature(awsmiddleware.UserAgentFeaturePaginator)
		return nil
	})
}

func addRetry(stack *middleware.Stack, o Options) error {
	attempt := retry.NewAttemptMiddleware(o.Retryer, smithyhttp.RequestCloner, func(m *retry.Attempt) {
		m.LogAttempts = o.ClientLogMode.IsRetries()
	})
	if err := stack.Finalize.Insert(attempt, "Signing", middleware.Before); err != nil {
		return err
	}
	if err := stack.Finalize.Insert(&retry.MetricsHeader{}, attempt.ID(), middleware.After); err != nil {
		return err
	}
	return nil
}

// resolves dual-stack endpoint configuration
func resolveUseDualStackEndpoint(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := internalConfig.ResolveUseDualStackEndpoint(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.EndpointOptions.UseDualStackEndpoint = value
	}
	return nil
}

// resolves FIPS endpoint configuration
func resolveUseFIPSEndpoint(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := internalConfig.ResolveUseFIPSEndpoint(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.EndpointOptions.UseFIPSEndpoint = value
	}
	return nil
}

func resolveAccountID(identity smithyauth.Identity, mode aws.AccountIDEndpointMode) *string {
	if mode == aws.AccountIDEndpointModeDisabled {
		return nil
	}

	if ca, ok := identity.(*internalauthsmithy.CredentialsAdapter); ok && ca.Credentials.AccountID != "" {
		return aws.String(ca.Credentials.AccountID)
	}

	return nil
}

func addTimeOffsetBuild(stack *middleware.Stack, c *Client) error {
	mw := internalmiddleware.AddTimeOffsetMiddleware{Offset: c.timeOffset}
	if err := stack.Build.Add(&mw, middleware.After); err != nil {
		return err
	}
	return stack.Deserialize.Insert(&mw, "RecordResponseTiming", middleware.Before)
}
func initializeTimeOffsetResolver(c *Client) {
	c.timeOffset = new(atomic.Int64)
}

func checkAccountID(identity smithyauth.Identity, mode aws.AccountIDEndpointMode) error {
	switch mode {
	case aws.AccountIDEndpointModeUnset:
	case aws.AccountIDEndpointModePreferred:
	case aws.AccountIDEndpointModeDisabled:
	case aws.AccountIDEndpointModeRequired:
		if ca, ok := identity.(*internalauthsmithy.CredentialsAdapter); !ok {
			return fmt.Errorf("accountID is required but not set")
		} else if ca.Credentials.AccountID == "" {
			return fmt.Errorf("accountID is required but not set")
		}
	// default check in case invalid mode is configured through request config
	default:
		return fmt.Errorf("invalid accountID endpoint mode %s, must be preferred/required/disabled", mode)
	}

	return nil
}

func addUserAgentRetryMode(stack *middleware.Stack, options Options) error {
	ua, err := getOrAddRequestUserAgent(stack)
	if err != nil {
		return err
	}

	switch options.Retryer.(type) {
	case *retry.Standard:
		ua.AddUserAgentFeature(awsmiddleware.UserAgentFeatureRetryModeStandard)
	case *retry.AdaptiveMode:
		ua.AddUserAgentFeature(awsmiddleware.UserAgentFeatureRetryModeAdaptive)
	}
	return nil
}

func addRecursionDetection(stack *middleware.Stack) error {
	return stack.Build.Add(&awsmiddleware.RecursionDetection{}, middleware.After)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) error {
	return stack.Deserialize.Insert(&awsmiddleware.RequestIDRetriever{}, "OperationDeserializer", middleware.Before)

}

func addResponseErrorMiddleware(stack *middleware.Stack) error {
	return stack.Deserialize.Insert(&awshttp.ResponseErrorWrapper{}, "RequestIDRetriever", middleware.Before)

}

func addRequestResponseLogging(stack *middleware.Stack, o Options) error {
	return stack.Deserialize.Add(&smithyhttp.RequestResponseLogger{
		LogRequest:          o.ClientLogMode.IsRequest(),
		LogRequestWithBody:  o.ClientLogMode.IsRequestWithBody(),
		LogResponse:         o.ClientLogMode.IsResponse(),
		LogResponseWithBody: o.ClientLogMode.IsResponseWithBody(),
	}, middleware.After)
}

type disableHTTPSMiddleware struct {
	DisableHTTPS bool
}

func (*disableHTTPSMiddleware) ID() string {
	return "disableHTTPS"
}

func (m *disableHTTPSMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.DisableHTTPS && !smithyhttp.GetHostnameImmutable(ctx) {
		req.URL.Scheme = "http"
	}

	return next.HandleFinalize(ctx, in)
}

func addDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Finalize.Insert(&disableHTTPSMiddleware{
		DisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "ResolveEndpointV2", middleware.After)
}
