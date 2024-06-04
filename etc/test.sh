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
echo "    model"
mkdir $output/model
bin/initialize model $output/model/ model ""
echo "    generic"
mkdir $output/generic
bin/initialize generic $output/generic/ generic ""
echo

echo "Validating the class models:"
for file in `ls $input`; do
	model=${file%?gcmn}
	echo "    $model"
	bin/validate $output/$model/Package.go
done
echo "    model"
bin/validate $output/model/Package.go
echo "    generic"
bin/validate $output/generic/Package.go
echo

echo "Formatting the class model:"
for file in `ls $input`; do
	model=${file%?gcmn}
	echo "    $model"
	bin/format $output/$model/Package.go
done
echo "    model"
bin/format $output/model/Package.go
echo "    generic"
bin/format $output/generic/Package.go
echo

echo "Generating the class implementation files:"
for file in `ls $input`; do
	model=${file%?gcmn}
	echo "    $model"
	bin/generate $output/$model
done
echo "    model"
bin/generate $output/model/
echo "    generic"
bin/generate $output/generic/
echo

echo "Running lint on the generated files:"
golangci-lint run
echo

echo "Done."
