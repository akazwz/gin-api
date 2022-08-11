package nsfw

import (
	"context"
	"os"
	"time"

	pb "github.com/akazwz/gin-api/grpc/nsfw/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClassifyResult struct {
	Drawing float64
	Hentai  float64
	Neutral float64
	Porn    float64
	Sexy    float64
}

func ClassifyImage(image []byte) (err error, result ClassifyResult) {
	conn, err := grpc.Dial(os.Getenv("NSFW_SERVICE_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	c := pb.NewClassifyClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	reply, err := c.ClassifyImage(ctx, &pb.ClassifyRequest{Image: image})
	if err != nil {
		return
	}

	result = ClassifyResult{
		Drawing: reply.Drawing,
		Hentai:  reply.Hentai,
		Neutral: reply.Neutral,
		Porn:    reply.Porn,
		Sexy:    reply.Sexy,
	}

	return
}
