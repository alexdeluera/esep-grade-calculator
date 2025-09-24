# ESEP Grade Calculator

This library is a simple utility to aid in the calculation of student grades.

Runtime requirements:
Golang 1.24+

If Golang is not yet installed, to install on Windows. Open Powershell as an administrator and run:
```
choco install golang
```
After installation has completed, verify that it was installed correctly and check the version by running:
```
go version
```
The version should be the same as the version listed in "grade-calculator/go.mod", or newer. 

Once golang is installed and the version has been checked, you may run the unit tests with:
```
cd grade-calculator/
go test ./...
```

Additionally, to run the unit tests with code coverage:
```
cd grade-calculator/
go test -coverprofile coverage.txt ./...
```
