module Rails
    class << self
        attr_writer :application

        def application
            @application
        end
    end
end

Rails.application = "majun"
puts Rails.application
