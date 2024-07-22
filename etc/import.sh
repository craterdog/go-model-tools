input=tst/input

echo "Removing the old example class models:"
rm -r $input/*
echo

echo "Importing the new example class models:"
echo "    ast.gcmn"
cp ../go-model-framework/v4/ast/Package.go $input/ast.gcmn
echo "    grammar.gcmn"
cp ../go-model-framework/v4/grammar/Package.go $input/grammar.gcmn
echo "    collection.gcmn"
cp ../go-collection-framework/v4/collection/Package.go $input/collection.gcmn
echo "    cdsn.gcmn"
cp ../go-grammar-framework/v4/ast/Package.go $input/cdsn.gcmn
echo

echo "Done."
