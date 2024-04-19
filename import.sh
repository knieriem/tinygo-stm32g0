TINYGO=../tinygo

set -e

if test "x$1" = x; then :
else
	TINYGO=$1
fi

cat files | while read f; do

	if test -z $f; then
		:
	elif echo $f | grep '\.patch$' > /dev/null; then
		src=`echo $f | sed 's,\.patch$,,'`
		git -C $TINYGO diff $src > ,,diff
		if ! cmp -s ,,diff $f; then
			echo updating $f
			mv -f ,,diff $f
		else
			rm -f ,,diff
		fi

	elif ! cmp -s $TINYGO/$f $f; then
		echo updating $f
		cat < $TINYGO/$f > $f
	fi
done
