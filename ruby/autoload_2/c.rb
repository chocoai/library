module Rails
  puts "c: Rails"
  module Initializable
    puts "c: Initializable"
    def self.included(base) #:nodoc:
      base.extend ClassMethods
    end

    module ClassMethods
      puts "c: ClassMethods"
      def initializer()
          puts "c: initializer called"
      end
    end
  end
end
