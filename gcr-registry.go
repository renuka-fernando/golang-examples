package main

import (
	"context"
	"fmt"
	"os"

	containeranalysis "cloud.google.com/go/containeranalysis/apiv1"
	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"
)

// getOccurrence retrieves and prints a specified Occurrence from the server.
func getOccurrence(occurrenceID, projectID string) (*grafeaspb.Occurrence, error) {
	// occurrenceID := path.Base(occurrence.Name)
	ctx := context.Background()
	client, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %v", err)
	}
	defer client.Close()

	req := &grafeaspb.GetOccurrenceRequest{
		Name: fmt.Sprintf("projects/%s/occurrences/%s", projectID, occurrenceID),
	}
	occ, err := client.GetGrafeasClient().GetOccurrence(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("client.GetOccurrence: %v", err)
	}
	return occ, nil
}

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/renuka/Desktop/apim-kube-4f6ee7acd8c3.json")
	fmt.Println(getOccurrence("", "apim-kube"))
}
