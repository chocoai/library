$i = 0
$num = 5
 
while $i < $num  do
   puts("在循环语句中 i = #$i" )
   $i +=1
end

$i = 0
begin
    puts("在循环语句中 i = #$i" )
    $i +=1
end while $i < $num
