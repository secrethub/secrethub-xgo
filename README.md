# SecretHub Cross Language Golang Client<sup>**experimental**</sup>

`secrethub-xgo` wraps the `secrethub-go` client with `cgo` exported functions so it can be called from other languages, e.g. C, Python, Ruby, NodeJS, and Java. 

For a proof of concept of how a client writting in another language can use `secrethub-xgo`, see the `clients/` folder. Right now it only contains a Java example, but more will be added in the future.

Make sure you build the `secrethub-xgo` client first before using the example client code. 

To build the `secrethub-xgo` c-shared object, run the `build.sh` script. It will place the build in a `linux` or `darwin` folder, depending on your OS. 

You'll need to have [Golang](https://golang.org/doc/install) installed. 


### Windows Support

Because we need `cgo` enabled for the compilation, Golang cannot cross compile for different operating systems as you might be used to. That's why I've only implemented it for Linux and macOS for now. 

Building for windows is *theoretically* possible by appliying the same concepts as written in the `build.sh` script. Difference would be to name the output `.dll` and to make the name of the output folder match your OS name. 

I haven't tried it yet though. 

In the future, we may want to check out [github.com/karalabe/xgo](https://github.com/karalabe/xgo) and see if we can use it to compile for all operating systems. 

## Experimental status

Note that this project is still very early stage and should **NOT** be considered anywhere near stable enough to use in production. 

- [ ] Ensure we free memory correctly. See https://golang.org/cmd/cgo/#hdr-Go_references_to_C
- [ ] Build and package the client with compiled C code for all required platforms, using e.g. [github.com/karalabe/xgo](https://github.com/karalabe/xgo).
- [ ] Return full client responses instead of primitive types.