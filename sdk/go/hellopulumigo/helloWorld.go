// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hellopulumigo

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type HelloWorld struct {
	pulumi.CustomResourceState

	Loud    pulumi.BoolPtrOutput `pulumi:"loud"`
	Message pulumi.StringOutput  `pulumi:"message"`
	Name    pulumi.StringOutput  `pulumi:"name"`
}

// NewHelloWorld registers a new resource with the given unique name, arguments, and options.
func NewHelloWorld(ctx *pulumi.Context,
	name string, args *HelloWorldArgs, opts ...pulumi.ResourceOption) (*HelloWorld, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HelloWorld
	err := ctx.RegisterResource("hello-pulumi-go:index:HelloWorld", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHelloWorld gets an existing HelloWorld resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHelloWorld(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HelloWorldState, opts ...pulumi.ResourceOption) (*HelloWorld, error) {
	var resource HelloWorld
	err := ctx.ReadResource("hello-pulumi-go:index:HelloWorld", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HelloWorld resources.
type helloWorldState struct {
}

type HelloWorldState struct {
}

func (HelloWorldState) ElementType() reflect.Type {
	return reflect.TypeOf((*helloWorldState)(nil)).Elem()
}

type helloWorldArgs struct {
	Loud *bool  `pulumi:"loud"`
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a HelloWorld resource.
type HelloWorldArgs struct {
	Loud pulumi.BoolPtrInput
	Name pulumi.StringInput
}

func (HelloWorldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*helloWorldArgs)(nil)).Elem()
}

type HelloWorldInput interface {
	pulumi.Input

	ToHelloWorldOutput() HelloWorldOutput
	ToHelloWorldOutputWithContext(ctx context.Context) HelloWorldOutput
}

func (*HelloWorld) ElementType() reflect.Type {
	return reflect.TypeOf((**HelloWorld)(nil)).Elem()
}

func (i *HelloWorld) ToHelloWorldOutput() HelloWorldOutput {
	return i.ToHelloWorldOutputWithContext(context.Background())
}

func (i *HelloWorld) ToHelloWorldOutputWithContext(ctx context.Context) HelloWorldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelloWorldOutput)
}

type HelloWorldOutput struct{ *pulumi.OutputState }

func (HelloWorldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelloWorld)(nil)).Elem()
}

func (o HelloWorldOutput) ToHelloWorldOutput() HelloWorldOutput {
	return o
}

func (o HelloWorldOutput) ToHelloWorldOutputWithContext(ctx context.Context) HelloWorldOutput {
	return o
}

func (o HelloWorldOutput) Loud() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HelloWorld) pulumi.BoolPtrOutput { return v.Loud }).(pulumi.BoolPtrOutput)
}

func (o HelloWorldOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *HelloWorld) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

func (o HelloWorldOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HelloWorld) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HelloWorldInput)(nil)).Elem(), &HelloWorld{})
	pulumi.RegisterOutputType(HelloWorldOutput{})
}
