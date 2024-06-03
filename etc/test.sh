input=tst/input
output=tst/output

echo "Validating the class model files:"
for file in `ls $input`; do
	echo "    $file"
	bin/validate $input/$file
done
echo

echo "Formatting the class model files:"
rm -r $output/*
for file in `ls $input`; do
	echo "    $file"
	model=${file%?gcmn}
	rm -rf $output/$model
	mkdir $output/$model
	cp $input/$file $output/$model/Package.go
	bin/format $output/$model/Package.go
done
echo

echo "Generating the class implementation files:"
for file in `ls $input`; do
	model=${file%?gcmn}
	echo "    $model"
	bin/generate $output/$model
done
echo

echo "Initializing a new model file:"
rm -rf $output/model
mkdir $output/model
bin/initialize model $output/model/ model ""
echo

echo "Initializing a new generic file:"
rm -rf $output/generic
mkdir $output/generic
bin/initialize generic $output/generic/ generic ""
echo

echo "Running lint on the generated files:"
golangci-lint run
echo

echo "Done."
