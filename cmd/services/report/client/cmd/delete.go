package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Deletes an report",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := reportcomm.NewReportSvcClient(conn)
			if _, err := cli.Delete(context.TODO(),
				&reportcomm.DeleteReportRequest{Id: id}); err != nil {
				log.Fatalf("error calling delete: %v", err)
			}

			log.Printf("deleted report %s", id)
		},
	}
)

func init() {
	cmdDelete.Flags().StringVarP(&id, "id", "i", "", "report's id")
}
