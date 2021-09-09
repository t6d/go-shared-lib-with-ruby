require 'bundler/setup'
require 'ffi'

module Pi
  extend FFI::Library
  ffi_lib File.expand_path("./main.so", File.dirname(__FILE__))
  attach_function :pi, [], :double
  attach_function :serve, [:int], :void
end

Pi.serve(8000)
