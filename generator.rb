require 'faker'

size = 100_000

# I'd love to use faker's phone number generation
# but it's too real.  We're not testing regex and
# data handling here.  So I need a MURICAN (bang bang)
# phone number format.

def phone_number
  area = rand(999).to_s.rjust(3, "0")
  prefix = rand(999).to_s.rjust(3, "0")
  line_number = rand(9999).to_s.rjust(4, "0")
  "(#{area}) #{prefix}-#{line_number}"
end

file = File.open("gen_#{size}.txt", 'w')
number_of_phones = 0

size.times do |i|
  file.puts Faker::Lorem.paragraphs(rand(3)+5)

  phone_chance = rand(1_000)
  if phone_chance == 0
    file.puts phone_number
    number_of_phones += 1
   end

  puts i if i % 10_000 == 0
end

file.puts "Number of Phone Numbers written: #{number_of_phones}"
file.close
