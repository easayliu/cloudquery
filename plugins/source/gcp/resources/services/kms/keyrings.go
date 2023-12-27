package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func KeyRings() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_keyrings",
		Description: `https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing`,
		Resolver:    fetchKeyrings,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudkms.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.KeyRing{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.ProjectIDColumn(true),
		},
		Relations: []*schema.Table{
			CryptoKeys(),
			ImportJobs(),
		},
	}
}
