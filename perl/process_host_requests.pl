#!/usr/bin/perl

use warnings;
use strict;
use autodie;

if (@ARGV == 0) {
    die "Please provide a log file.";
}

my $logfile = $ARGV[0];
#print $logfile,"\n";
my $output = "records_$logfile";

open OUTPUT, '>', $output;
#print $output,"\n";

my %requests_by_hosts;

while (<>) {
    chomp;
    if (/^((\w*\.*)*)\s.*/){
        $requests_by_hosts{$1} += 1;
    }
}

my @k = keys %requests_by_hosts;
foreach (@k) {
    print OUTPUT "$_ $requests_by_hosts{$_}\n";
}

close OUTPUT;
