import * as pulumi from "@pulumi/pulumi";
import { HelloWorld } from "../sdk/nodejs/bin";

const helloworld = new HelloWorld("helloworld", { name: "world!" });
export const greeting = helloworld.message;
