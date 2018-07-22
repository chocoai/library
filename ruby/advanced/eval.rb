class A
    def initialize(name)
        @name = name
    end
end

a = A.new("majun")

#############
A.class_eval do
    def defined_instance_method
        puts "defined_instance_method"
    end

    def self.cm
    end

    class << self
        def sm
        end
    end
end

A.instance_eval do
    def defined_class_method
        puts "defined_class_method"
    end
end

puts A.singleton_methods
a.defined_instance_method

#########################

a.instance_eval do
    puts self
    puts @name

    def a_method_only
        puts "a_method_only-#@name"
    end
end

a.a_method_only
