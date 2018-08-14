require 'memoist'

class Person
    extend Memoist
    def social_security
        Time.now
    end
    memoize :social_security
end

person = Person.new
puts person.social_security
sleep 5
puts person.social_security
