module P
    @j = "j"
    def p
        puts @j
    end
end

class A
    extend P
end

puts A.p
