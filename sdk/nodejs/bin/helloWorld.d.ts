import * as pulumi from "@pulumi/pulumi";
export declare class HelloWorld extends pulumi.CustomResource {
    /**
     * Get an existing HelloWorld resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HelloWorld;
    /**
     * Returns true if the given object is an instance of HelloWorld.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is HelloWorld;
    readonly loud: pulumi.Output<boolean | undefined>;
    readonly message: pulumi.Output<string>;
    readonly name: pulumi.Output<string>;
    /**
     * Create a HelloWorld resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HelloWorldArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a HelloWorld resource.
 */
export interface HelloWorldArgs {
    loud?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
}
