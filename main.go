// Copyright 2022 Nils Knieling
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"context"
	"fmt"
	"log"
	"flag"
	"cloud.google.com/go/pubsub"
)

func main() {
	program := os.Args[0]
	// https://cloud.google.com/sdk/gcloud/reference/pubsub/topics/publish
	project := flag.String("project", "", "Google Cloud project ID")
	topic   := flag.String("topic",   "", "Topic ID or fully qualified identifier for the topic")
	message := flag.String("message", "", "The body of the message to publish to the given topic name")

	flag.Usage = func() {
		fmt.Printf("Publish a message to the specified Google Cloud Pub/Sub topic.\n")
		fmt.Printf("\nSynopsis:\n")
		fmt.Printf("%v -project=PROJECT_ID -topic=TOPIC -message=MESSAGE \n", program)
		flag.PrintDefaults()  // prints default usage
		fmt.Printf("\nSet the environment variable GOOGLE_APPLICATION_CREDENTIALS to ")
		fmt.Printf("the path of the JSON file that contains your service account key.\n")
		fmt.Printf("\nAuthentication:\n")
		fmt.Printf("export GOOGLE_APPLICATION_CREDENTIALS=\"KEY_PATH\"\n\n")
	}
	flag.Parse() // after declaring flags we need to call it
	
	if *project == "" {
		log.Fatalf("[ERROR] Project ID not set!")
	}
	if *topic == "" {
		log.Fatalf("[ERROR] Topic ID not set!")
	}
	if *message == "" {
		log.Fatalf("[ERROR] No message provided!")
	}

	ctx := context.Background()

	// Creates a client
	client, err := pubsub.NewClient(ctx, *project)
	if err != nil {
		log.Fatalf("[ERROR] Failed to create client: %v", err)
	}
	defer client.Close()

	t := client.Topic(*topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(*message),
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("[ERROR] Failed to publish message: %v", err)
	}
	log.Printf("[OK] Published message with ID: %v\n", id)
}
