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
	if len(osx.Args) < 2 {
		fmt.Println("Usage: generate <package-directory>")
		return
	}
	var directory = osx.Args[1]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}

	// Parse the class model.
	var packageFile = directory + "Package.go"
	var bytes, err = osx.ReadFile(packageFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = mod.Parser()
	var model = parser.ParseSource(source)

	// Generate the class files.
	var generator = mod.Generator()
	var iterator = model.GetClasses().GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		var name = sts.ToLower(sts.TrimSuffix(
			class.GetDeclaration().GetIdentifier(),
			"ClassLike",
		))
		var source = generator.GenerateClass(model, name)
		var bytes = []byte(source)
		var filename = directory + name + ".go"
		var err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
}
