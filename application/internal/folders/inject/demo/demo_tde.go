//go:build tde

package demo

import (
	"tde/pkg/testing"
)

func TDE_WordReverse(t *testing.T) {
	// e.SetConfig(tde.Config{
	// 	MaxCompute:           100,
	// 	MaxMemory:            1000,
	// 	MaxSize:              1000,
	// 	MaxTime:              10,
	// 	ComputeToMemoryRatio: 3 / 2,
	// })

	testParameters := map[string]string{
		"Hello world":                    "dlrow olleH",
		"dlrow olleH":                    "Hello world",
		"The quick brown fox":            "xof nworb kciuq ehT",
		"Fusce a orci leo":               "oel icro a ecsuF",
		"Nulla a mollis est":             "tse sillom a alluN",
		"Etiam a semper nisl":            "lsin repmes a maitE",
		"Cras eget nulla diam":           "maid allun tege sarC",
		"Etiam in diam ligula":           "alugil maid ni maitE",
		"Aliquam erat volutpat":          "taptulov tare mauqilA",
		"Integer vel tincidunt erat":     "tare tnudicnit lev regetnI",
		"Phasellus ut commodo neque":     "euqen odommoc tu sullesahP",
		"Curabitur id elementum augue":   "eugua mutnemele di rutibaruC",
		"Duis tempus in dolor eu varius": "suirav ue rolod ni supmet siuD",
		"Nam rutrum et turpis eu tempor": "ropmet ue siprut te murtur maN",
		// "In hac habitasse platea dictumst":                                                           "tsmutcid aetalp essatibah cah nI",
		// "Phasellus venenatis placerat porta":                                                         "atrop tarecalp sitanenev sullesahP",
		// "Integer molestie fringilla pulvinar":                                                        "ranivlup allignirf eitselom regetnI",
		// "Praesent blandit lacinia vulputate":                                                         "etatupluv ainical tidnalb tnesearP",
		// "Vivamus ullamcorper vehicula mattis":                                                        "sittam alucihev reprocmallu sumaviV",
		// "Suspendisse convallis rhoncus dictum":                                                       "mutcid sucnohr sillavnoc essidnepsuS",
		// "Fusce laoreet vitae quam eget placerat":                                                     "tarecalp tege mauq eativ teeroal ecsuF",
		// "Donec feugiat dui ut venenatis lobortis":                                                    "sitrobol sitanenev tu iud taiguef cenoD",
		// "Suspendisse placerat consectetur ligula":                                                    "alugil rutetcesnoc tarecalp essidnepsuS",
		// "Nam efficitur magna non aliquet aliquet":                                                    "teuqila teuqila non angam ruticiffe maN",
		// "Sed condimentum mollis libero sed rutrum":                                                   "murtur des orebil sillom mutnemidnoc deS",
		// "Fusce mollis aliquam nibh pretium commodo":                                                  "odommoc muiterp hbin mauqila sillom ecsuF",
		// "Nam sed consequat est, ac bibendum libero":                                                  "orebil mudnebib ca ,tse tauqesnoc des maN",
		// "Praesent semper lobortis ex et sollicitudin":                                                "niduticillos te xe sitrobol repmes tnesearP",
		// "Quisque vel laoreet lorem, id gravida erat":                                                 "tare adivarg di ,merol teeroal lev euqsiuQ",
		// "Donec ultrices bibendum eros vel pellentesque":                                              "euqsetnellep lev sore mudnebib secirtlu cenoD",
		// "Etiam rutrum libero ac sapien vehicula commodo":                                             "odommoc alucihev neipas ca orebil murtur maitE",
		// "Morbi tempus felis quis risus dapibus aliquet":                                              "teuqila subipad susir siuq silef supmet ibroM",
		// "Ut a nisi ac sem viverra lobortis in vel purus":                                             "surup lev ni sitrobol arreviv mes ca isin a tU",
		// "Maecenas ac pretium velit, nec scelerisque diam":                                            "maid euqsirelecs cen ,tilev muiterp ca saneceaM",
		// "Vestibulum vel posuere velit, eget elementum ante":                                          "etna mutnemele tege ,tilev ereusop lev mulubitseV",
		// "In posuere nunc mi, quis dapibus risus placerat et":                                         "te tarecalp susir subipad siuq ,im cnun ereusop nI",
		// "Phasellus eu lectus a mauris placerat condimentum":                                          "mutnemidnoc tarecalp siruam a sutcel ue sullesahP",
		// "Cras efficitur dolor nisl, sed ultricies quam semper a":                                     "a repmes mauq seicirtlu des ,lsin rolod ruticiffe sarC",
		// "Maecenas consequat cursus odio, ornare bibendum neque":                                      "euqen mudnebib eranro ,oido susruc tauqesnoc saneceaM",
		// "Lorem ipsum dolor sit amet, consectetur adipiscing elit":                                    "tile gnicsipida rutetcesnoc ,tema tis rolod muspi meroL",
		// "Ut gravida hendrerit felis, id eleifend quam tempor in":                                     "ni ropmet mauq dnefiele di ,silef tirerdneh adivarg tU",
		// "Sed eu leo nec massa consectetur pharetra vitae id urna":                                    "anru di eativ arterahp rutetcesnoc assam cen oel ue deS",
		// "Ut semper tellus lorem, eu tristique ipsum fringilla vel":                                   "lev allignirf muspi euqitsirt ue ,merol sullet repmes tU",
		// "Aenean feugiat aliquam orci, at ultrices arcu lobortis eu":                                  "ue sitrobol ucra secirtlu ta ,icro mauqila taiguef naeneA",
		// "Donec laoreet ornare lectus, quis rhoncus nulla viverra a":                                  "a arreviv allun sucnohr siuq ,sutcel eranro teeroal cenoD",
		// "Ut purus lectus, bibendum in mattis et, posuere non metus":                                  "sutem non ereusop ,te sittam ni mudnebib ,sutcel surup tU",
		// "Curabitur id massa ut ex facilisis eleifend eu nec tortor":                                  "rotrot cen ue dnefiele sisilicaf xe tu assam di rutibaruC",
		// "Aenean tempor tempor enim, id varius risus dictum sit amet":                                 "tema tis mutcid susir suirav di ,mine ropmet ropmet naeneA",
		// "Interdum et malesuada fames ac ante ipsum primis in faucibus":                               "subicuaf ni simirp muspi etna ca semaf adauselam te mudretnI",
		// "Curabitur sed magna rhoncus, fringilla mauris ut, sodales elit":                             "tile selados ,tu siruam allignirf ,sucnohr angam des rutibaruC",
		// "Maecenas malesuada euismod erat, id tempus eros vestibulum id":                              "di mulubitsev sore supmet di ,tare domsiue adauselam saneceaM",
		// "Praesent convallis eros in nisi mattis, eu posuere est egestas":                             "satsege tse ereusop ue ,sittam isin ni sore sillavnoc tnesearP",
		// "Morbi ornare ex quis elit semper, sit amet lacinia diam cursus":                             "susruc maid ainical tema tis ,repmes tile siuq xe eranro ibroM",
		// "Aliquam fringilla pretium nulla, rutrum vehicula erat ornare sed":                           "des eranro tare alucihev murtur ,allun muiterp allignirf mauqilA",
		// "Donec lectus nisl, rutrum in commodo vitae, tempor sit amet nibh":                           "hbin tema tis ropmet ,eativ odommoc ni murtur ,lsin sutcel cenoD",
		// "Etiam volutpat ligula eget arcu tempus, a facilisis dui eleifend":                           "dnefiele iud sisilicaf a ,supmet ucra tege alugil taptulov maitE",
		// "Integer lacus dolor, hendrerit nec faucibus at, rutrum id augue":                            "eugua di murtur ,ta subicuaf cen tirerdneh ,rolod sucal regetnI",
		// "Vestibulum eu ante dignissim, fermentum leo ac, ultricies tortor":                           "rotrot seicirtlu ,ca oel mutnemref ,missingid etna ue mulubitseV",
		// "Integer lectus elit, semper vel magna vel, iaculis lobortis ipsum":                          "muspi sitrobol silucai ,lev angam lev repmes ,tile sutcel regetnI",
		// "Nullam auctor nisi sit amet tortor luctus, at auctor nibh pellentesque":                     "euqsetnellep hbin rotcua ta ,sutcul rotrot tema tis isin rotcua malluN",
		// "Maecenas magna neque, accumsan sit amet dolor sed, viverra posuere velit":                   "tilev ereusop arreviv ,des rolod tema tis nasmucca ,euqen angam saneceaM",
		// "Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos": "soeanemih sotpecni rep ,artson aibunoc rep tneuqrot arotil da uqsoicos iticat tnetpa ssalC",
		// "Morbi volutpat, justo at sodales porta, diam justo scelerisque odio, et fermentum velit ipsum sit amet ipsum":          "muspi tema tis muspi tilev mutnemref te ,oido euqsirelecs otsuj maid ,atrop selados ta otsuj ,taptulov ibroM",
		// "Proin condimentum, quam et ultricies scelerisque, augue dolor convallis massa, eget maximus risus mi ac metus":         "sutem ca im susir sumixam tege ,assam sillavnoc rolod eugua ,euqsirelecs seicirtlu te mauq ,mutnemidnoc niorP",
		// "Maecenas pretium, enim vestibulum ullamcorper maximus, orci nisi posuere odio, in dapibus ante velit eu ligula":        "alugil ue tilev etna subipad ni ,oido ereusop isin icro ,sumixam reprocmallu mulubitsev mine ,muiterp saneceaM",
		// "In scelerisque, nulla pulvinar molestie tincidunt, ante urna pellentesque elit, bibendum accumsan enim felis eu purus": "surup ue silef mine nasmucca mudnebib ,tile euqsetnellep anru etna ,tnudicnit eitselom ranivlup allun ,euqsirelecs nI",
	}

	for input, want := range testParameters {
		output := WordReverse(input)
		t.AssertEqual(output, want)
	}
}
