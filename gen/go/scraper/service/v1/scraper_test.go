package scraper

import (
	"context"
	"encoding/base64"
	"net"
	"testing"

	scraper_model "github.com/VampireAotD/anilibrary-grpc-gateway/gen/go/scraper/model/v1"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type scraperServerExample struct {
	UnimplementedScraperServiceServer
}

func (s scraperServerExample) Scrape(_ context.Context, request *ScrapeRequest) (*ScrapeResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "unsupported url")
	}

	return &ScrapeResponse{
		Anime: &scraper_model.Anime{
			Image:    base64.StdEncoding.EncodeToString([]byte("random")),
			Title:    "Наруто",
			Status:   "Вышел",
			Episodes: "220",
			Rating:   9.2,
		},
	}, nil
}

type ScraperServiceExampleSuite struct {
	suite.Suite

	listener *bufconn.Listener
	server   *grpc.Server
	client   ScraperServiceClient
	closer   func()
}

func TestScraperServiceExampleSuite(t *testing.T) {
	suite.Run(t, new(ScraperServiceExampleSuite))
}

func (suite *ScraperServiceExampleSuite) SetupSuite() {
	buffer := 1024 * 1024
	listener := bufconn.Listen(buffer)
	grpcServer := grpc.NewServer()

	RegisterScraperServiceServer(grpcServer, scraperServerExample{})

	suite.listener = listener
	suite.server = grpcServer
	suite.closer = func() {
		err := suite.listener.Close()
		suite.Require().NoError(err)
		suite.server.Stop()
	}

	go func() {
		err := suite.server.Serve(suite.listener)
		suite.Require().NoError(err)
	}()

	conn, err := grpc.Dial("", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return suite.listener.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	suite.Require().NoError(err)

	suite.client = NewScraperServiceClient(conn)
}

func (suite *ScraperServiceExampleSuite) TearDownSuite() {
	suite.closer()
}

func (suite *ScraperServiceExampleSuite) TestScrape() {
	require := suite.Require()

	suite.Run("Unsupported url", func() {
		response, err := suite.client.Scrape(context.Background(), &ScrapeRequest{Url: "unsupported.com"})
		require.Error(err)
		require.Nil(response)

		grpcError := status.Convert(err)
		require.Equal("unsupported url", grpcError.Message())
		require.Equal(codes.InvalidArgument, grpcError.Code())
	})

	suite.Run("Supported url", func() {
		expected := &ScrapeResponse{
			Anime: &scraper_model.Anime{
				Image:    base64.StdEncoding.EncodeToString([]byte("random")),
				Title:    "Наруто",
				Status:   "Вышел",
				Episodes: "220",
				Rating:   9.2,
			},
		}
		response, err := suite.client.Scrape(context.Background(), &ScrapeRequest{
			Url: "https://animego.org/anime/naruto-102",
		})
		require.NoError(err)
		require.NotNil(response)

		require.Equal(expected.Anime.Image, response.Anime.Image)
		_, err = base64.StdEncoding.DecodeString(response.Anime.Image)
		require.NoError(err)

		require.Equal(expected.Anime.Title, response.Anime.Title)
		require.Equal(expected.Anime.Status, response.Anime.Status)
		require.Equal(expected.Anime.Episodes, response.Anime.Episodes)
		require.Equal(expected.Anime.Rating, response.Anime.Rating)
	})
}
