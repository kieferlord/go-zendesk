package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// JiraLink is struct for jira link payload
// https://developer.zendesk.com/api-reference/ticketing/jira/links 
type JiraLink struct {
	ID                 int64                  `json:"id"`
	CreatedAt          string                 `json:"created_at"`
	IssueID		   int64                  `json:"issue_id"`
	IssueKey	   string                 `json:"issue_key"`
	TicketID	   int64                  `json:"ticket_id"`
	SharedTickets      bool                   `json:"shared_tickets"`
	UpdatedAt	   string		  `json:"updated_at"`
}

// JiraLinkListOptions is options for GetJiraLinks
//
// ref: https://developer.zendesk.com/rest_api/docs/support/organizations#list-organizations
type JiraLinkListOptions struct {
	PageOptions
}

// JiraLinkAPI an interface containing all methods associated with Jira links 
type JiraLinkAPI interface {
	GetJiraLinks(ctx context.Context, opts *JiraLinkListOptions) ([]JiraLink, Page, error)
	CreateJiraLink(ctx context.Context, link JiraLink) (JiraLink, error)
	GetJiraLink(ctx context.Context, linkID int64) (JiraLink, error)
	UpdateJiraLink(ctx context.Context, linkID int64, link JiraLink) (JiraLink, error)
	DeleteJiraLink(ctx context.Context, linkID int64) error
}

// GetJiraLinks fetch link list
//
// ref: https://developer.zendesk.com/rest_api/docs/support/organizations#getting-organizations
func (z *Client) GetJiraLinks(ctx context.Context, opts *JiraLinkListOptions) ([]JiraLink, Page, error) {
	var data struct {
		JiraLinks []JiraLink `json:"jira_links"`
		Page
	}

	if opts == nil {
		return []JiraLink{}, Page{}, &OptionsError{opts}
	}

	u, err := addOptions("/jira/links.json", opts)
	if err != nil {
		return []JiraLink{}, Page{}, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return []JiraLink{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []JiraLink{}, Page{}, err
	}

	return data.JiraLinks, data.Page, nil
}

// CreateJiraLink creates new organization
// https://developer.zendesk.com/rest_api/docs/support/organizations#create-organization
func (z *Client) CreateJiraLink(ctx context.Context, org JiraLink) (JiraLink, error) {
	var data, result struct {
		JiraLink JiraLink `json:"organization"`
	}

	data.JiraLink = org

	body, err := z.post(ctx, "/organizations.json", data)
	if err != nil {
		return JiraLink{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return JiraLink{}, err
	}

	return result.JiraLink, nil
}

// GetJiraLink gets a specified link 
// ref: https://developer.zendesk.com/rest_api/docs/support/organizations#show-organization
func (z *Client) GetJiraLink(ctx context.Context, linkID int64) (JiraLink, error) {
	var result struct {
		JiraLink JiraLink `json:"jira_link"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/jira/links/%d.json", linkID))

	if err != nil {
		return JiraLink{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return JiraLink{}, err
	}

	return result.JiraLink, err
}

// UpdateJiraLink updates a organization with the specified organization
// ref: https://developer.zendesk.com/rest_api/docs/support/organizations#update-organization
func (z *Client) UpdateJiraLink(ctx context.Context, orgID int64, org JiraLink) (JiraLink, error) {
	var result, data struct {
		JiraLink JiraLink `json:"organization"`
	}

	data.JiraLink = org

	body, err := z.put(ctx, fmt.Sprintf("/organizations/%d.json", orgID), data)

	if err != nil {
		return JiraLink{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return JiraLink{}, err
	}

	return result.JiraLink, err
}

// DeleteJiraLink deletes the specified organization
// ref: https://developer.zendesk.com/rest_api/docs/support/organizations#delete-organization
func (z *Client) DeleteJiraLink(ctx context.Context, orgID int64) error {
	err := z.delete(ctx, fmt.Sprintf("/organizations/%d.json", orgID))

	if err != nil {
		return err
	}

	return nil
}
