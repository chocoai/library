def test1(a, b)
    puts a, b
end
test1 1, 2
test1({a: 1, b: 2}, 100)
test1(100, 200)

def test2(a: 3, b: 4)
    puts a, b
end
test2 a: 1, b: 2
test2
test2 a: 5, b: 6
test2(a: 5, b: 6)
arg = {a: 7, b: 8}
test2 arg
test2(arg)
test2 :a => 9, :b => 10

def test3(a=11, b=12)
    puts a, b
end
test3 a: 1, b: 2
test3
test3 a: 5, b: 6
test3({a: 5, b: 6}, 100)
arg = {a: 7, b: 8}
test3 arg, 100
test3(arg, 100)
test3 :a => 9, :b => 10
