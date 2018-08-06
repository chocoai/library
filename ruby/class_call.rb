class A
    def call
        puts "A call"
    end

    def m
        puts "m"
    end
end

a = A.new
a.m
a.call
A.instance_method(:call).bind(a).call
