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
	var modelFile = directory + "Package.go"
	if !pathExists(modelFile) {
		fmt.Printf(
			"The model file %v does not exist, so aborting...",
			modelFile,
		)
		return
	}
	var model = parseModel(directory)
	validateModel(model)
	generatePrivate(directory, model)
	generateClasses(directory, model)
}

func generatePrivate(
	directory string,
	model mod.ModelLike,
) {
	var generator = mod.Generator()
	var source = generator.GeneratePrivate(model)
	var bytes = []byte(source)
	var filename = directory + "Private.go"
	if pathExists(filename) {
		fmt.Printf(
			"The %v file exists, so skipping it...",
			filename,
		)
		return
	}
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateClasses(
	directory string,
	model mod.ModelLike,
) {
	var generator = mod.Generator()
	var iterator = model.GetClasses().GetClasses().GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		var name = sts.ToLower(sts.TrimSuffix(
			class.GetDeclaration().GetName(),
			"ClassLike",
		))
		var source = generator.GenerateClass(model, name)
		var bytes = []byte(source)
		var filename = directory + name + ".go"
		if pathExists(filename) {
			fmt.Printf(
				"The class %v exists, so skipping it...",
				filename,
			)
			continue
		}
		var err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
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

func pathExists(path string) bool {
	var _, err = osx.Stat(path)
	if err == nil {
		return true
	}
	if osx.IsNotExist(err) {
		return false
	}
	panic(err)
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
