class Pp
    class << self
        delegate :config, to: :instance
        def instance
            @instance ||= new
        end
    end
    def config
        @config ||= {}
    end
end

class Qq < Pp
    config["name"] = "majun"
end

Qq.config
# {"name"=>"majun"}
