package api

import (
	"context"

	"github.com/ozoncp/ocp-template-api/internal/repo"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozoncp/ocp-template-api/pkg/ocp-template-api"
)

var (
	totalSuccessCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_template_api_success_created_total",
		Help: "Total number of requests for templates successfully created",
	})
	totalSuccessUpdated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_template_api_success_updated_total",
		Help: "Total number of requests for templates successfully updated",
	})
	totalSuccessDeleted = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_template_api_success_deleted_total",
		Help: "Total number of requests for templates successfully deleted",
	})
)

type templateAPI struct {
	pb.UnimplementedOcpTemplateApiServiceServer
	repo repo.Repo
}

func NewTemplateAPI(r repo.Repo) pb.OcpTemplateApiServiceServer {
	return &templateAPI{repo: r}
}

func (o *templateAPI) DescribeTemplateV1(
	ctx context.Context,
	req *pb.DescribeTemplateV1Request,
) (*pb.DescribeTemplateV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeTemplateV1 - invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	template, err := o.repo.DescribeTemplate(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("DescribeTemplateV1 -- failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if template == nil {
		return nil, status.Error(codes.NotFound, "template not found")
	}

	log.Debug().Msg("DescribeTemplateV1 - success")

	return &pb.DescribeTemplateV1Response{
		Value: &pb.Template{
			Id:  template.ID,
			Foo: template.Foo,
		},
	}, nil
}
