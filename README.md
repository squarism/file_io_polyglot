# File IO Experiments

Sometimes Perl beats Ruby with I/O speed.  Surprising but not offensive.
Outside the bottom-feeding battles of dynamic languages, what would
something assumed to be "fast" behave like?

Also, this may serve as a nice interview quiz.

## The Experiment

_Maybe other experiments could be added ..._

1. `gen_100k.txt` is a plain text file in utf-8 with unix line
   encodings.  It's 87MB so the `git clone` might be slow.  Sorry.
1. Go through the file as fast as you can and find all the phone
   numbers.  This is not a test of i18n so the phone numbers will all
   look like this: (xxx) xxx-xxxx
1. Print the number of phone numbers found.
1. Verify your result by `tail gen_100k.txt` and looking at the answer.
   When the file was generated, the answer was saved within itself.
1. Time it with `time`.  _(see results section for how)_
1. Add it to `results.txt` in this repo.

I want these languages tested:

* Perl
* Ruby
* Go

Later, I'd like these just for fun:

* PHP - yes, php cli app, woooo
* Scala
* Python

## Results
Pipe time stderr to results.txt.

    { time ruby telephones.rb gen_100k.txt; } 2>> results.txt

Or substitute your script/program there.  You should probably do at least three runs.  We can do averages later.


## The Challenge

IO speed is fixed (maybe).  The real challenge is in program design to get through the file fast enough while asyncing/whatever out to a processing layer.  At least this is what I imagine.  I'm not sure if there's anything like concurrent or parallel IO.

Originally I was measuring with a 1.4MB file but it was producing
sub-second runs in Ruby and that's not enough.  There's a `generator.rb`
file that's included but I think we should all be working off a known
file.  We can change this later if we want (add more lines etc) with the
generator.
