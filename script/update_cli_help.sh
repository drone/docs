# Update drone cli's each command's helps
# When the new version of drone cli has released,
# update the local drone cli and run this script.
#
# USAGE:
#   $ bash script/update_cli_help.sh

drone --version || exit 1
cd `dirname $0`/.. || exit 1

for file in `find content/cli -name "*.out.txt"`; do
  COMMAND=`basename $file | sed -e "s/drone_\(.*\)\.out\.txt/\1/" | sed -e "s/_/ /g"`
  echo "+ drone $COMMAND --help > $file"
  drone $COMMAND --help > $file || exit 1
done
