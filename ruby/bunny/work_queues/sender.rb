require 'bunny'

connection = Bunny.new
connection.start

channel = connection.create_channel
queue = channel.queue('hello')

puts ARGV.inspect
message = ARGV.empty? ? "Hello World!" : ARGV.join(' ')
queue.publish(message, routing_key: queue.name)
puts "[x] Sent #{message}"

connection.close
