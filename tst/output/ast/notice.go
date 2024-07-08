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

var noticeClass = &noticeClass_{
	// Initialize class constants.
}

// Function

func Notice() NoticeClassLike {
	return noticeClass
}

// CLASS METHODS

// Target

type noticeClass_ struct {
	// Define class constants.
}

// Constructors

func (c *noticeClass_) Make(comment string) NoticeLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(comment):
		panic("The comment attribute is required for each Notice.")
	default:
		return &notice_{
			// Initialize instance attributes.
			class_: c,
			comment_: comment,
		}
	}
}

// Private

func (c *noticeClass_) isUndefined(value any) bool {
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

type notice_ struct {
	// Define instance attributes.
	class_ NoticeClassLike
	comment_ string
}

// Attributes

func (v *notice_) GetClass() NoticeClassLike {
	return v.class_
}

func (v *notice_) GetComment() string {
	return v.comment_
}

// Private
