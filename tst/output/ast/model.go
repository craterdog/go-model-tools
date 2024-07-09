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

// CLASS ACCESS

// Reference

var modelClass = &modelClass_{
	// Initialize class constants.
}

// Function

func Model() ModelClassLike {
	return modelClass
}

// CLASS METHODS

// Target

type modelClass_ struct {
	// Define class constants.
}

// Constructors

func (c *modelClass_) Make(
	notice NoticeLike,
	header HeaderLike,
	imports ImportsLike,
	types TypesLike,
	functionals FunctionalsLike,
	classes ClassesLike,
	instances InstancesLike,
	aspects AspectsLike,
) ModelLike {
	// Validate the arguments.
	switch {
	case isUndefined(notice):
		panic("The notice attribute is required for each Model.")
	case isUndefined(header):
		panic("The header attribute is required for each Model.")
	case isUndefined(imports):
		panic("The imports attribute is required for each Model.")
	case isUndefined(types):
		panic("The types attribute is required for each Model.")
	case isUndefined(functionals):
		panic("The functionals attribute is required for each Model.")
	case isUndefined(classes):
		panic("The classes attribute is required for each Model.")
	case isUndefined(instances):
		panic("The instances attribute is required for each Model.")
	case isUndefined(aspects):
		panic("The aspects attribute is required for each Model.")
	default:
		return &model_{
			// Initialize instance attributes.
			class_: c,
			notice_: notice,
			header_: header,
			imports_: imports,
			types_: types,
			functionals_: functionals,
			classes_: classes,
			instances_: instances,
			aspects_: aspects,
		}
	}
}

// INSTANCE METHODS

// Target

type model_ struct {
	// Define instance attributes.
	class_ ModelClassLike
	notice_ NoticeLike
	header_ HeaderLike
	optionalImports_ ImportsLike
	optionalTypes_ TypesLike
	optionalFunctionals_ FunctionalsLike
	classes_ ClassesLike
	instances_ InstancesLike
	optionalAspects_ AspectsLike
	imports_ ImportsLike
	types_ TypesLike
	functionals_ FunctionalsLike
	aspects_ AspectsLike
}

// Attributes

func (v *model_) GetClass() ModelClassLike {
	return v.class_
}

func (v *model_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *model_) GetHeader() HeaderLike {
	return v.header_
}

func (v *model_) GetOptionalImports() ImportsLike {
	return v.optionalImports_
}

func (v *model_) GetOptionalTypes() TypesLike {
	return v.optionalTypes_
}

func (v *model_) GetOptionalFunctionals() FunctionalsLike {
	return v.optionalFunctionals_
}

func (v *model_) GetClasses() ClassesLike {
	return v.classes_
}

func (v *model_) GetInstances() InstancesLike {
	return v.instances_
}

func (v *model_) GetOptionalAspects() AspectsLike {
	return v.optionalAspects_
}

// Private
