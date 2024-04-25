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
	age "github.com/craterdog/go-model-framework/v4/agent"
	osx "os"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		fmt.Println("Usage: format <package-file>")
		return
	}
	var packageFile = osx.Args[1]

	// Parse the package file.
	var bytes, err = osx.ReadFile(packageFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = age.Parser().Make()
	var model = parser.ParseSource(source)

	// Validate the model.
	var validator = age.Validator().Make()
	validator.ValidateModel(model)

	// Reformat the package file.
	var formatter = age.Formatter().Make()
	source = formatter.FormatModel(model)
	bytes = []byte(source)
	err = osx.WriteFile(packageFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
