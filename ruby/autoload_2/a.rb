$:.unshift File.expand_path("../", __FILE__)
#require "b"

module Rails
  puts "a: Rails"
  class Application
    puts "a: Application"
    autoload :Finisher, "./b.rb"

    puts "a: call Finisher.initializer"
    Finisher.initializer
  end
end
