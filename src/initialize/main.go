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
	var modelType, directory, name, copyright = retrieveArguments()
	var model = createModel(modelType, name, copyright)
	saveModel(directory, model)
}

func retrieveArguments() (
	modelType string,
	directory string,
	name string,
	copyright string,
) {
	if len(osx.Args) < 6 {
		fmt.Println(
			"Usage: initialize (class | generic) (type | structure) <directory> <name> <copyright>",
		)
		osx.Exit(1)
	}
	modelType = osx.Args[1] + " " + osx.Args[2]
	directory = osx.Args[3]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	name = osx.Args[4]
	copyright = osx.Args[5]
	return modelType, directory, name, copyright
}

func createModel(
	modelType string,
	name string,
	copyright string,
) (model mod.ModelLike) {
	var generator = mod.Generator()
	switch modelType {
	case "class type":
		model = generator.CreateClassType(name, copyright)
	case "generic type":
		model = generator.CreateGenericType(name, copyright)
	case "class structure":
		model = generator.CreateClassStructure(name, copyright)
	case "generic structure":
		model = generator.CreateGenericStructure(name, copyright)
	default:
		fmt.Printf(
			"Illegal model type: %v.",
			modelType,
		)
		osx.Exit(1)
	}
	return model
}

func saveModel(directory string, model mod.ModelLike) {
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
	var modelFile = directory + "Package.go"
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(model)
	var bytes = []byte(source)
	err = osx.WriteFile(modelFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
