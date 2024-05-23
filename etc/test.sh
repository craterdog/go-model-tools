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
	mkdir $output/$model
	cp $input/$file $output/$model/Package.go
	bin/format $output/$model/Package.go
done
echo

echo "Generating the class implementation files:"
for model in `ls $output`; do
	echo "    $model"
	bin/generate $output/$model
done
echo

echo "Initializing a new model file:"
bin/initialize $output/example/ example ""
echo
