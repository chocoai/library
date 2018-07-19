#$:.unshift File.expand_path("../", __FILE__)
require_relative "b.rb"

module M
    class N
        #autoload :B, "./b.rb" # "./b.rb" is good while "b.rb" is bad
        def create
            B.new
        end
    end
end

M::N.new.create
