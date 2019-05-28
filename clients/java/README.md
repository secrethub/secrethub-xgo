# SecretHub Java Client<sup>**experimental**</sup>

> Before getting started, make sure you've built the SecretHub xgo client. See the root of this project.

To get started, first [install]() the SecretHub CLI, [sign up]() for a free developer account. 

> Make sure you replace `your-username` in the commands below with your own username.

Then create a repository named `javatest`:

```
secrethub repo init your-username/javatest
```

And create a secret:

```
echo "Hello World" | secrethub write your-username/javatest/hello
```

Because of the experimental nature of this codebase, we highly recommend you don't use your personal account credentials in the code and instead create a dedicated service account for this project.

```
export SECRETHUB_CREDENTIAL=$(secrethub service init your-username/java --permission read --descr "Testing out the SecretHub Java Client")
```

Now make sure you've installed [Gradle](https://gradle.org/) and execute the example code:

```
gradle run --args="your-username/javatest/hello"
```

It should print the secret you've written to SecretHub.

See `io.secrethub.example.Main` for the full example code.

## Experimental status

Note that this project is still very early stage and should **NOT** be considered anywhere near stable enough to use in production. 

A couple of TODOs are:

- [ ] Extract the example from the library and into a separate application.
- [ ] Ensure we free memory correctly. See https://golang.org/cmd/cgo/#hdr-Go_references_to_C
- [ ] Package the client with compiled C code for all required platforms.
- [ ] Return full client responses instead of primitive types.