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
	mod "github.com/craterdog/go-model-framework/v2"
	osx "os"
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

	// Generate the class files.
	var generator = mod.Generator().Make()
	generator.CreateModel(directory, name, copyright)
}
