#结论：class里面的self指的是class本身，而实例方法里面的self指的是实例本身
class C
    puts self
    attr_reader :encoding

    def initialize(name)
        puts "initialize:"
        puts self
        puts "######"
        @encoding = "utf-8"
        @name = name
    end

    def encoding=(value)
        @encoding = value
    end

    def instance_method
        puts self
        puts "instance_method:#@name"
    end

    def call_method
        self.instance_method
        instance_method
    end

    def self.singleton_method_local
        puts "singleton_method"
    end
end

c = C.new("majun")
puts "C:#{C}"
puts "c:#{c}"
c.call_method
puts c.encoding
