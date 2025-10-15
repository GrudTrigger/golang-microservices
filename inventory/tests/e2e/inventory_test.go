package integration

import (
	"context"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ = ginkgo.Describe("InvenrotyService", func() {
	var (
		ctx             context.Context
		cancel          context.CancelFunc
		inventoryClient inventoryV1.InventoryServiceClient
	)

	ginkgo.BeforeEach(func() {
		ctx, cancel = context.WithCancel(suiteCtx)

		conn, err := grpc.NewClient(env.App.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		gomega.Expect(err).ToNot(gomega.HaveOccurred(), "ожидали успешное подключение к grpc приложению")
		inventoryClient = inventoryV1.NewInventoryServiceClient(conn)
	})

	ginkgo.AfterEach(func() {
		err := env.ClearPartsCollection(ctx)
		gomega.Expect(err).ToNot(gomega.HaveOccurred(), "ожидали успешную очистку коллекции parts")
		cancel()
	})

	ginkgo.Describe("Get part for uuid", func() {
		var partUuid string
		ginkgo.BeforeEach(func() {
			var err error
			partUuid, err = env.InsertTestPartData(ctx)
			gomega.Expect(err).ToNot(gomega.HaveOccurred(), "ожидали успешную вставку тестовой запчасти в mongoDB")
		})

		ginkgo.It("должен успешно получить запчасть по uuid", func() {
			resp, err := inventoryClient.GetPart(ctx, &inventoryV1.GetPartRequest{Uuid: partUuid})
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			gomega.Expect(resp.GetPart()).ToNot(gomega.BeNil())
			gomega.Expect(resp.GetPart().Uuid).To(gomega.Equal(partUuid))
		})
	})

	ginkgo.Describe("Get parts list with filter name", func() {
		ginkgo.BeforeEach(func() {
			err := env.InsertTestSlicePartsData(ctx)
			gomega.Expect(err).ToNot(gomega.HaveOccurred(), "ожидали успешную вставку тестовой запчастей в mongoDB")
		})
		ginkgo.It("должен успешно получить запчасть с фильтром name:name-1", func() {
			resp, err := inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{
				Filter: &inventoryV1.PartsFilter{
					Names: []string{"name-1"},
				},
			})
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			gomega.Expect(resp.GetParts()).ToNot(gomega.BeNil())
			gomega.Expect(resp.GetParts()[0].Name).To(gomega.Equal("name-1"))
		})
	})
})
