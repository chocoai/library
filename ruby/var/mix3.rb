module P
    @@g = "g"
    @j = "j"
    def p
        puts @@g
    end
end

class A
    extend P
end
