package encoder

import (
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	member := &ESMember{
		DeviceType: "ios",
		RegSource:  "1,2,3",
		RegTime: &RegTime{
			RegTimeStart: time.Now().Unix(),
			RegTimeEnd:   time.Now().Add(time.Second * 5).Unix(),
		},
	}

	query, err := Marshal(member)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println(query)
}
