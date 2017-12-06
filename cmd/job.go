package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/kbence/rendr/rpc"
	elastictable "github.com/rdsubhas/go-elastictable"
	"github.com/spf13/cobra"
)

func connect() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:5679", grpc.WithInsecure())

	if err != nil {
		log.Panic(err)
	}

	return conn
}

func NewJobCommand() *cobra.Command {
	jobCommand := &cobra.Command{
		Use: "job",
	}
	jobCommand.AddCommand(&cobra.Command{
		Use: "list",
		Run: JobList,
	})
	jobCommand.AddCommand(&cobra.Command{
		Use: "create",
		Run: JobCreate,
	})

	return jobCommand
}

func JobList(cmd *cobra.Command, args []string) {
	conn := connect()
	defer conn.Close()

	jobList, err := rpc.NewJobClient(conn).List(context.Background(), &rpc.JobListRequest{})

	if err != nil {
		log.Panic(err)
	}

	if jobs := jobList.GetJobs(); jobs != nil {
		table := elastictable.NewElasticTable([]string{"ID", "Name", "Status"})

		for _, job := range jobs {
			table.AddRow([]string{job.GetId(), job.GetName(), job.GetStatus()})
		}

		table.Render(os.Stdout)
	} else {
		fmt.Println("No matching jobs found!")
	}
}

func JobCreate(cmd *cobra.Command, args []string) {
	conn, err := grpc.Dial("localhost:5679")

	if err != nil {
		log.Panic(err)
	}

	rpc.NewJobClient(conn).Create(context.Background(), &rpc.JobCreateRequest{})
}
