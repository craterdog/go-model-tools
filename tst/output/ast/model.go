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
	ref "reflect"
)

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
	case c.isUndefined(notice):
		panic("The notice attribute is required for each Model.")
	case c.isUndefined(header):
		panic("The header attribute is required for each Model.")
	case c.isUndefined(imports):
		panic("The imports attribute is required for each Model.")
	case c.isUndefined(types):
		panic("The types attribute is required for each Model.")
	case c.isUndefined(functionals):
		panic("The functionals attribute is required for each Model.")
	case c.isUndefined(classes):
		panic("The classes attribute is required for each Model.")
	case c.isUndefined(instances):
		panic("The instances attribute is required for each Model.")
	case c.isUndefined(aspects):
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

// Private

func (c *modelClass_) isUndefined(value any) bool {
	switch actual := value.(type) {
	case string:
		return len(actual) > 0
	default:
		var meta = ref.ValueOf(actual)
		return (meta.Kind() == ref.Ptr ||
			meta.Kind() == ref.Interface ||
			meta.Kind() == ref.Slice ||
			meta.Kind() == ref.Map ||
			meta.Kind() == ref.Chan ||
			meta.Kind() == ref.Func) && meta.IsNil()
	}
}

// INSTANCE METHODS

// Target

type model_ struct {
	// Define instance attributes.
	class_ ModelClassLike
	notice_ NoticeLike
	header_ HeaderLike
	imports_ ImportsLike
	types_ TypesLike
	functionals_ FunctionalsLike
	classes_ ClassesLike
	instances_ InstancesLike
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
	return v.imports_
}

func (v *model_) GetOptionalTypes() TypesLike {
	return v.types_
}

func (v *model_) GetOptionalFunctionals() FunctionalsLike {
	return v.functionals_
}

func (v *model_) GetClasses() ClassesLike {
	return v.classes_
}

func (v *model_) GetInstances() InstancesLike {
	return v.instances_
}

func (v *model_) GetOptionalAspects() AspectsLike {
	return v.aspects_
}

// Private
