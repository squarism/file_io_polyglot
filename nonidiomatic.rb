# Find all telephone numbers as fast as possible
# the non-idiomatic ruby approach

i = 0
infile = File.new(ARGV[0], 'r')
until infile.eof? do
  if /\(\d{3}\) \d{3}-\d{4}/ =~ infile.readline
    i += 1
  end
end
infile.close
puts i
