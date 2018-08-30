require 'bunny'

connection = Bunny.new(automatically_recover: false)
connection.start

channel = connection.create_channel
channel.prefetch(1)
queue = channel.queue('hello')

begin
    puts '[*] Waiting for messages. To exit press CTRL+C'
    queue.subscribe(block: false) do |_delivery_info, _properties, body|
        puts "[x] Received #{body}"
    end
rescue Interrupt => _
    connection.close
    exit(0)
end
