package txt

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendBasicSms(msg, num string) {
	senderPhone := os.Getenv("M_PHONE")

	client := twilio.NewRestClient()

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(num)
	params.SetFrom(senderPhone)
	params.SetBody(msg)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		// fmt.Println(err.Error())
		fmt.Printf("Enter params for num & msg\nex:\n ./munch text --msg 'alexa show me this guy' --num '+1XXXYYYZZZ'\n\t for num put country code and 10 digits\n")
	}
	fmt.Println(resp)
}
