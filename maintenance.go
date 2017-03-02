package pagerduty

import (
	"fmt"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/nlopes/slack"
)

func getMaintenanceWindows(client *pagerduty.Client, limit uint) slack.Attachment {
	attachment := slack.Attachment{
		Title:     "PagerDuty Maintenance Windows",
		TitleLink: getDomainURL("/maintenance_windows"),
	}

	opts := pagerduty.ListMaintenanceWindowsOptions{
		APIListObject: pagerduty.APIListObject{
			Limit: limit,
		},
	}

	if resp, err := client.ListMaintenanceWindows(opts); err != nil {
		panic(err)
	} else {
		for _, mw := range resp.MaintenanceWindows {
			mwField := slack.AttachmentField{
				Title: fmt.Sprintf("%d", mw.SequenceNumber),
				Value: mw.Description,
				Short: false,
			}
			attachment.Fields = append(attachment.Fields, mwField)
		}
	}

	return attachment
}
