// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.hellopulumigo;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HelloWorldArgs extends com.pulumi.resources.ResourceArgs {

    public static final HelloWorldArgs Empty = new HelloWorldArgs();

    @Import(name="loud")
    private @Nullable Output<Boolean> loud;

    public Optional<Output<Boolean>> loud() {
        return Optional.ofNullable(this.loud);
    }

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    private HelloWorldArgs() {}

    private HelloWorldArgs(HelloWorldArgs $) {
        this.loud = $.loud;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HelloWorldArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HelloWorldArgs $;

        public Builder() {
            $ = new HelloWorldArgs();
        }

        public Builder(HelloWorldArgs defaults) {
            $ = new HelloWorldArgs(Objects.requireNonNull(defaults));
        }

        public Builder loud(@Nullable Output<Boolean> loud) {
            $.loud = loud;
            return this;
        }

        public Builder loud(Boolean loud) {
            return loud(Output.of(loud));
        }

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public HelloWorldArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}