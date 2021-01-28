## Go-Boilerplate

An easy way to build projects while maintaining a flexible structure

## POSTMAN Collection Link
**https://www.getpostman.com/collections/45c2a6a6cc6fbaa100b0**

## swagger endpoint
http://host/swagger/index.html

## Concepts Covered
The following concepts are currently covered in this boilerplate.
- [x] Directory structure
- [x] Dependency injection
- [x] Request interceptors
- [x] Vendoring
- [x] Commenting
- [x] Containerizing
- [x] Error
- [x] Logging
- [x] Documenting
- [x] Apm
- [x] Remote KV store
- [ ] Database & Orms
- [x] Unit Testing
- [x] Unit Test Framework
- [x] Makefile & build scripts
- [ ] Integration testing


## Directory structure

### apis
It contains controllers for the pkgs that need to be exposed. Each protocol can have different implementations. Each protocol's implementation can reside in its respective directories.

### config
It contains the configuration model. This model will be used in the project.

### initiate
It would contain code to start up the project. All dependencies would also be created here and then passed to the respective packages.

### pkg
It would mainly contain library code. It could contain multiple packages some of which may depend on other packages. Eventually, some pkg will be used by apis to get exposed.

### internal
contains internal business logic , which is used by the api layers and respectively expose required endpoints

### docker
It contains the scripts need to build the code. Containerization scripts would reside here.

### vendor
All dependencies would be here.

## Flow
The flow of the program would look like this.

```js
 main.go --> initiate --> config --> apis --> pkg
```
