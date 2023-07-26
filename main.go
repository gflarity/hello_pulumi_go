package main

import (
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"strings"
)

func main() {
	p.RunProvider("hello-pulumi-go", "1.0.0",
		// We tell the provider what resources it needs to support.
		// In this case, a single custom resource.
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[HelloWorld, HelloWorldArgs, HelloWorldState](),
			},
		}))
}

// HelloWorld is the controlling struct for the resource.
type HelloWorld struct{}

// HelloWorldArgs is the input struct, defining the inputs to necessary to creat the resource.
type HelloWorldArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but it's generally a
	// good idea.
	Name string `pulumi:"name"`
	// Fields marked `optional` are optional, so they should have a pointer
	// ahead of their type.
	Loud *bool `pulumi:"loud,optional"`
}

// HelloWorldState is the resource state, it describes the fields that exist on the created resource (outputs).
type HelloWorldState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	HelloWorldArgs
	// Here we define a required output called message.
	Message string `pulumi:"message"`
}

// Create is responsible for performing the external actions representative of creating this resource. All resources must implement Create at a minumum.
func (HelloWorld) Create(ctx p.Context, name string, input HelloWorldArgs, preview bool) (string, HelloWorldState, error) {
	state := HelloWorldState{HelloWorldArgs: input}
	if preview {
		return name, state, nil
	}
	state.Message = fmt.Sprintf("Hello, %s", input.Name)
	if input.Loud != nil && *input.Loud {
		state.Message = strings.ToUpper(state.Message)
	}
	return name, state, nil
}
