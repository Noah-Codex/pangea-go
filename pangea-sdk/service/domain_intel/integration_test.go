// go:build integration
package domain_intel_test

import (
	"context"
	"testing"
	"time"

	"github.com/pangeacyber/pangea-go/pangea-sdk/internal/pangeatesting"
	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/domain_intel"
	"github.com/stretchr/testify/assert"
)

const (
	testingEnvironment = pangeatesting.Live
)

func intelDomainIntegrationCfg(t *testing.T) *pangea.Config {
	t.Helper()
	return pangeatesting.IntegrationConfig(t, testingEnvironment)
}

// Lookup domain malicious
func Test_Integration_DomainLookup(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainLookupInput{
		Domain:   "737updatesboeing.com",
		Raw:      true,
		Verbose:  true,
		Provider: "domaintools",
	}
	out, err := domainintel.Lookup(ctx, input)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	assert.NotNil(t, out)
	assert.NotNil(t, out.Result.Data)
	assert.Equal(t, out.Result.Data.Verdict, "malicious")
}

// Lookup domain unknown
func Test_Integration_DomainLookup_2(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainLookupInput{
		Domain:   "google.com",
		Raw:      true,
		Verbose:  true,
		Provider: "domaintools",
	}

	out, err := domainintel.Lookup(ctx, input)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	assert.NotNil(t, out)
	assert.NotNil(t, out.Result.Data)
	assert.Equal(t, out.Result.Data.Verdict, "benign")
}

// Test empty domain
func Test_Integration_DomainLookup_Error(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainLookupInput{
		Domain:   "",
		Raw:      true,
		Verbose:  true,
		Provider: "domaintools",
	}

	out, err := domainintel.Lookup(ctx, input)

	assert.Error(t, err)
	assert.Nil(t, out)
	err = err.(*pangea.APIError)
	apiErr := err.(*pangea.APIError)
	assert.Equal(t, len(apiErr.PangeaErrors.Errors), 1)
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Code, "BadFormatHostname")
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Detail, "'domain' must be a valid RFC1123 hostname")
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Source, "/domain")
}

// Bad auth token
func Test_Integration_DomainLookup_Error_BadAuthToken(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	cfg.Token = "notavalidtoken"
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainLookupInput{
		Domain:   "737updatesboeing.com",
		Raw:      true,
		Verbose:  true,
		Provider: "domaintools",
	}
	out, err := domainintel.Lookup(ctx, input)

	assert.Error(t, err)
	assert.Nil(t, out)
	apiErr := err.(*pangea.APIError)
	assert.Equal(t, apiErr.Err.Error(), "API error: Not authorized to access this resource.")
}

// Not valid provider
func Test_Integration_DomainLookup_Error_Provider(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainLookupInput{
		Domain:   "737updatesboeing.com",
		Raw:      true,
		Verbose:  true,
		Provider: "notaprovider",
	}
	out, err := domainintel.Lookup(ctx, input)

	assert.Error(t, err)
	assert.Nil(t, out)
	err = err.(*pangea.APIError)
	apiErr := err.(*pangea.APIError)
	assert.Equal(t, len(apiErr.PangeaErrors.Errors), 1)
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Code, "NotEnumMember")
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Source, "/provider")
}

// Reputation domain malicious
func Test_Integration_DomainReputation(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainReputationRequest{
		Domain:   "737updatesboeing.com",
		Raw:      true,
		Verbose:  true,
		Provider: "domaintools",
	}
	out, err := domainintel.Reputation(ctx, input)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	assert.NotNil(t, out)
	assert.NotNil(t, out.Result.Data)
	assert.Equal(t, out.Result.Data.Verdict, "malicious")
}

// Reputation domain unknown
func Test_Integration_DomainReputation_2(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainReputationRequest{
		Domain:   "google.com",
		Raw:      true,
		Verbose:  true,
		Provider: "domaintools",
	}

	out, err := domainintel.Reputation(ctx, input)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	assert.NotNil(t, out)
	assert.NotNil(t, out.Result.Data)
	assert.Equal(t, out.Result.Data.Verdict, "benign")
}

// Test empty domain
func Test_Integration_DomainReputation_Error(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainReputationRequest{
		Domain:   "",
		Raw:      true,
		Verbose:  true,
		Provider: "domaintools",
	}

	out, err := domainintel.Reputation(ctx, input)

	assert.Error(t, err)
	assert.Nil(t, out)
	err = err.(*pangea.APIError)
	apiErr := err.(*pangea.APIError)
	assert.Equal(t, len(apiErr.PangeaErrors.Errors), 1)
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Code, "BadFormatHostname")
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Detail, "'domain' must be a valid RFC1123 hostname")
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Source, "/domain")
}

// Bad auth token
func Test_Integration_DomainReputation_Error_BadAuthToken(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	cfg.Token = "notavalidtoken"
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainReputationRequest{
		Domain:   "737updatesboeing.com",
		Raw:      true,
		Verbose:  true,
		Provider: "domaintools",
	}
	out, err := domainintel.Reputation(ctx, input)

	assert.Error(t, err)
	assert.Nil(t, out)
	apiErr := err.(*pangea.APIError)
	assert.Equal(t, apiErr.Err.Error(), "API error: Not authorized to access this resource.")
}

// Not valid provider
func Test_Integration_DomainReputation_Error_Provider(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := intelDomainIntegrationCfg(t)
	domainintel := domain_intel.New(cfg)

	input := &domain_intel.DomainReputationRequest{
		Domain:   "737updatesboeing.com",
		Raw:      true,
		Verbose:  true,
		Provider: "notaprovider",
	}
	out, err := domainintel.Reputation(ctx, input)

	assert.Error(t, err)
	assert.Nil(t, out)
	err = err.(*pangea.APIError)
	apiErr := err.(*pangea.APIError)
	assert.Equal(t, len(apiErr.PangeaErrors.Errors), 1)
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Code, "NotEnumMember")
	assert.Equal(t, apiErr.PangeaErrors.Errors[0].Source, "/provider")
}
