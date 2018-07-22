require_relative "rails/application.rb"

module Rails
    class << self
        @app_class = nil
        def app_class
            @app_class
        end
    end
end

puts Rails.app_class
