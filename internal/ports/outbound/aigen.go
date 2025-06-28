package outbound

import "context"

// AIGenerator defines the interface for AI-based content generation.
type AIGenerator interface {
	// GeneratePersona creates a character persona based on a given prompt.
	GeneratePersona(ctx context.Context, prompt string) (string, error)
	// GenerateEventMessage creates a message for a specific event in the persona's voice.
	GenerateEventMessage(ctx context.Context, persona string, event string) (string, error)
	// GenerateImage creates an image based on a given prompt.
	GenerateImage(ctx context.Context, prompt string) (string, error) // returns image URL
	// GenerateVideo creates a short video using Veo3 model based on persona and image.
	GenerateVideo(ctx context.Context, persona, imageURL string) (string, error) // returns video URL
	// GenerateVideoFromPrompt creates a short video using Veo3 model based on text prompt and image.
	GenerateVideoFromPrompt(ctx context.Context, prompt, imageURL string) (string, error) // returns video URL
}
