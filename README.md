# File IO Experiments

Sometimes Perl beats Ruby with I/O speed.  Surprising but not offensive.
Outside the bottom-feeding battles of dynamic languages, what would
something assumed to be "fast" behave like?

Also, this may serve as a nice interview quiz.

## The Experiment

_Maybe other experiments could be added ..._

1. `alice_telephone.txt` is a plain text file in utf-8 with unix line
   encodings.
2. Go through the file as fast as you can and find all the phone
   numbers.  This is not a test of i18n so the phone numbers will all
   look like this: (xxx) xxx-xxxx
3. Print the number of phone numbers found.
4. Time it with `time`.
5. Add it to `results.txt` in this repo.

I want these languages tested:

* Perl
* Ruby
* Go

Later, I'd like these just for fun:

* PHP - yes, php cli app, woooo
* Scala
* Python

## The Challenge

IO speed is fixed (maybe).  The real challenge is in program design to get through the file fast enough while asyncing/whatever out to a processing layer.  At least this is what I imagine.  I'm not sure if there's anything like concurrent or parallel IO.

Also this data file is only 1.4MB so hopefully that's enough.

The answer should be 10 phone numbers btw.
