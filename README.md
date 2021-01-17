## Go-Boilerplate

An easier way to bootstrap web projects.
The philosophies in this boilerplate are not official standards defined by the Go core dev team.

## POSTMAN Collection Link
**https://www.getpostman.com/collections/45c2a6a6cc6fbaa100b0**

# swagger endpoint :
http://host/swagger/index.html

## Concepts Covered
The following concepts are currently covered in this boilerplate.
The rest will be added in time.
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
- [x] Remote KV Store
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
.yml is used as the configuration file as it is much readable and also supports comments.
The yml config bind to the model which then passed through the project.

### pkg
It would mainly contain library code. It could contain multiple packages some of which may depend on other packages. Eventually, some pkg will be used by /apis to get exposed.

### docker
It contains the scripts need to build the code. Containerization scripts would reside here.

### vendor
All dependencies would be here.

## Flow
The flow of the program would look like this.

```
 main.go --> config --> apis --> pkg
```
