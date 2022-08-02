package services

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"testing"
)

func buildZones(t *testing.T, ctrl *gomock.Controller) client.Api {
	mock := mocks.NewMockApi(ctrl)

	var zonesResp cloudflare.ZonesResponse
	if err := faker.FakeData(&zonesResp); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListZonesContext(
		gomock.Any(),
		gomock.Any(),
	).Return(
		zonesResp,
		nil,
	)

	return mock
}

func TestZones(t *testing.T) {
	client.CFMockTestHelper(t, Zones(), buildZones)
}
