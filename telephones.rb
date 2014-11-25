# Find all telephone numbers as fast as possible

i = 0
File.readlines(ARGV[0], 'r').each do |line|
  if !line.scan(/\(\d+\)\s\d+-\d+/).empty?
    i += 1
    # debug lol
    # puts line
  end
end

puts i
