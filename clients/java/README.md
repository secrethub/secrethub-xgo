# SecretHub Java Client<sup>**experimental**</sup>

> Before getting started, make sure you've built the `secrethub-xgo` client. See the `README.md` at the root of this project.

To get started, first [install](https://secrethub.io/docs/getting-started/install/) the SecretHub CLI, [sign up](https://secrethub.io/docs/getting-started/signup/) for a free developer account. 

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
export SECRETHUB_CREDENTIAL=$(secrethub service init your-username/java --permission read --desc "Testing out the SecretHub Java Client")
```

Now make sure you've installed [Gradle](https://gradle.org/) and execute the example code:

```
gradle run --args="your-username/javatest/hello"
```

It should print the secret you've written to SecretHub.

See `io.secrethub.example.Main` for the full example code.

## Experimental status

Note that this project is still very early stage and should **NOT** be considered anywhere near stable enough to use in production. 