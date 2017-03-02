package pagerduty

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/nlopes/slack"
)

func getUsers(client *pagerduty.Client) slack.Attachment {
	attachment := slack.Attachment{
		Title:     "PagerDuty Users",
		TitleLink: getDomainURL("/users"),
	}

	var opts pagerduty.ListUsersOptions

	if resp, err := client.ListUsers(opts); err != nil {
		panic(err)
	} else {
		for _, user := range resp.Users {
			userField := slack.AttachmentField{
				Title: user.Name,
				Value: user.Email,
				Short: false,
			}
			attachment.Fields = append(attachment.Fields, userField)
		}
	}

	return attachment
}
