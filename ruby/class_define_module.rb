class C
    module M
        def a
            puts "a"
        end
    end
    extend M
end
C.a
