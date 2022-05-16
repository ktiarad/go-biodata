package params

import "biodata/models"

func AddNewBiodata(id, name, address, job, reason string) models.Biodata {
	return models.Biodata{
		Id:      id,
		Name:    name,
		Address: address,
		Job:     job,
		Reason:  reason,
	}
}
