package repositories

import (
	"biodata/models"
	"biodata/params"
)

func GetBiodataById(id string) models.Biodata {
	var biodata models.Biodata

	biodatas := GetBiodata()

	for _, val := range biodatas {
		if val.Id == id {
			biodata = val
			break
		}
	}
	return biodata
}

func GetBiodata() []models.Biodata {
	var biodatas []models.Biodata

	bio := params.AddNewBiodata("1", "Tiara", "Jalan Mawar", "Backend Engineer", "Mendalami Golang")
	biodatas = append(biodatas, bio)

	bio = params.AddNewBiodata("2", "Dewangga", "Jalan Melati", "DevOps", "Memahami Golang secara umum")
	biodatas = append(biodatas, bio)

	bio = params.AddNewBiodata("3", "Edi", "Jalan Layangan", "Backend Engineer", "Menguasai tech Golang")
	biodatas = append(biodatas, bio)

	bio = params.AddNewBiodata("4", "Erik", "Jalan Jalan", "Web Programmer", "Minat pindah ke backend")
	biodatas = append(biodatas, bio)

	bio = params.AddNewBiodata("5", "Mira", "Jalan Kenangan", "Guru TIK", "Tertarik dengan Golang")
	biodatas = append(biodatas, bio)

	return biodatas

}
