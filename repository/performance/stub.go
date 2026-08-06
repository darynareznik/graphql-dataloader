package performance

import "graphql-dataloader/graph/model"

type performanceByAthleteID map[string]*model.Performance

var performanceStub = performanceByAthleteID{
	"12345": &model.Performance{
		AvgSpeed:   3.5,
		InTraining: true,
	},
	"25331": &model.Performance{
		AvgSpeed:   7,
		InTraining: false,
	},
	"2345": &model.Performance{
		AvgSpeed:   4.1,
		InTraining: true,
	},
	"7826": &model.Performance{
		AvgSpeed:   3.55,
		InTraining: false,
	},
	"1110": &model.Performance{
		AvgSpeed:   5,
		InTraining: true,
	},
}
