package api

import (
	"github.com/kbence/rendr/db"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type JobModel struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name"`
}

var jobCollection = db.GetDatabase().C("job")

func NewUntypedJobModel() interface{} {
	return &JobModel{}
}

func GetJob(key string) (interface{}, error) {
	job := &JobModel{}

	err := jobCollection.Find(bson.M{"_id": key}).One(job)
	if err != nil {
		job = nil
	}

	return job, nil
}

func PostJob(content interface{}) (interface{}, error) {
	job := content.(*JobModel)
	job.ID = uuid.NewV1().String()
	err := jobCollection.Insert(job)

	if err != nil {
		job = nil
	}

	return job, err
}

func PutJob(key string, content interface{}) error {
	job := content.(*JobModel)
	job.ID = key
	return jobCollection.UpdateId(key, bson.M{"$set": job})
}

func DeleteJob(key string) error {
	return jobCollection.RemoveId(key)
}
