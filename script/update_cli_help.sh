cd `dirname $0`/..

for file in `find content/cli -name "*.out.txt"`; do
  COMMAND=`basename $file | sed -e "s/drone_\(.*\)\.out\.txt/\1/" | sed -e "s/_/ /g"`
  echo "+ drone $COMMAND --help > $file"
  drone $COMMAND --help > $file
done
