module P
    @@g = "g"
    def p
        puts @@g
    end
end

class A
    include P
end
