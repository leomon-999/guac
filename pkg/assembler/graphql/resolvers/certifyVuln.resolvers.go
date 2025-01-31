package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"strings"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestCertifyVuln is the resolver for the ingestCertifyVuln field.
func (r *mutationResolver) IngestCertifyVuln(ctx context.Context, pkg model.PkgInputSpec, vulnerability model.VulnerabilityInputSpec, certifyVuln model.ScanMetadataInput) (*model.CertifyVuln, error) {
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	return r.Backend.IngestCertifyVuln(ctx, pkg,
		model.VulnerabilityInputSpec{Type: strings.ToLower(vulnerability.Type), VulnerabilityID: strings.ToLower(vulnerability.VulnerabilityID)},
		certifyVuln)
}

// IngestCertifyVulns is the resolver for the ingestCertifyVulns field.
func (r *mutationResolver) IngestCertifyVulns(ctx context.Context, pkgs []*model.PkgInputSpec, vulnerabilities []*model.VulnerabilityInputSpec, certifyVulns []*model.ScanMetadataInput) ([]*model.CertifyVuln, error) {
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	var lowercaseVulnInputList []*model.VulnerabilityInputSpec
	for _, v := range vulnerabilities {
		lowercaseVulnInput := model.VulnerabilityInputSpec{
			Type:            strings.ToLower(v.Type),
			VulnerabilityID: strings.ToLower(v.VulnerabilityID),
		}
		lowercaseVulnInputList = append(lowercaseVulnInputList, &lowercaseVulnInput)
	}
	return r.Backend.IngestCertifyVulns(ctx, pkgs, lowercaseVulnInputList, certifyVulns)
}

// CertifyVuln is the resolver for the CertifyVuln field.
func (r *queryResolver) CertifyVuln(ctx context.Context, certifyVulnSpec model.CertifyVulnSpec) ([]*model.CertifyVuln, error) {
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase

	if certifyVulnSpec.Vulnerability != nil {

		if certifyVulnSpec.Vulnerability.NoVuln != nil && !*certifyVulnSpec.Vulnerability.NoVuln {
			if certifyVulnSpec.Vulnerability.Type != nil && *certifyVulnSpec.Vulnerability.Type == "novuln" {
				return []*model.CertifyVuln{}, gqlerror.Errorf("novuln boolean set to false, cannot specify vulnerability type to be novuln")
			}
		}

		lowercaseVulnFilter := model.VulnerabilitySpec{
			ID:              certifyVulnSpec.Vulnerability.ID,
			Type:            toLower(certifyVulnSpec.Vulnerability.Type),
			VulnerabilityID: toLower(certifyVulnSpec.Vulnerability.VulnerabilityID),
			NoVuln:          certifyVulnSpec.Vulnerability.NoVuln,
		}

		lowercaseCertifyVulnFilter := model.CertifyVulnSpec{
			ID:             certifyVulnSpec.ID,
			Package:        certifyVulnSpec.Package,
			Vulnerability:  &lowercaseVulnFilter,
			TimeScanned:    certifyVulnSpec.TimeScanned,
			DbURI:          certifyVulnSpec.DbURI,
			DbVersion:      certifyVulnSpec.DbVersion,
			ScannerURI:     certifyVulnSpec.ScannerURI,
			ScannerVersion: certifyVulnSpec.ScannerVersion,
			Origin:         certifyVulnSpec.Origin,
			Collector:      certifyVulnSpec.Collector,
		}
		return r.Backend.CertifyVuln(ctx, &lowercaseCertifyVulnFilter)
	} else {
		return r.Backend.CertifyVuln(ctx, &certifyVulnSpec)
	}
}
