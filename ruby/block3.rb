(1...4).each {|i| puts i.to_s}
puts (1...4).map(&:to_s)
[5, 6, 7, 8].each {|i| puts i.to_s}

class A
    def test
        puts "test a"
    end
end

class B
    def test
        puts "test b"
    end
end

[A.new, B.new].each(&:test)
