package pagerduty

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/domudall/doiici/plugins"
	"github.com/nlopes/slack"
)

type plugin struct{}

var (
	token     string
	subdomain string

	client *pagerduty.Client
)

func (p *plugin) Match(input string, params slack.PostMessageParameters) slack.PostMessageParameters {
	if input == "" {
		return p.Help(params)
	}

	inputParts := strings.Split(input, " ")

	if inputParts[0] == "list" {
		if len(inputParts) > 1 {
			if inputParts[1] == "users" {
				params.Attachments = append(params.Attachments, getUsers(client))
				return params
			}

			if inputParts[1] == "mw" {
				limit := 5
				if len(inputParts) > 2 {
					limit, _ = strconv.Atoi(inputParts[2])
					if limit == 0 {
						limit = 5
					}
				}

				params.Attachments = append(params.Attachments, getMaintenanceWindows(client, uint(limit)))
				return params
			}
		}
	}

	return p.Help(params)
}

func (p *plugin) GetName() string {
	return "pd"
}

func (p *plugin) Help(params slack.PostMessageParameters) slack.PostMessageParameters {
	return slack.PostMessageParameters{
		Text: "TODO",
	}
}

func getDomainURL(url string) string {
	if subdomain != "" {
		return fmt.Sprintf("https://%s.pagerduty.com%s", subdomain, url)
	}

	return ""
}

func init() {
	token = os.Getenv("PAGERDUTY_TOKEN")
	if token == "" {
		log.Fatal("token must be provided")
	}

	subdomain = os.Getenv("PAGERDUTY_SUBDOMAIN")

	client = pagerduty.NewClient(token)

	plugins.Add(&plugin{})
}
