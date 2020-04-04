package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

var something string
var connection *bongo.Connection

type UserModel struct {
	bongo.DocumentBase `bson:",inline"`
	Username           string `json:"username" bson:"username"`
	Password           string `json:"password" bason:"password"`
	IsSuperUser        bool   `json:"is_superuser" bson:"is_supersuser"`
}

func Connect() {
	config := &bongo.Config{
		ConnectionString: "mongodb://mqtt_admin:kiotapp_mqtt_super@139.59.30.136:21007/mqtt",
		Database:         "mqtt",
	}
	var err error
	connection, err = bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}
}

/*
 */
func Finduser(username string, password string) {
	user := &UserModel{}
	err := connection.Collection("mqtt_user").FindOne(bson.M{}, user)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Found User : ", user.Username)
	}
}

func Anything() {
	// initialize progress container, with custom width
	p := mpb.New(mpb.WithWidth(64))

	total := 100
	name := "Single Bar:"
	// adding a single bar, which will inherit container's width
	bar := p.AddBar(int64(total),
		// override DefaultBarStyle, which is "[=>-]<+"
		mpb.BarStyle(" "),
		mpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	// simulating some work
	max := 100 * time.Millisecond
	for i := 0; i < total; i++ {
		time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
		bar.Increment()
	}
	// wait for our bar to complete and flush
	p.Wait()
}
