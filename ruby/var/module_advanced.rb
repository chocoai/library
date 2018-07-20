module M
    @@g = "g"
    @j = "j"
    puts self.singleton_class
    class << self
        puts self
        puts @j # 不可见的
        @i = "i"
        puts @i

        def p
            puts "hello#@i"
            puts "welcome#@@g"
        end
    end
end

M.p
