# demo project to show problems with custom connection types using the go sdk

## what did I do

1. create the project according to the [go template from dai](https://github.com/digital-ai/release-integration-template-go)
2. realize that even in the template examples the connection property is
   not passed as function arg to the actual implementation (contrary to the other properties!)
   -> instead the boilerplate code (in main.go) tries to find a connection property with name "server",
   then creates a http client out of that and passes that as function args to the implementation.
   In my use case I want to use openapi client side code generation for accessing a REST api, which ships its own
   http client and thereby cannot work with the one that dai release provided.
   Apart from this, I would opt for user freedom to have complete control here.
   Another reason not to use the default client is that there is no option to provide a CA certificate for that.
3. So I tried to create a custom connection type called `portal.Server` (based on BasicAuthHttpConnection)
   and extended it only with the additional certificate property. Then I tried to use this connection in my Task as in
   input parameter.
4. The result is that the custom connection type property called `server` only contains empty strings, e.g. empty URL,
   emtpy user etc., contrary to what is configured in (my) dai release instance where I ran this task.

## problem summary

So far I can only see that this go sdk seems to lack proper support for custom connection types.
I could probably get my code working by switching from the `ci` type to other types of params, but so far I can
only say that managing the connections via dai release ui in the Connections gui is quite nice and I would like
to get this running cleanly.

Help in this regard would be much appreciated!

...now follows the original repo readme...

# lkww-release-version-deployed-integration

Dai release go sdk based plugin to get the currently deployed release (helm chart/helm release) version
(via lkww [portal](http://lkwbitbucket.lkw-walter.com/projects/MM/repos/portal/browse)).

This repo name follows the naming convention proposed by
the dai release go template - more information on how the base code for this repo was obtained
and how to work on this project in general can be found in the file
[README_orig.md](README_orig.md).
