=begin
module Majun
  #extend self
  #class << self
  def public_method
    puts "public method"
  end
  def module_method
    puts "module method"
  end
  private
  def private_method_1
    puts "private method 1"
  end
  def private_method_2
    puts "private method 2"
  end
  module_function :module_method, :private_method_1
end

class Kavon
  include Majun
end

o = Kavon.new
Majun.public_method
o.public_method
Majun.module_method
Majun.private_method_1
o.private_method_1
o.private_method_2
=end

module Trig
    #def Trig.test1
    #extend self
    class << self
    def test1
        puts "test1"
    end

    def test2
        puts "test2"
    end

    end
    #module_function :test2
end

Trig.test1
Trig.test2
