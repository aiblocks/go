package aiblockstoml

import "log"

// ExampleGetTOML gets the aiblocks.toml file for coins.asia
func ExampleClient_GetAiBlocksToml() {
	_, err := DefaultClient.GetAiBlocksToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
