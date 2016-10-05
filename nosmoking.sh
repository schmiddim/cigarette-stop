#!/bin/bash

if [ "$#" -ne 3 ]; then
	echo $0 [date-stopped \(Y-m-D\)] [cigarettes per day] [price / cigarette]
	exit 1
fi
if [ -z "$(which bc)" -o ! -x "$(which bc)" ];then
	echo bc is required for floating point calculations
	exit 1
fi

days=$((($(date +%s)-$(date --date="$1" +%s))/86400))
echo You stopped smoking $days days ago. \
	Saved $(LC_NUMERIC=C printf %.2f $(echo $days*$2*$3 | bc)) Euro \
	and you didn\'t smoked $(($days*$2)) cigarettes.

