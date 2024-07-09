input=tst/input
output=tst/output

echo "Initializing the existing class models:"
rm -r $output/*
for file in `ls $input`; do
	model=${file%?gcmn}
	echo "    $model"
	mkdir $output/$model
	cp $input/$file $output/$model/Package.go
done
echo

echo "Creating the new class models:"
echo "    angle"
bin/initialize class type $output/angle/ angle ""
echo "    array"
bin/initialize generic type $output/array/ array ""
echo "    complex"
bin/initialize class structure $output/complex/ complex ""
echo "    catalog"
bin/initialize generic structure $output/catalog/ catalog ""
echo

echo "Validating the class models:"
for file in `ls $input`; do
	model=${file%?gcmn}
	echo "    $model"
	bin/validate $output/$model/Package.go
done
echo "    angle"
bin/validate $output/angle/Package.go
echo "    array"
bin/validate $output/array/Package.go
echo "    complex"
bin/validate $output/complex/Package.go
echo "    catalog"
bin/validate $output/catalog/Package.go
echo

echo "Formatting the class models:"
for file in `ls $input`; do
	model=${file%?gcmn}
	echo "    $model"
	bin/format $output/$model/Package.go
done
echo "    angle"
bin/format $output/angle/Package.go
echo "    array"
bin/format $output/array/Package.go
echo "    complex"
bin/format $output/complex/Package.go
echo "    catalog"
bin/format $output/catalog/Package.go
echo

echo "Generating the class implementation files:"
for file in `ls $input`; do
	model=${file%?gcmn}
	echo "    $model"
	bin/generate $output/$model
done
echo "    angle"
bin/generate $output/angle/
echo "    array"
bin/generate $output/array/
echo "    complex"
bin/generate $output/complex/
echo "    catalog"
bin/generate $output/catalog/
echo

echo "Running lint on the generated files:"
rm $output/angle/Private.go
rm $output/array/Private.go
golangci-lint run
echo

echo "Done."
