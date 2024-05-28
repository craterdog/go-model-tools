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
	var directory = retrieveArguments()
	var model = parseModel(directory)
	validateModel(model)
	generateClasses(directory, model)
}

func retrieveArguments() (directory string) {
	if len(osx.Args) < 2 {
		fmt.Println("Usage: generate <directory>")
		osx.Exit(1)
	}
	directory = osx.Args[1]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	return directory
}

func validateModel(model mod.ModelLike) {
	var validator = mod.Validator()
	validator.ValidateModel(model)
}

func parseModel(directory string) mod.ModelLike {
	var modelFile = directory + "Package.go"
	var bytes, err = osx.ReadFile(modelFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = mod.Parser()
	var model = parser.ParseSource(source)
	return model
}

func generateClasses(
	directory string,
	model mod.ModelLike,
) {
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
