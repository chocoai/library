class A
    @@g = "g"
    @j = "j"
    def self.n
        puts @j
        puts @@g
    end
    def m
        @i = "i"
        puts "#@i"
        puts @@g
    end
    def b
        puts "#@i"
        puts @@g
    end
end
a = A.new
a.m
a.b
A.n
puts a.instance_variables
puts A.instance_variables
puts A.class_variables
