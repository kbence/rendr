package rpc

import (
	"context"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kbence/rendr/db"
	uuid "github.com/satori/go.uuid"
)

func getJobCollection() *mgo.Collection {
	return db.GetDatabase().C("job")
}

type jobService struct{}

func NewJobService() *jobService {
	return &jobService{}
}

func (s *jobService) Create(ctx context.Context, req *JobCreateRequest) (*JobCreateResponse, error) {
	job := db.JobModel{
		ID:     uuid.NewV1().String(),
		Name:   req.GetName(),
		Status: Status_QUEUED.String(),
	}
	getJobCollection().Insert(job)
	return &JobCreateResponse{}, nil
}

func (s *jobService) List(ctx context.Context, req *JobListRequest) (*JobListResponse, error) {
	query := bson.M{}

	if states := req.GetStates(); states != nil {
		var stateStrings []string
		for _, state := range states {
			stateStrings = append(stateStrings, state.String())
		}
		query["status"] = bson.M{"$in": stateStrings}
	}

	var jobsFound []db.JobModel
	err := getJobCollection().Find(query).All(&jobsFound)

	if err != nil {
		return nil, err
	}

	jobDefinitions := []*JobDefinition{}
	for _, job := range jobsFound {
		jobDefinitions = append(jobDefinitions, &JobDefinition{
			Id:     job.ID,
			Name:   job.Name,
			Status: job.Status,
		})
	}

	return &JobListResponse{
		Jobs: jobDefinitions,
	}, nil
}
