<?php
/**
 *
 * example 	php cli.php 2016-09-19 30 0.285714286
 */

if (4 !== $argc) {
	echo 'usage: php '. __FILE__ . ' [date-stopped (Y-m-D)] [cigarettes per day] [price / cigarette]' . PHP_EOL;
	exit;
}

function getSmokingString(\DateTime $dateStopped, $cigarettesPerDay, $pricePerCigarette)
{
	$daysWithoutSmoking = (new \DateTime('now'))->diff($dateStopped)->format('%a');
	$cigarettesNotSmoked = $daysWithoutSmoking * $cigarettesPerDay;
	$moneySaved = $pricePerCigarette * $cigarettesNotSmoked;
	return sprintf('You stopped smoking %d days ago. Saved %d Euro, didnt\'t smoke %d cigarettes', $daysWithoutSmoking, $moneySaved, $cigarettesNotSmoked);
}


$dateStopped = new \DateTime($argv[1]);
$cigarettesPerDay = $argv[2];
$pricePerCigarette = $argv[3];

echo getSmokingString($dateStopped, $cigarettesPerDay, $pricePerCigarette);