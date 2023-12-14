TINYGO=../tinygo

set -e

if test "x$1" = x; then :
else
	TINYGO=$1
fi

wd=`pwd`
cd $TINYGO
TINYGO=`pwd`
cd $wd
mkdir dist
cd dist

copy() {
	mkdir -p `dirname $2`
	cat < $1 > $2
}

cat $wd/files | while read f; do

	if test -z $f; then
		:
	elif echo $f | grep '\.patch$' > /dev/null; then
		src=`echo $f | sed 's,\.patch$,,'`
		copy $TINYGO/$src  $src
		if ! patch -s -R --dry-run -p1 < $wd/$f; then
			if ! patch -s -f --dry-run -p1 < $wd/$f; then
				echo cannot apply $f
				exit 1
			else
				echo patching $src
				patch -p1 < $wd/$f
			fi
		else
			echo copying already patched file $src
		fi

	else
		echo copying $f
		copy $wd/$f  $f
	fi
done
