package clouddeploy

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	deploy "cloud.google.com/go/deploy/apiv1"
)

func DeliveryPipelines() *schema.Table {
	return &schema.Table{
		Name:        "gcp_clouddeploy_delivery_pipelines",
		Description: `https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines#DeliveryPipeline`,
		Resolver:    fetchDeliveryPipelines,
		Multiplex:   client.ProjectMultiplexEnabledServices("clouddeploy.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.DeliveryPipeline{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.ProjectIDColumn(true),
		},
		Relations: []*schema.Table{
			Releases(),
		},
	}
}

func fetchDeliveryPipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListDeliveryPipelinesRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListDeliveryPipelines(ctx, req, c.CallOptions...)
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
