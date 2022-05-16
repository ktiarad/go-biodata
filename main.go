package main

import (
	"biodata/repositories"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	id := args[1]

	biodata := repositories.GetBiodataById(id)
	fmt.Printf("Id\t\t:%s\nNama\t\t:%s\nAlamat\t\t:%s\nPekerjaan\t:%s\nAlasan\t\t:%s\n", biodata.Id, biodata.Name, biodata.Address, biodata.Job, biodata.Reason)
}
