package tts

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

// New tts plugin
func New() pgs.Module {
	return &tts{
		ModuleBase: &pgs.ModuleBase{},
	}
}

type tts struct {
	*pgs.ModuleBase
}

// Name is the identifier used to identify the module. This value is
// automatically attached to the BuildContext associated with the ModuleBase.
func (t *tts) Name() string { return "tts" }

type templateContext struct {
	pgs.File
	ProtoFileName string
}

func (t *tts) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		fname := f.Name().LowerSnakeCase().String()
		fname = fname[:len(fname)-6] // strip _proto from the value.
		c := templateContext{
			File:          f,
			ProtoFileName: fname,
		}
		t.AddGeneratorTemplateFile(fname+"_twirp_pb.js", template.Must(template.New("js").Parse(jsTemplate)), c)
		t.AddGeneratorTemplateFile(fname+"_twirp_pb.d.ts", template.Must(template.New("ts").Parse(tsTemplate)), c)
	}
	return t.Artifacts()
}
