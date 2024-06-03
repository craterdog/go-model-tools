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

func main() {
	var isGeneric, directory, name, copyright = retrieveArguments()
	var model = createModel(isGeneric, name, copyright)
	saveModel(directory, model)
}

func retrieveArguments() (
	isGeneric bool,
	directory string,
	name string,
	copyright string,
) {
	if len(osx.Args) < 5 {
		fmt.Println(
			"Usage: initialize (model | generic) <directory> <name> <copyright>",
		)
		osx.Exit(1)
	}
	isGeneric = osx.Args[1] == "generic"
	directory = osx.Args[2]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	name = osx.Args[3]
	copyright = osx.Args[4]
	return isGeneric, directory, name, copyright
}

func createModel(
	isGeneric bool,
	name string,
	copyright string,
) (model mod.ModelLike) {
	var generator = mod.Generator()
	if isGeneric {
		model = generator.CreateGeneric(name, copyright)
	} else {
		model = generator.CreateModel(name, copyright)
	}
	return model
}

func saveModel(directory string, model mod.ModelLike) {
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
	var modelFile = directory + "Package.go"
	fmt.Printf(
		"    Creating %q now...\n",
		modelFile,
	)
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(model)
	var bytes = []byte(source)
	err = osx.WriteFile(modelFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
