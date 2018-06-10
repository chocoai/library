colors1 = {"red" => 0xf00, "green" => 0x0f0, "blue" => 0x00f }
puts colors1, colors1.keys(), colors1["red"]
colors1.keys.each do |c|
    puts c.object_id
end

colors2 = {"red": 0xf00, "green": 0x0f0, "blue": 0x00f }
puts colors2, colors2.keys(), colors2[:red]
colors2.keys.each do |d|
    puts d.object_id
end

