# SecretHub-XGO Cross Language Golang Client (C#)

SecretHub - XGO wraps the `secrethub-go` client with `cgo` exported functions so it can be called form other languages, e.g. C, C#, Python, Ruby, NodeJS, and Java. To generate the code that will then be wrapped in the library used by a certian programming language, we use `swig`.

## Prerequisites

To make use of the package, you will need to have the following installed:
 - [.NET Core](https://docs.microsoft.com/en-gb/dotnet/core/install/)
 - [Golang](https://golang.org/doc/install)

## Usage

To install SecretHub package from NuGet Gallery, run the following command in your project's directory: 
```bash
dotnet add package SecretHub
```
or you can go to the project's `.csproj` file and add the following line:
```xml
<PackageReference Include="SecretHub" Version="0.1.0" />
```

## Building from source 
1. Execute `make nupkg` from the Makefile
2. Go to your .NET project direcotry and run the following command: `dotnet add package SecretHub -s <path_to_your_secrethub-xgo_repo>`.

## How to call library functions
Before doing any calls to the library, you need to create you SecretHub client. This is done in the following way:
```csharp
var client = new SecretHub.Client();
```

After you have your client, you can perform on of the following functions:

### `Read(string path)`
Retrieve a secret, including all its metadata:
```csharp
SecretHub.SecretVersion secret = client.Read("path/to/secret");
Console.WriteLine("The secret value is " + secret.Data);
```
`SecretHub.SecretVersion` object represents a version of a secret with sensitive data.

### `ReadString(string path)`
Retrieve a secret as a string:
 ```csharp
 string secret = client.Read("path/to/secret");
 Console.WriteLine("The secret value is " + secret);
 ```

### `Exists(string path)`
Check if a secret exists at `path`:
```csharp
bool isSecret = client.Exists("path/to/secret");
```

### `Write(string path, string secret)`
Write a secret containing the contents of `secret` at `path`:
```csharp
client.Write("path/to/secret", "secret_value");
```

### `Remove(string path)`
Delete the secret found at `path`, if it exists:
```csharp
client.Remove("path/to/secret");
if (!client.Exists("path/to/secret"))
{
    Console.WriteLine("Secret deleted successfully");
}
```

### `Resolve(string ref)`
Fetch the value of a secret from SecretHub, when the `ref` has the format `secrethub://<path>`, otherwise it returns `ref` unchanged:
```csharp
string resolvedRef = client.Resolve("secrethub://path/to/secret");
Console.WriteLine("The secret value got from reference is " + resolvedRef);
```

### `ResolveEnv()`
Take a map of system's environment variables and replaces the values of those which store references of secrets in SecretHub (`secrethub://<path>`) with the value of the respective secret. The other entries in the map remain untouched.
```csharp
Dictionary<string, string> resolvedEnv = client.ResolveEnv();
// resolvedEnv = Dictionary<string, string> {
// 		"REFERENCE_TO_SECRET_ENV_VAR": "some-secret-value",
//		"NOT_A_REFERENCE_ENV_VAR": "some-other-value"
// }
```

### Exceptions
In order to catch any exceptions thrown by the library (thus avoiding your program to crash), you have to put the lines that use the package in a `try-catch` block.
```csharp
try 
{
	string secret = client.Read("path/to/secret");
} 
catch(Exception ex)
{
	Console.WriteLine(ex.Message);
}
```

## Resources

For example code and more reading on the specific tech behind this cross language magic, see these resources:

- https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf
- https://github.com/vladimirvivien/go-cshared-examples/
- https://github.com/draffensperger/go-interlang
- https://golang.org/cmd/cgo/
- http://www.swig.org/

## Getting Help

Come chat with us on [Discord](https://discord.gg/EQcE87s) or email us at [support@secrethub.io](mailto:support@secrethub.io)
