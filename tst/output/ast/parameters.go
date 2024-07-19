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

package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var parametersClass = &parametersClass_{
	// Initialize class constants.
}

// Function

func Parameters() ParametersClassLike {
	return parametersClass
}

// CLASS METHODS

// Target

type parametersClass_ struct {
	// Define class constants.
}

// Constructors

func (c *parametersClass_) Make(
	parameter ParameterLike,
	additionalParameters abs.Sequential[AdditionalParameterLike],
) ParametersLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(parameter):
		panic("The parameter attribute is required by this class.")
	case col.IsUndefined(additionalParameters):
		panic("The additionalParameters attribute is required by this class.")
	default:
		return &parameters_{
			// Initialize instance attributes.
			class_: c,
			parameter_: parameter,
			additionalParameters_: additionalParameters,
		}
	}
}

// INSTANCE METHODS

// Target

type parameters_ struct {
	// Define instance attributes.
	class_ ParametersClassLike
	parameter_ ParameterLike
	additionalParameters_ abs.Sequential[AdditionalParameterLike]
}

// Attributes

func (v *parameters_) GetClass() ParametersClassLike {
	return v.class_
}

func (v *parameters_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *parameters_) GetAdditionalParameters() abs.Sequential[AdditionalParameterLike] {
	return v.additionalParameters_
}

// Private
