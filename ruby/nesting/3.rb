module A
    module B
        puts "#{Module.nesting}"
    end
end

module X
    module Y
        puts "#{Module.nesting}"
    end
end

module X::Y
    module A::B
        puts "#{Module.nesting}"
    end
end
