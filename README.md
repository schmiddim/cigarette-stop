Description
===========

shows you the number of days since you stopped smoking, how many cigarettes you didn't smoked and how much money you saved

Requirements
============
- nonsmoking.php need php
- nosmoking.sh just needs bash and bc
- nonsmoking.go can be comiled with go
- nonsmoking.js needs node.js

Usage
=====

>php cli.php [date-stopped (Y-m-D)] [cigarettes per day] [price / cigarette]

or

>nosmoking.sh [date-stopped (Y-m-D)] [cigarettes per day] [price / cigarette]

or

>go run nonsmoking.go [date-stopped (Y-m-D)] [cigarettes per day] [price / cigarette]

or

> node  nonsmoking.js [date-stopped (Y-m-D)] [cigarettes per day] [price / cigarette]

Example
------
>php cli.php 2016-09-19 30 0.285714286
 
2016-09-19 is the day you stopped 30 cigarettes per Day 0.285714286 price for one cigarette
