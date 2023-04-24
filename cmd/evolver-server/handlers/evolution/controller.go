package evolution

import (
	"fmt"
	"tde/models/dto"
)

func Controller(request dto.EvolverService_Evolve_Request) (response dto.EvolverService_Evolve_Response) {
	fmt.Println(request.ArchiveID)
	return
}
