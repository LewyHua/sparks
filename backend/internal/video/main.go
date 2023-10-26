package video

import "google.golang.org/grpc"

func main() {
	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{}
	grpc.NewServer(opts...)
}
