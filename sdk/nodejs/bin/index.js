"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.Provider = exports.HelloWorld = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
exports.HelloWorld = null;
utilities.lazyLoad(exports, ["HelloWorld"], () => require("./helloWorld"));
exports.Provider = null;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));
const _module = {
    version: utilities.getVersion(),
    construct: (name, type, urn) => {
        switch (type) {
            case "hello-pulumi-go:index:HelloWorld":
                return new exports.HelloWorld(name, undefined, { urn });
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("hello-pulumi-go", "index", _module);
pulumi.runtime.registerResourcePackage("hello-pulumi-go", {
    version: utilities.getVersion(),
    constructProvider: (name, type, urn) => {
        if (type !== "pulumi:providers:hello-pulumi-go") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new exports.Provider(name, undefined, { urn });
    },
});
//# sourceMappingURL=index.js.map