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
	col "github.com/craterdog/go-collection-framework/v4/collection"
	ref "reflect"
)

// CLASS ACCESS

// Reference

var instancesClass = &instancesClass_{
	// Initialize class constants.
}

// Function

func Instances() InstancesClassLike {
	return instancesClass
}

// CLASS METHODS

// Target

type instancesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *instancesClass_) Make(
	note string,
	instances col.Sequential[InstanceLike],
) InstancesLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(note):
		panic("The note attribute is required for each Instances.")
	case c.isUndefined(instances):
		panic("The instances attribute is required for each Instances.")
	default:
		return &instances_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			instances_: instances,
		}
	}
}

// Private

func (c *instancesClass_) isUndefined(value any) bool {
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

type instances_ struct {
	// Define instance attributes.
	class_ InstancesClassLike
	note_ string
	instances_ col.Sequential[InstanceLike]
}

// Attributes

func (v *instances_) GetClass() InstancesClassLike {
	return v.class_
}

func (v *instances_) GetNote() string {
	return v.note_
}

func (v *instances_) GetInstances() col.Sequential[InstanceLike] {
	return v.instances_
}

// Private
