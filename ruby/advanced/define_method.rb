class A
    define_method :m do
        puts self
    end
end

A.instance_eval do
    puts self
    define_method :p do #特殊的例外
        puts self
    end
end

a = A.new
puts a
a.m
a.p
