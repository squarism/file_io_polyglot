### Ruby Runs
_2.8ghz Mac Pro Tower, Ruby 2.1.5_
ruby telephones.rb gen_100k.txt  5.83s user 0.94s system 99% cpu 6.777 total
ruby telephones.rb gen_100k.txt  5.84s user 0.95s system 99% cpu 6.805 total
ruby telephones.rb gen_100k.txt  5.82s user 0.95s system 99% cpu 6.770 total

### Perl Runs
_2.8ghz Mac Pro Tower, Perl v5.18.2_
perl telephones.pl gen_100k.txt  0.41s user 0.05s system 99% cpu 0.465 total
perl telephones.pl gen_100k.txt  0.41s user 0.05s system 99% cpu 0.463 total
perl telephones.pl gen_100k.txt  0.41s user 0.06s system 99% cpu 0.465 total

### Go
_2.8ghz Mac Pro Tower, go version go1.3.3 darwin/amd64_
./telephones gen_100k.txt  0.46s user 0.19s system 144% cpu 0.453 total
./telephones gen_100k.txt  0.47s user 0.20s system 146% cpu 0.455 total
./telephones gen_100k.txt  0.46s user 0.21s system 147% cpu 0.453 total

### Go concurrent version
./telephones gen_100k.txt  2.69s user 1.40s system 244% cpu 1.673 total
./telephones gen_100k.txt  2.63s user 1.35s system 245% cpu 1.616 total
./telephones gen_100k.txt  2.66s user 1.43s system 243% cpu 1.684 total
