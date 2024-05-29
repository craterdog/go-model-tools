<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Class Model Tools

### Overview
This project provides a set of Go based command-line tools that can be used to
parse, validate and format a class model defined in a `Package.go` file using Go
Class Model Notation™ (aka GCMN) using the Go Class Model Framework.  It can
also generate the corresponding concrete class files for each abstract class
defined in the package file.

### Quick Links
For more information on this project click on the following links:
 * [project documentation](https://github.com/craterdog/go-model-tools/wiki)
 * [grammar framework](https://github.com/craterdog/go-model-framework/wiki)

### Getting Started
To install the class model tools do the following from a terminal window:
```
$ git clone git@github.com:craterdog/go-model-tools.git
$ cd go-model-tools/
$ etc/build.sh
$ ls
LICENSE		bin		go.mod		src
README.md	etc		go.sum		tst

$ ls bin/
format		generate	initialize	validate	version

$ bin/version
bin/version: v4.6.0
```

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
