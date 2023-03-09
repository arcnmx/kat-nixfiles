// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package provider

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OpenSSHCertificate struct {
	pulumi.CustomResourceState

	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	Cakey     pulumi.StringOutput `pulumi:"cakey"`
	Content   pulumi.StringOutput `pulumi:"content"`
	Duration  pulumi.StringOutput `pulumi:"duration"`
	Hostname  pulumi.StringOutput `pulumi:"hostname"`
	Kind      pulumi.StringOutput `pulumi:"kind"`
	Userkey   pulumi.StringOutput `pulumi:"userkey"`
}

// NewOpenSSHCertificate registers a new resource with the given unique name, arguments, and options.
func NewOpenSSHCertificate(ctx *pulumi.Context,
	name string, args *OpenSSHCertificateArgs, opts ...pulumi.ResourceOption) (*OpenSSHCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Algorithm == nil {
		return nil, errors.New("invalid value for required argument 'Algorithm'")
	}
	if args.Cakey == nil {
		return nil, errors.New("invalid value for required argument 'Cakey'")
	}
	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Userkey == nil {
		return nil, errors.New("invalid value for required argument 'Userkey'")
	}
	var resource OpenSSHCertificate
	err := ctx.RegisterResource("opensshcertificate:provider:OpenSSHCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenSSHCertificate gets an existing OpenSSHCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenSSHCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenSSHCertificateState, opts ...pulumi.ResourceOption) (*OpenSSHCertificate, error) {
	var resource OpenSSHCertificate
	err := ctx.ReadResource("opensshcertificate:provider:OpenSSHCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenSSHCertificate resources.
type openSSHCertificateState struct {
}

type OpenSSHCertificateState struct {
}

func (OpenSSHCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*openSSHCertificateState)(nil)).Elem()
}

type openSSHCertificateArgs struct {
	Algorithm string `pulumi:"algorithm"`
	Cakey     string `pulumi:"cakey"`
	Duration  string `pulumi:"duration"`
	Hostname  string `pulumi:"hostname"`
	Kind      string `pulumi:"kind"`
	Userkey   string `pulumi:"userkey"`
}

// The set of arguments for constructing a OpenSSHCertificate resource.
type OpenSSHCertificateArgs struct {
	Algorithm pulumi.StringInput
	Cakey     pulumi.StringInput
	Duration  pulumi.StringInput
	Hostname  pulumi.StringInput
	Kind      pulumi.StringInput
	Userkey   pulumi.StringInput
}

func (OpenSSHCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openSSHCertificateArgs)(nil)).Elem()
}

type OpenSSHCertificateInput interface {
	pulumi.Input

	ToOpenSSHCertificateOutput() OpenSSHCertificateOutput
	ToOpenSSHCertificateOutputWithContext(ctx context.Context) OpenSSHCertificateOutput
}

func (*OpenSSHCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenSSHCertificate)(nil)).Elem()
}

func (i *OpenSSHCertificate) ToOpenSSHCertificateOutput() OpenSSHCertificateOutput {
	return i.ToOpenSSHCertificateOutputWithContext(context.Background())
}

func (i *OpenSSHCertificate) ToOpenSSHCertificateOutputWithContext(ctx context.Context) OpenSSHCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenSSHCertificateOutput)
}

// OpenSSHCertificateArrayInput is an input type that accepts OpenSSHCertificateArray and OpenSSHCertificateArrayOutput values.
// You can construct a concrete instance of `OpenSSHCertificateArrayInput` via:
//
//	OpenSSHCertificateArray{ OpenSSHCertificateArgs{...} }
type OpenSSHCertificateArrayInput interface {
	pulumi.Input

	ToOpenSSHCertificateArrayOutput() OpenSSHCertificateArrayOutput
	ToOpenSSHCertificateArrayOutputWithContext(context.Context) OpenSSHCertificateArrayOutput
}

type OpenSSHCertificateArray []OpenSSHCertificateInput

func (OpenSSHCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OpenSSHCertificate)(nil)).Elem()
}

func (i OpenSSHCertificateArray) ToOpenSSHCertificateArrayOutput() OpenSSHCertificateArrayOutput {
	return i.ToOpenSSHCertificateArrayOutputWithContext(context.Background())
}

func (i OpenSSHCertificateArray) ToOpenSSHCertificateArrayOutputWithContext(ctx context.Context) OpenSSHCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenSSHCertificateArrayOutput)
}

// OpenSSHCertificateMapInput is an input type that accepts OpenSSHCertificateMap and OpenSSHCertificateMapOutput values.
// You can construct a concrete instance of `OpenSSHCertificateMapInput` via:
//
//	OpenSSHCertificateMap{ "key": OpenSSHCertificateArgs{...} }
type OpenSSHCertificateMapInput interface {
	pulumi.Input

	ToOpenSSHCertificateMapOutput() OpenSSHCertificateMapOutput
	ToOpenSSHCertificateMapOutputWithContext(context.Context) OpenSSHCertificateMapOutput
}

type OpenSSHCertificateMap map[string]OpenSSHCertificateInput

func (OpenSSHCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OpenSSHCertificate)(nil)).Elem()
}

func (i OpenSSHCertificateMap) ToOpenSSHCertificateMapOutput() OpenSSHCertificateMapOutput {
	return i.ToOpenSSHCertificateMapOutputWithContext(context.Background())
}

func (i OpenSSHCertificateMap) ToOpenSSHCertificateMapOutputWithContext(ctx context.Context) OpenSSHCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenSSHCertificateMapOutput)
}

type OpenSSHCertificateOutput struct{ *pulumi.OutputState }

func (OpenSSHCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenSSHCertificate)(nil)).Elem()
}

func (o OpenSSHCertificateOutput) ToOpenSSHCertificateOutput() OpenSSHCertificateOutput {
	return o
}

func (o OpenSSHCertificateOutput) ToOpenSSHCertificateOutputWithContext(ctx context.Context) OpenSSHCertificateOutput {
	return o
}

func (o OpenSSHCertificateOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenSSHCertificate) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

func (o OpenSSHCertificateOutput) Cakey() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenSSHCertificate) pulumi.StringOutput { return v.Cakey }).(pulumi.StringOutput)
}

func (o OpenSSHCertificateOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenSSHCertificate) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

func (o OpenSSHCertificateOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenSSHCertificate) pulumi.StringOutput { return v.Duration }).(pulumi.StringOutput)
}

func (o OpenSSHCertificateOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenSSHCertificate) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

func (o OpenSSHCertificateOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenSSHCertificate) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o OpenSSHCertificateOutput) Userkey() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenSSHCertificate) pulumi.StringOutput { return v.Userkey }).(pulumi.StringOutput)
}

type OpenSSHCertificateArrayOutput struct{ *pulumi.OutputState }

func (OpenSSHCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OpenSSHCertificate)(nil)).Elem()
}

func (o OpenSSHCertificateArrayOutput) ToOpenSSHCertificateArrayOutput() OpenSSHCertificateArrayOutput {
	return o
}

func (o OpenSSHCertificateArrayOutput) ToOpenSSHCertificateArrayOutputWithContext(ctx context.Context) OpenSSHCertificateArrayOutput {
	return o
}

func (o OpenSSHCertificateArrayOutput) Index(i pulumi.IntInput) OpenSSHCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OpenSSHCertificate {
		return vs[0].([]*OpenSSHCertificate)[vs[1].(int)]
	}).(OpenSSHCertificateOutput)
}

type OpenSSHCertificateMapOutput struct{ *pulumi.OutputState }

func (OpenSSHCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OpenSSHCertificate)(nil)).Elem()
}

func (o OpenSSHCertificateMapOutput) ToOpenSSHCertificateMapOutput() OpenSSHCertificateMapOutput {
	return o
}

func (o OpenSSHCertificateMapOutput) ToOpenSSHCertificateMapOutputWithContext(ctx context.Context) OpenSSHCertificateMapOutput {
	return o
}

func (o OpenSSHCertificateMapOutput) MapIndex(k pulumi.StringInput) OpenSSHCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OpenSSHCertificate {
		return vs[0].(map[string]*OpenSSHCertificate)[vs[1].(string)]
	}).(OpenSSHCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OpenSSHCertificateInput)(nil)).Elem(), &OpenSSHCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*OpenSSHCertificateArrayInput)(nil)).Elem(), OpenSSHCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OpenSSHCertificateMapInput)(nil)).Elem(), OpenSSHCertificateMap{})
	pulumi.RegisterOutputType(OpenSSHCertificateOutput{})
	pulumi.RegisterOutputType(OpenSSHCertificateArrayOutput{})
	pulumi.RegisterOutputType(OpenSSHCertificateMapOutput{})
}
