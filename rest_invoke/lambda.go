package rest_invoke

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func (obj *Lambda) RestInvoke() error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	var response Response

	client := lambda.New(sess, &aws.Config{Region: aws.String(obj.Region)})

	payload, err := json.Marshal(&obj.Payload)
	if err != nil {
		return err
	}

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String(obj.Function), Payload: payload})
	if err != nil {
		return err
	}

	err = json.Unmarshal(result.Payload, &response)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("Error getting items, StatusCode: %d\n", response.StatusCode)
	}

	err = json.Unmarshal([]byte(response.Body), &obj.Object)
	if err != nil {
		return err
	}

	return nil
}
