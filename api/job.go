package api

type JobModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUntypedJobModel() interface{} {
	return &JobModel{}
}

func GetJob(key string) (interface{}, error) {
	return nil, nil
}

func PostJob(key string, content interface{}) error {
	return nil
}

func PutJob(key string, content interface{}) error {
	return nil
}

func DeleteJob(key string) error {
	return nil
}
