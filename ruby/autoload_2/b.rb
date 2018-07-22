$:.unshift File.expand_path("../", __FILE__)
require "c"

module Rails
  puts "b: Rails"
  class Application
    puts "b: Application"
    module Finisher
      puts "b: Finisher"
      include Initializable
    end
  end
end
