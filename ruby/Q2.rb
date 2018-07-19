rails/application.rb
  class Application < Engine
    autoload :Bootstrap,              'rails/application/bootstrap'
    autoload :Configuration,          'rails/application/configuration'
    autoload :DefaultMiddlewareStack, 'rails/application/default_middleware_stack'
    autoload :Finisher,               'rails/application/finisher'
    autoload :Railties,               'rails/engine/railties'
    autoload :RoutesReloader,         'rails/application/routes_reloader'

    class << self
      def inherited(base)
        super
        Rails.app_class = base
        add_lib_to_load_path!(find_root(base.called_from))
        ActiveSupport.run_load_hooks(:before_configuration, base)
      end
