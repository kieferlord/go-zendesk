package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
)

// JiraLink is struct for jira link payload
// https://developer.zendesk.com/api-reference/ticketing/jira/links 
type JiraLink struct {
	ID             int64                  `json:"id"`
	CreatedAt      string                 `json:"created_at"`
	IssueID		   json.Number            `json:"issue_id"`
	IssueKey	   string                 `json:"issue_key"`
	TicketID	   json.Number			  `json:"ticket_id"`
	SharedTickets  bool                   `json:"shared_tickets"`
	UpdatedAt	   string		  		  `json:"updated_at"`
}

// JiraLinkListOptions is options for GetJiraLinks
//
// ref: https://developer.zendesk.com/api-reference/ticketing/jira/links/ 
type JiraLinkListOptions struct {
	PageOptions
}

// JiraLinkAPI an interface containing all methods associated with Jira links 
type JiraLinkAPI interface {
	GetJiraLinks(ctx context.Context, opts *JiraLinkListOptions) ([]JiraLink, Page, error)
	GetJiraLink(ctx context.Context, linkID int64) (JiraLink, error)
}

// GetJiraLinks fetch link list
//
// ref: https://developer.zendesk.com/api-reference/ticketing/jira/links/ 
func (z *Client) GetJiraLinks(ctx context.Context) ([]JiraLink, Page, error) {
	var data struct {
		JiraLinks []JiraLink `json:"links"`
		Page
	}	

	body, err := z.get(ctx, "/jira/links")
	if err != nil {
		return []JiraLink{}, Page{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []JiraLink{}, Page{}, err
	}
	return data.JiraLinks, data.Page, nil
}


// GetJiraLink gets a specified link 
// ref: https://developer.zendesk.com/api-reference/ticketing/jira/links/ 
func (z *Client) GetJiraLink(ctx context.Context, linkID int64) (JiraLink, error) {
	var result struct {
		JiraLink JiraLink `json:"link"`
	}

    body, err := z.get(ctx, fmt.Sprintf("/jira/links/%d", linkID))
	if err != nil {
		return JiraLink{}, err
    }

	err = json.Unmarshal(body, &result)
    if err != nil {
		return JiraLink{}, err
	}
	return result.JiraLink, err
}


