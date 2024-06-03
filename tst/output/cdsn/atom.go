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

import ()

// CLASS ACCESS

// Reference

var atomClass = &atomClass_{
	// Initialize class constants.
}

// Function

func Atom() AtomClassLike {
	return atomClass
}

// CLASS METHODS

// Target

type atomClass_ struct {
	// Define class constants.
}

// Constructors

func (c *atomClass_) MakeWithGlyph(glyph GlyphLike) AtomLike {
	return &atom_{
		// Initialize instance attributes.
		class_: c,
		glyph_: glyph,
	}
}

func (c *atomClass_) MakeWithIntrinsic(intrinsic string) AtomLike {
	return &atom_{
		// Initialize instance attributes.
		class_: c,
		intrinsic_: intrinsic,
	}
}

// INSTANCE METHODS

// Target

type atom_ struct {
	// Define instance attributes.
	class_ AtomClassLike
	glyph_ GlyphLike
	intrinsic_ string
}

// Attributes

func (v *atom_) GetClass() AtomClassLike {
	return v.class_
}

func (v *atom_) GetGlyph() GlyphLike {
	return v.glyph_
}

func (v *atom_) GetIntrinsic() string {
	return v.intrinsic_
}

// Private
