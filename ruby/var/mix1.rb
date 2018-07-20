module P
    @i = "i"
    def p
        puts @i
    end
end

class A
    include P
end

class B
    extend P
end
