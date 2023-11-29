package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	xi "github.com/rwrz/elevenlabs-go-unofficial"
)

var key *string

func init() {
	k := os.Getenv("XI_API_KEY")
	if k == "" {
		panic("XI_API_KEY is not set")
	}
	key = &k
}

func main() {
	filePtr := flag.String("file", "", "the mp3 file to save")
	textPtr := flag.String("text", "", "the text for the API call")
	flag.Parse()

	if *filePtr == "" || *textPtr == "" {
		flag.PrintDefaults()
		return
	}

	cl, err := xi.NewClientWithResponses("https://api.elevenlabs.io")
	if err != nil {
		panic(fmt.Errorf("error creating client %w", err))
	}
	getVoicesResp, err := cl.GetVoicesV1VoicesGetWithResponse(context.Background(), &xi.GetVoicesV1VoicesGetParams{
		XiApiKey: key,
	})
	if err != nil {
		panic(fmt.Errorf("error getting voices %w", err))
	}
	var voiceId string
	if len(getVoicesResp.JSON200.Voices) > 0 {
		voiceId = getVoicesResp.JSON200.Voices[0].VoiceId
	} else {
		panic("no voices found")
	}
	ttsResp, err := cl.TextToSpeechV1TextToSpeechVoiceIdStreamPostWithResponse(context.Background(), voiceId, &xi.TextToSpeechV1TextToSpeechVoiceIdStreamPostParams{
		XiApiKey: key,
	}, xi.BodyTextToSpeechV1TextToSpeechVoiceIdStreamPost{
		Text: *textPtr,
	})
	err = SaveMPEGToFile(ttsResp, *filePtr)
	if err != nil {
		panic(fmt.Errorf("error saving mp3 %w", err))
	}
}

func SaveMPEGToFile(response *xi.TextToSpeechV1TextToSpeechVoiceIdStreamPostResponse, filepath string) error {
	// Check if there was an error
	if response.JSON422 != nil {
		// Handle the error
		log.Printf("Server error: %v", response.JSON422)
		return fmt.Errorf("server error: %v", response.JSON422)
	}

	// Write the body to file
	err := ioutil.WriteFile(filepath, response.Body, 0644)
	if err != nil {
		log.Printf("Failed to write to file: %v", err)
		return err
	}

	return nil
}
