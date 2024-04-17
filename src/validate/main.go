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
	gcm "github.com/craterdog/go-model-framework/v3/gcmn"
	osx "os"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		fmt.Println("Usage: validate <package-file>")
		return
	}
	var packageFile = osx.Args[1]

	// Parse the package file.
	var bytes, err = osx.ReadFile(packageFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gcm.Parser().Make()
	var model = parser.ParseSource(source)

	// Validate the model.
	var validator = gcm.Validator().Make()
	validator.ValidateModel(model)
}
