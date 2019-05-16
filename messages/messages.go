package messages

import (
	"fmt"
	"github.com/autom8ter/memeblast/cats"
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

type Blast struct {
	TwilioAccount string            `validate:"required"`
	TwilioKey     string            `validate:"required"`
	Count         int               `validate:"required"`
	Messages      map[string]string `validate:"required"`
	From          []string          `validate:"required"`
	Targets       []string          `validate:"required"`
}

func NewBlast() *Blast {
	viper.SetDefault("count", 1)
	viper.SetDefault("messages", cats.Messages())
	return &Blast{
		TwilioAccount: viper.GetString("twilio_account"),
		TwilioKey:     viper.GetString("twilio_key"),
		Count:         viper.GetInt("count"),
		Messages:      viper.GetStringMapString("messages"),
		Targets:       viper.GetStringSlice("targets"),
		From:          viper.GetStringSlice("from"),
	}
}

func (b *Blast) Validate() error {
	return tool.WrapErr(tool.Validate(b), "place an memeblast.yaml config in $PWD\n")
}

func (b *Blast) Twilio() *gotwilio.Twilio {
	return gotwilio.NewTwilioClient(b.TwilioAccount, b.TwilioKey)
}

func (b *Blast) JSONString() string {
	return string(tool.MarshalJSON(b))
}

func (b *Blast) AppendMessages(msgs map[string]string) {
	if b.Messages == nil {
		b.Messages = make(map[string]string)
	}
	for k, v := range msgs {
		b.Messages[k] = v
	}
}

func (b *Blast) AppendFromTwilioNumbers(from ...string) {
	b.From = append(b.From, from...)
}

func (b *Blast) AppendTargetNumbers(targets ...string) {
	b.Targets = append(b.Targets, targets...)
}

func (b *Blast) Once(number string) error {
	for msg, img := range b.Messages {
		_, ex, err := b.Twilio().SendMMS(b.From[rand.Intn(len(b.From))], number, msg, img, "", "")
		fmt.Println(string(tool.MarshalJSON(ex)))
		if err != nil {
			return tool.WrapErr(err, string(tool.MarshalJSON(ex)))
		}
	}
	return nil
}

func (b *Blast) Attack() {
	for x := 0; x < b.Count; x++ {
		for _, number := range b.Targets {
			err := b.Once(number)
			if err != nil {
				fmt.Println("error", b.Once(number).Error())
			}
		}
	}
}
