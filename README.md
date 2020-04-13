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

- [X] Serialize data through one single type (e.g. `GoString`)
- [X] Implement a single contract for multiple languages with generated boilerplate.
- [ ] Review data serialization & abstract encoding away to avoid boilerplate.
- [ ] Improve/implement error handling
- [ ] Ensure we free memory correctly. See https://golang.org/cmd/cgo/#hdr-Go_references_to_C
- [ ] Organize repo structure, naming, and separate repos vs. monorepo.
- [ ] Dockerize build process
- [ ] Build for Linux, macOS & Windows
- [ ] Package for language native platforms, including the `.so` files etc.
- [ ] Make a beautiful README.md
- [ ] Implement Java Client:
    - [X] `Read`
    - [ ] `Write`
    - [ ] `Remove`
    - [ ] `Exists`
- [ ] Implement Python Client:
    - [X] `Read`
    - [ ] `Write`
    - [ ] `Remove`
    - [ ] `Exists`
- [ ] Implement Ruby Client:
    - [ ] `Read`
    - [ ] `Write`
    - [ ] `Remove`
    - [ ] `Exists`
- [ ] Implement JavaScript/Node Client:
    - [ ] `Read`
    - [ ] `Write`
    - [ ] `Remove`
    - [ ] `Exists`


## Resources

For example code and more reading on the specific tech behind this cross language magic, see these resources:

- https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf
- https://github.com/vladimirvivien/go-cshared-examples/
- https://github.com/draffensperger/go-interlang
- https://golang.org/cmd/cgo/

## Getting Help

Come chat with us on [Discord](https://discord.gg/EQcE87s) or email us at [support@secrethub.io](mailto:support@secrethub.io)


