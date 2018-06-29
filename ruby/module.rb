module Majun
  #extend self
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
