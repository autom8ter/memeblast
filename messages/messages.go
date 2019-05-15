package messages

import (
	"fmt"
	"github.com/autom8ter/objectify"
	"github.com/sfreiberg/gotwilio"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var tool = objectify.Default()

type Messages []string

type Images []string

type Blast struct {
	Messages []string
	Images   []string
}

func NewBlast(messages []string, images []string) *Blast {
	return &Blast{Messages: messages, Images: images}
}

func (b *Blast) Attack() {
	twilio := gotwilio.NewTwilioClient(viper.GetString("twilio_account"), viper.GetString("twilio_key"))
	from := viper.GetStringSlice("from")
	to := viper.GetStringSlice("to")

	for _, p := range to {
		_, ex, err := twilio.SendMMS(from[rand.Intn(len(from))], p, b.Messages[rand.Intn(len(b.Messages))], b.Images[rand.Intn(len(b.Images))], "", "")
		fmt.Println(string(tool.MarshalYAML(ex)))
		if err != nil {
			fmt.Println(tool.WrapErr(err, string(tool.MarshalJSON(ex))).Error())
		}
	}
}
