package zendesk

import (
	"net/http"
	"testing"
)

// Note: For these tests to pass, you must uncomment the lines that set the subdomain
// and API credential and replace the parameters with your information.

func TestGetJiraLinks(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "jira_links.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()
	//client.SetSubdomain("your subdomain")
	//client.SetCredential(NewAPITokenCredential("your email", "your API token"))
	_, _, err := client.GetJiraLinks(ctx)
	if err != nil {
		t.Fatalf("Failed to get links: %s", err)
	}
}

func TestJiraLink(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "jira_link.json")
	client := newTestClient(mockAPI)
    defer mockAPI.Close()
	//client.SetSubdomain("your subdomain")
	//client.SetCredential(NewAPITokenCredential("your email", "your API token"))
	_, err := client.GetJiraLink(ctx, 24777562)
	if err != nil {
		t.Fatalf("Failed to get links: %s", err)
	}
}
