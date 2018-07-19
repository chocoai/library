$LOAD_PATH << "."

require "c"

#module M
module N
    class C
        def initialize
            puts "b.rb - M::C"
        end
    end
end
