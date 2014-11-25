# Find all telephone numbers as fast as possible

i = 0
File.readlines('./alice_telephones.txt', 'r').each do |line|
  i += 1 if !line.scan(/\(\d+\)\s\d+-\d+/).empty?
end

puts i
