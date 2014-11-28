#!/usr/bin/env perl

open(IN,"<$ARGV[0]") or die "Cannot open file [$ARGV[0]]: $!\n";

my $phonenumbers = 0;
while(<IN>){
  $phonenumbers++ if /\(\d{3}\) \d{3}\-\d{4}/;
}
print "$phonenumbers\n";

