package main

import (
	"github.com/InVisionApp/protoc-gen-tts/tts"
	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(tts.New()).
		RegisterPostProcessor(tts.NewFormatter()).
		Render()
}
