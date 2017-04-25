package cmd

import (
	"os"

	"google.golang.org/grpc/metadata"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/dcwangmit01/goapi/app/logutil"

	pb "github.com/dcwangmit01/goapi/app/pb"
)

var keyvalRootCmd = &cobra.Command{
	Use:   "keyval",
	Short: "Client used to set Key/Value on gRPC service",
	Long: `
Create a key;
    goapi keyval create <key> <value>

Read a key:
    goapi keyval read <key>

Update a key:
    goapi keyval update <key> <value>

Delete a key:
    goapi keyval delete <key>
`,
}

var keyvalCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Key/Value on gRPC service",
	Run: func(cmd *cobra.Command, args []string) {
		grpcDialAndRunKeyVal(keyvalCreate)
	},
}

func keyvalCreate(client pb.KeyValClient, ctx context.Context) {
	msg, _ := client.KeyValCreate(ctx, &pb.KeyValMessage{os.Args[3], os.Args[4]})
	logutil.AddCtx(log.WithFields(log.Fields{
		"message": msg,
	})).Info("Sent RPC Request")
}

var keyvalReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read new Key/Value on gRPC service",
	Run: func(cmd *cobra.Command, args []string) {
		grpcDialAndRunKeyVal(keyvalRead)
	},
}

func keyvalRead(client pb.KeyValClient, ctx context.Context) {

	// add metadata to context
	md := metadata.Pairs(
		"myauth", "myauthtoken",
	)
	ctx = metadata.NewContext(ctx, md)

	msg, _ := client.KeyValRead(ctx, &pb.KeyValMessage{os.Args[3], ""})
	logutil.AddCtx(log.WithFields(log.Fields{
		"message": msg,
	})).Info("Sent RPC Request")
}

var keyvalUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update new Key/Value on gRPC service",
	Run: func(cmd *cobra.Command, args []string) {
		grpcDialAndRunKeyVal(keyvalUpdate)
	},
}

func keyvalUpdate(client pb.KeyValClient, ctx context.Context) {
	msg, _ := client.KeyValUpdate(ctx, &pb.KeyValMessage{os.Args[3], os.Args[4]})
	logutil.AddCtx(log.WithFields(log.Fields{
		"message": msg,
	})).Info("Sent RPC Request")
}

var keyvalDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete new Key/Value on gRPC service",
	Run: func(cmd *cobra.Command, args []string) {
		grpcDialAndRunKeyVal(keyvalDelete)
	},
}

func keyvalDelete(client pb.KeyValClient, ctx context.Context) {
	msg, _ := client.KeyValDelete(ctx, &pb.KeyValMessage{os.Args[3], ""})
	logutil.AddCtx(log.WithFields(log.Fields{
		"message": msg,
	})).Info("Sent RPC Request")
}

func init() {
	RootCmd.AddCommand(keyvalRootCmd)
	keyvalRootCmd.AddCommand(keyvalCreateCmd)
	keyvalRootCmd.AddCommand(keyvalReadCmd)
	keyvalRootCmd.AddCommand(keyvalUpdateCmd)
	keyvalRootCmd.AddCommand(keyvalDeleteCmd)
}
