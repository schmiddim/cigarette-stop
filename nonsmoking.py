#!/usr/bin/python

#
# example ./nonsmoking.py 2016-09-19 30 0.285714286
#

from datetime import  datetime
from sys import argv


def get_smoking_string(dateStopped, cigarettesPerDay, pricePerCigarette):
    today = datetime.today()

    daysWithoutSmoking = abs((today - dateStopped).days)
    cigarettesNotSmoked = cigarettesPerDay * daysWithoutSmoking
    moneySaved = round(pricePerCigarette * cigarettesNotSmoked, 2)
    return "You stopped smoking " + str(daysWithoutSmoking) + " days ago. Saved " + str(
        moneySaved) + " Euro and you didnt\'t smoked " + str(cigarettesNotSmoked) + " cigarettes"


dateStopped = datetime.strptime(argv[1], "%Y-%m-%d")
print get_smoking_string(dateStopped, int(argv[2]), float(argv[3]))
