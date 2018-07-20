module M
    @@j = "j"
    @i = "i"
    def self.m
        puts self
        puts @@j
        puts @i
    end
end

M.m
puts M.class_variables
puts M.instance_variables
