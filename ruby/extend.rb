module Base
    def hello
        puts "hello"
    end
end

class Car
    extend Base
end

Car.hello
c = Car.new
c.extend(Base)
c.hello
