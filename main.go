package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/InVisionApp/protoc-gen-tts/tts"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(tts.New()).
		RegisterPostProcessor(tts.NewFormatter()).
		Render()
}
