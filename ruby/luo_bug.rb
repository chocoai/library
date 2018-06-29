class Controller
    def sms
        attempt_left ||= 3
        puts "#{attempt_left}"
        raise "sending sms failed"
    rescue
        (attempt_left -= 1) > 0 ? retry : (puts "retry over")
    end
end

c = Controller.new
c.sms
