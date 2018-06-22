def test0(head, options, tail)
    puts "options:", options
    puts "type:", "#{options.class}"
end
test0(1, "name": "majun", "company": "zhenyou", "location": "qingdao", 2)

def test1(*options)
    puts "options:", options
    puts "length:", options.length
    puts "type:", "#{options.class}"
end

test1(1, 2, 3)
test1("name": "majun", "company": "zhenyou", "location": "qingdao")

puts "#################"

def test2(**options)
    puts "options:", options
    puts "type:", "#{options.class}"
end

test2("name": "majun", "company": "zhenyou", "location": "qingdao")
test2(options={"name": "majun", "company": "zhenyou", "location": "qingdao"})

def test3(head, **options, tail)
    puts "options:", options
    puts "type:", "#{options.class}"
end

test3(1, "name": "majun", "company": "zhenyou", "location": "qingdao", 2)
