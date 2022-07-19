package main

import (
	"errors"
	"io/ioutil"
	"net/http"

	"fmt"
	"os"

	twitter "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/hashicorp/go-multierror"
	flag "github.com/spf13/pflag"
)

var (
	tweet, apiKey, apiKeySecret, accessToken, accessTokenSecret string
	dryRun                                                      bool
)

func main() {
	parseAndValidateInput()

	if dryRun {
		setOutput("successMessage", tweet)
		return
	}

	config := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	_, resp, err := client.Statuses.Update(tweet, nil)
	if err != nil {
		setOutput("errorMessage", err.Error())
		os.Exit(1)
	}

	if resp != nil && resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		setOutput("errorMessage", fmt.Sprintf("StatusCode: %v Response: %v", resp.Status, string(body)))
		os.Exit(1)
	}

	setOutput("successMessage", tweet)
}

func parseAndValidateInput() {
	flag.StringVar(&tweet, "tweet", "", "Tweet message body")
	flag.StringVar(&apiKey, "apiKey", "", "Twitter api key")
	flag.StringVar(&apiKeySecret, "apiKeySecret", "", "Twitter api secret key")
	flag.StringVar(&accessToken, "accessToken", "", "Twitter access token")
	flag.StringVar(&accessTokenSecret, "accessTokenSecret", "", "Twitter access token secret")
	flag.BoolVar(&dryRun, "dryRun", false, "if true or if env var DRY_RUN=true, then a tweet will not be sent")
	flag.Parse()

	if os.Getenv("DRY_RUN") == "true" {
		dryRun = true
	}

	var err error
	if tweet == "" {
		err = multierror.Append(err, errors.New("tweet message is required"))
	}

	if !dryRun {
		if apiKey == "" {
			err = multierror.Append(err, errors.New("apiKey is required"))
		}

		if apiKeySecret == "" {
			err = multierror.Append(err, errors.New("apiKeySecret is required"))
		}

		if accessToken == "" {
			err = multierror.Append(err, errors.New("accessToken is required"))
		}

		if accessTokenSecret == "" {
			err = multierror.Append(err, errors.New("accessTokenSecret is requried"))
		}
	}

	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func setOutput(key, value string) {
	fmt.Printf("::set-output name=%s::%s\n", key, value)
}
