require 'rubygems'
require 'plivo'

include Plivo

client = RestClient.new('MAMZI3MJZLNGFKMTK1NZ', 'MDZkOTRkMjY2NDY1Mjc2MmQwODNlMDhhOWZkZjE3')

client.messages.create(
  '+14848747788',
  %w[8618661827982],
  'Hello, Majun!'
)
