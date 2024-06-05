output=tst/output
wiki=../../wiki/go-model-tools.wiki

echo "Removing the old examples:"
rm -r $wiki/Example-*
echo

echo "Exporting the angle examples:"
echo "    angle/Package.md"
echo '```go' >$wiki/Example-Angle-Package.md
cat $output/angle/Package.go >>$wiki/Example-Angle-Package.md
echo '```' >>$wiki/Example-Angle-Package.md
echo "    angle/angle.md"
echo '```go' >$wiki/Example-Angle-Class.md
cat $output/angle/angle.go >>$wiki/Example-Angle-Class.md
echo '```' >>$wiki/Example-Angle-Class.md
echo

echo "Exporting the array examples:"
echo "    array/Package.md"
echo '```go' >$wiki/Example-Array-Package.md
cat $output/array/Package.go >>$wiki/Example-Array-Package.md
echo '```' >>$wiki/Example-Array-Package.md
echo "    array/array.md"
echo '```go' >$wiki/Example-Array-Class.md
cat $output/array/array.go >>$wiki/Example-Array-Class.md
echo '```' >>$wiki/Example-Array-Class.md
echo

echo "Exporting the complex examples:"
echo "    complex/Package.md"
echo '```go' >$wiki/Example-Complex-Package.md
cat $output/complex/Package.go >>$wiki/Example-Complex-Package.md
echo '```' >>$wiki/Example-Complex-Package.md
echo "    complex/complex.md"
echo '```go' >$wiki/Example-Complex-Class.md
cat $output/complex/complex.go >>$wiki/Example-Complex-Class.md
echo '```' >>$wiki/Example-Complex-Class.md
echo

echo "Exporting the catalog examples:"
echo "    catalog/Package.md"
echo '```go' >$wiki/Example-Catalog-Package.md
cat $output/catalog/Package.go >>$wiki/Example-Catalog-Package.md
echo '```' >>$wiki/Example-Catalog-Package.md
echo "    catalog/association.md"
echo '```go' >$wiki/Example-Association-Class.md
cat $output/catalog/association.go >>$wiki/Example-Association-Class.md
echo '```' >>$wiki/Example-Association-Class.md
echo "    catalog/catalog.md"
echo '```go' >$wiki/Example-Catalog-Class.md
cat $output/catalog/catalog.go >>$wiki/Example-Catalog-Class.md
echo '```' >>$wiki/Example-Catalog-Class.md
echo

echo "Done."
