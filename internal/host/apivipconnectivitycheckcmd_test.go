package host

import (
	"context"

	"github.com/openshift/assisted-service/internal/common"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/models"
)

var _ = Describe("apivipconnectivitycheckcmd", func() {
	ctx := context.Background()
	var host models.Host
	var cluster common.Cluster
	var db *gorm.DB
	var apivipConnectivityCheckCmd *apivipConnectivityCheckCmd
	var id, clusterID strfmt.UUID
	var stepReply *models.Step
	var stepErr error
	dbName := "apivipconnectivitycheckcmd"

	BeforeEach(func() {
		db = common.PrepareTestDB(dbName)
		apivipConnectivityCheckCmd = NewAPIVIPConnectivityCheckCmd(getTestLog(), db, "quay.io/ocpmetal/assisted-installer-agent:latest", true)

		id = strfmt.UUID(uuid.New().String())
		clusterID = strfmt.UUID(uuid.New().String())
		host = getTestHostAddedToCluster(id, clusterID, models.HostStatusInsufficient)
		Expect(db.Create(&host).Error).ShouldNot(HaveOccurred())
		apiVipDNSName := "test.com"
		cluster = common.Cluster{Cluster: models.Cluster{ID: &clusterID, APIVipDNSName: &apiVipDNSName}}
		Expect(db.Create(&cluster).Error).ShouldNot(HaveOccurred())
	})

	It("get_step", func() {
		stepReply, stepErr = apivipConnectivityCheckCmd.GetStep(ctx, &host)
		Expect(stepReply).ShouldNot(BeNil())
		Expect(stepReply.Args[len(stepReply.Args)-1]).Should(Equal("{\"url\":\"http://test.com:22624/config/worker\",\"verify_cidr\":true}"))
		Expect(stepErr).ShouldNot(HaveOccurred())
	})

	It("get_step_unknown_cluster_id", func() {
		host.ClusterID = strfmt.UUID(uuid.New().String())
		stepReply, stepErr = apivipConnectivityCheckCmd.GetStep(ctx, &host)
		Expect(stepReply).To(BeNil())
		Expect(stepErr).Should(HaveOccurred())
	})

	AfterEach(func() {
		common.DeleteTestDB(db, dbName)
		stepReply = nil
		stepErr = nil
	})
})
