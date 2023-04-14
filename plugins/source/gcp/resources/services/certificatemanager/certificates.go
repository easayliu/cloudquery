package certificatemanager

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	certificatemanager "cloud.google.com/go/certificatemanager/apiv1"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:        "gcp_certificatemanager_certificates",
		Description: `https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificates#Certificate`,
		Resolver:    fetchCertificates,
		Multiplex:   client.ProjectMultiplexEnabledServices("certificatemanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Certificate{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListCertificatesRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := certificatemanager.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListCertificates(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}
