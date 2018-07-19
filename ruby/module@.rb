#在模块中定义变量，并在模块方法中使用
module M
    @loaded = "majun"
    def self.m
        puts "#@loaded"
    end
end

M.m
