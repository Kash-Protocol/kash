package client

import (
	"context"
	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the kashwalletd server, and returns the client instance
func Connect(address string) (pb.KashwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("kashwallet daemon is not running, start it with `kashwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewKashwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
