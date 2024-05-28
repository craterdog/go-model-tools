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
)

func main() {
	var modelFile = retrieveArguments()
	var model = parseModel(modelFile)
	validateModel(model)
	reformatModel(modelFile, model)
}

func retrieveArguments() (modelFile string) {
	if len(osx.Args) < 2 {
		fmt.Println("Usage: format <model-file>")
		osx.Exit(1)
	}
	modelFile = osx.Args[1]
	return modelFile
}

func parseModel(modelFile string) mod.ModelLike {
	var bytes, err = osx.ReadFile(modelFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = mod.Parser()
	var model = parser.ParseSource(source)
	return model
}

func validateModel(model mod.ModelLike) {
	var validator = mod.Validator()
	validator.ValidateModel(model)
}

func reformatModel(
	modelFile string,
	model mod.ModelLike,
) {
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(model)
	var bytes = []byte(source)
	var err = osx.WriteFile(modelFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
