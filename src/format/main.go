/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package main

import (
	fmt "fmt"
	mod "github.com/craterdog/go-model-framework/v3/gcmn"
	osx "os"
	sts "strings"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		fmt.Println("Usage: format <package-directory>")
		return
	}
	var directory = osx.Args[1]

	// Parse the package file.
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var packageFile = directory + "Package.go"
	var bytes, err = osx.ReadFile(packageFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = mod.Parser().Make()
	var model = parser.ParseSource(source)

	// Reformat the package file.
	var formatter = mod.Formatter().Make()
	source = formatter.FormatModel(model)
	bytes = []byte(source)
	err = osx.WriteFile(packageFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
