# SecretHub Cross Language Golang Client<sup>**experimental**</sup>

`secrethub-xgo` wraps the `secrethub-go` client with `cgo` exported functions so it can be called from other languages, e.g. C, C#, Python, Ruby, NodeJS, and Java. 

At the moment, we provide the library for C# usage. 

In order to make use of the package, you will need to have the following installed:
 - [.NET Core](https://docs.microsoft.com/en-gb/dotnet/core/install/)
 - [Golang](https://golang.org/doc/install)

## Create and use the package

First, you have to make the NuGet Package and use it in your project. To do so, Here are the steps:
1. execute `make nupkg` from the Makefile
2. Go to your .NET project and run the following command: `dotnet add package SecretHub -s <path_to_your_secrethub-xgo_repo>`.
3. Since our library uses Json.Net, you also need to add this dependency to your project. To do so, run the following command: `dotnet add package Newtonsoft.Json`.

## How to call library functions

- Read
`SecretHub.SecretVersion secret = SecretHub.Client.Read("path_to_secret");`

- ReadString
`string secret = SecretHub.Client.ReadString("path_to_secret");`

- Exists
`bool isSecret = SecretHub.Client.Exists("path_to_secret");`

- Write
`SecretHub.Client.Write("path_to_secret", "secret_value");`

- Remove
`SecretHub.Client.Remove("path_to_secret");`

- Resolve
`string secret = SecretHub.Client.Resolve("reference");`

- ResolveEnv
`Doctionary<string, string> resolvedEnv = SecretHub.Client.ResolveEnv(dictionaryToResolve);`

## Example

An example project is found in `azure-example`. There you can see the library in action.

## Resources

For example code and more reading on the specific tech behind this cross language magic, see these resources:

- https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf
- https://github.com/vladimirvivien/go-cshared-examples/
- https://github.com/draffensperger/go-interlang
- https://golang.org/cmd/cgo/

## Getting Help

Come chat with us on [Discord](https://discord.gg/EQcE87s) or email us at [support@secrethub.io](mailto:support@secrethub.io)