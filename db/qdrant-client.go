package db

import (
	"context"
	"fmt"
	"log"

	pb "github.com/qdrant/go-client/qdrant"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Collections_client pb.CollectionsClient

func ClientToQdrant(ctx context.Context, v *viper.Viper) error {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", v.GetString("qdrant.ip"), v.GetInt("qdrant.port")),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	qdrantClient := pb.NewQdrantClient(conn)
	_, err = qdrantClient.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if err != nil {
		return err
	}
	Collections_client = pb.NewCollectionsClient(conn)
	num := v.GetUint64("qdrant.defaultSegmentNumber")
	// Create new collection:
	_, err = Collections_client.Create(ctx, &pb.CreateCollection{
		CollectionName: v.GetString("qdrant.collectionName"),
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     v.GetUint64("qdrant.vectorSize"),
				Distance: pb.Distance(v.GetInt32("qdrant.Distance_Dot")),
			},
		}},
		OptimizersConfig: &pb.OptimizersConfigDiff{
			DefaultSegmentNumber: &num,
		},
	})
	return nil
}
