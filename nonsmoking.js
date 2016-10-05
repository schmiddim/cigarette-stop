#!/usr/local/bin/node
/**
 * example    ./nonsmoking.js 2016-09-19 30 0.285714286
 */


/**
 *  parse a date in yyyy-mm-dd format
 * http://stackoverflow.com/a/2587398/1433869
 * @param input
 * @returns {Date}
 */
function parseDate(input) {
    var parts = input.split('-');
    // new Date(year, month [, day [, hours[, minutes[, seconds[, ms]]]]])
    return new Date(parts[0], parts[1] - 1, parts[2]); // Note: months are 0-based
}

function getSmokingString(dateStopped, cigarettesPerDay, pricePerCigarette) {
    var today = new Date();
    dateStopped = parseDate(dateStopped);
    var timeDiff = Math.abs(today.getTime() - dateStopped.getTime());
    var daysWithoutSmoking = (timeDiff / (1000 * 3600 * 24));
    var cigarettesNotSmoked = Math.ceil(daysWithoutSmoking * cigarettesPerDay);
    var moneySaved = Math.ceil(cigarettesNotSmoked * pricePerCigarette);
    return "You stopped smoking " + Math.floor(daysWithoutSmoking )+ " days ago. Saved " + moneySaved + " Euro and you didnt't smoked " + cigarettesNotSmoked + " cigarettes";

}

console.log(getSmokingString(process.argv[2], process.argv[3], process.argv[4]));
