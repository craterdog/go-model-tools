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
	mod "github.com/craterdog/go-model-framework/v4"
	osx "os"
	sts "strings"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 4 {
		fmt.Println("Usage: initialize <package-directory> <package-name> <copyright>")
		return
	}
	var directory = osx.Args[1]
	var name = osx.Args[2]
	var copyright = osx.Args[3]

	// Create a new model.
	var source = mod.CreateModel(name, copyright)

	// Create a new directory for the package.
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}

	// Save the new model template.
	var packageFile = directory + "Package.go"
	fmt.Printf(
		"The package file %q for the model does not yet exist.\n\tCreating a template for it...\n",
		packageFile,
	)
	var bytes = []byte(source)
	err = osx.WriteFile(packageFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
