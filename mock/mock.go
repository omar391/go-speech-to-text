//Package mock build mock config and data for testing
package mock

import (
	"log"
	"strings"
	"stt-service/conf"

	"github.com/naoina/toml"
)

//SetupAppConfig setup mock config and clinics for testing
func SetupAppConfig() {
	appToml := `# listen addr of http api server
	addr = ":80"
	
	# urls of remote clinics 
	urls = ["http://127.0.0.1/data/dental-clinics.json",
	"http://127.0.0.1/data/dental-clinics.json"]
	
	[states]
	AL= "Alabama"
	AK= "Alaska"
	AZ= "Arizona"
	AR= "Arkansas"
	CA= "California"
	CO= "Colorado"
	CT= "Connecticut"
	DE= "Delaware"
	DC= "District of Columbia"
	FL= "Florida"
	GA= "Georgia"
	HI= "Hawaii"
	ID= "Idaho"
	IL= "Illinois"
	IN= "Indiana"
	IA= "Iowa"
	KS= "Kansas"
	KY= "Kentucky"
	LA= "Louisiana"
	ME= "Maine"
	MD= "Maryland"
	MA= "Massachusetts"
	MI= "Michigan"
	MN= "Minnesota"
	MS= "Mississippi"
	MO= "Missouri"
	MT= "Montana"
	NE= "Nebraska"
	NV= "Nevada"
	NH= "New Hampshire"
	NJ= "New Jersey"
	NM= "New Mexico"
	NY= "New York"
	NC= "North Carolina"
	ND= "North Dakota"
	OH= "Ohio"
	OK= "Oklahoma"
	OR= "Oregon"
	PA= "Pennsylvania"
	RI= "Rhode Island"
	SC= "South Carolina"
	SD= "South Dakota"
	TN= "Tennessee"
	TX= "Texas"
	UT= "Utah"
	VT= "Vermont"
	VA= "Virginia"
	WA= "Washington"
	WV= "West Virginia"
	WI= "Wisconsin"
	WY= "Wyoming"
	`

	if err := toml.NewDecoder(strings.NewReader(appToml)).Decode(&conf.Config); err != nil {
		log.Println("cfg: ", err)
	}

}

//MockClinicBuffer mock clinincs
var MockClinicBuffer = []byte(`[
	{
	   "clinicName":"Good Health Home",
	   "stateCode":"FL",
	   "opening":{
		  "from":"15:00",
		  "to":"20:00"
	   }
	},{
	  "name":"Good Health Home",
	  "stateName":"Alaska",
	  "availability":{
		 "from":"10:00",
		 "to":"19:30"
	  }
   },{
	   "clinicName":"National Veterinary Clinic",
	   "stateCode":"CA",
	   "opening":{
		  "from":"15:00",
		  "to":"22:30"
	   }
	},{
	   "clinicName":"Scratchpay Test Pet Medical Center",
	   "stateCode":"CA",
	   "opening":{
		  "from":"00:00",
		  "to":"24:00"
	   }
	},{
	  "name":"Cleveland Clinic",
	  "stateName":"California",
	  "availability":{
		 "from":"11:00",
		 "to":"22:00"
	  }
   },
   {
	  "name":"Mount Sinai Hospital",
	  "stateName":"California",
	  "availability":{
		 "from":"12:00",
		 "to":"22:00"
	  }
   },
   {
	  "name":"Tufts Medical Center",
	  "stateName":"Kansas",
	  "availability":{
		 "from":"10:00",
		 "to":"23:00"
	  }
   },
   {
	  "name":"UAB Hospital",
	  "stateName":"Alaska",
	  "availability":{
		 "from":"11:00",
		 "to":"22:00"
	  }
   },
   {
	  "name":"Swedish Medical Center",
	  "stateName":"Arizona",
	  "availability":{
		 "from":"07:00",
		 "to":"20:00"
	  }
   },
   {
	  "name":"Scratchpay Test Pet Medical Center",
	  "stateName":"California",
	  "availability":{
		 "from":"00:00",
		 "to":"24:00"
	  }
   }
]	
`)
