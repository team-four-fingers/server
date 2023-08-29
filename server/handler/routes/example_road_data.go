package routes

import "encoding/json"

func LoadExampleRoadData() []*Road {
	var roads []*Road
	err := json.Unmarshal([]byte(exampleRoadJSON), &roads)
	if err != nil {
		panic(err)
	}
	return roads
}

var exampleRoadJSON = `[
            {
              "name": "",
              "distance": 10,
              "duration": 2,
              "traffic_speed": 16,
              "traffic_state": 0,
              "vertexes": [
                126.94645738299782,
                37.540648597491284,
                126.94653569519582,
                37.54071236917905
              ]
            },
            {
              "name": "큰우물로",
              "distance": 43,
              "duration": 7,
              "traffic_speed": 22,
              "traffic_state": 0,
              "vertexes": [
                126.94653569519582,
                37.54071236917905,
                126.94689011122018,
                37.54045422807421
              ]
            },
            {
              "name": "마포대로",
              "distance": 2353,
              "duration": 182,
              "traffic_speed": 54,
              "traffic_state": 4,
              "vertexes": [
                126.94689011122018,
                37.54045422807421,
                126.9462178533875,
                37.53997968834302,
                126.94533317532066,
                37.539323039149224,
                126.9450757454925,
                37.539122510164596,
                126.94489624519147,
                37.539012781266834,
                126.94471725331348,
                37.53886701491795,
                126.9444824545317,
                37.538666686888234,
                126.94422502924127,
                37.53846615609569,
                126.9434846691338,
                37.53800901281748,
                126.94248575621307,
                37.53743241413595,
                126.94065691518657,
                37.53633477018781,
                126.93056313229121,
                37.53005362446494,
                126.92930810284975,
                37.52920428221219,
                126.92618073935613,
                37.52716656044489
              ]
            },
            {
              "name": "국제금융로",
              "distance": 350,
              "duration": 82,
              "traffic_speed": 12,
              "traffic_state": 2,
              "vertexes": [
                126.92618073935613,
                37.52716656044489,
                126.92616663902785,
                37.526571743557135,
                126.9263506012221,
                37.52636617613021,
                126.92638505367019,
                37.5263304475975,
                126.92669538237965,
                37.52599087179203,
                126.92681017959238,
                37.52587477936874,
                126.92740781589528,
                37.5252224468054,
                126.92803886124038,
                37.52460645608779,
                126.92814156911192,
                37.524544315221185
              ]
            },
            {
              "name": "국제금융로",
              "distance": 285,
              "duration": 89,
              "traffic_speed": 23,
              "traffic_state": 3,
              "vertexes": [
                126.92814156911192,
                37.524544315221185,
                126.92835481641949,
                37.52466338614908,
                126.92830789400247,
                37.52478009607315,
                126.92704424577205,
                37.5261201879427,
                126.92668815371744,
                37.52649539050219,
                126.92660767870801,
                37.52658476344289
              ]
            },
            {
              "name": "여의대로",
              "distance": 166,
              "duration": 23,
              "traffic_speed": 32,
              "traffic_state": 3,
              "vertexes": [
                126.92660767870801,
                37.52658476344289,
                126.92658415097829,
                37.52664762264244,
                126.9265832472903,
                37.52671068747428,
                126.92658234359953,
                37.52677375230543,
                126.9266039349301,
                37.52684603203225,
                126.92775888985206,
                37.527568348258484,
                126.92781480595625,
                37.527613908101785
              ]
            },
            {
              "name": "여의대로6길",
              "distance": 94,
              "duration": 18,
              "traffic_speed": 28,
              "traffic_state": 0,
              "vertexes": [
                126.92781480595625,
                37.527613908101785,
                126.92784861355264,
                37.52762322546749,
                126.92790504552839,
                37.52763274821296,
                126.9279276698993,
                37.52763295359797,
                126.92799592981201,
                37.52760654193381,
                126.92855782979024,
                37.52708002546751
              ]
            },
            {
              "name": "",
              "distance": 125,
              "duration": 28,
              "traffic_speed": 16,
              "traffic_state": 0,
              "vertexes": [
                126.92855782979024,
                37.52708002546751,
                126.92744786436701,
                37.52637614646941
              ]
            }
          ]`
