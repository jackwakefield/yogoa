#!/usr/bin/env ruby

require 'watir'
require 'fileutils'

caps = Selenium::WebDriver::Remote::Capabilities.chrome(
  "loggingPrefs"=>{
    "browser"=>"ALL",
    "performance"=>"ALL"
  }
)
browser = Watir::Browser.new(:chrome, :desired_capabilities => caps, :switches => ['--force-device-scale-factor=1', '--window-position=0,0'])

Dir.chdir(File.dirname($0))

output = File.open("../yogoa_test.go", 'w')
output.write "package yogoa\n\n"

Dir['fixtures/*.html'].each do |file|
  fixture = File.read(file)
  name = File.basename(file, '.*')
  puts "Generate #{name}"

  ltr_fixture = fixture.gsub('start', 'left')
                       .gsub('end', 'right')
                       .gsub('flex-left', 'flex-start')
                       .gsub('flex-right', 'flex-end')

  rtl_fixture = fixture.gsub('start', 'right')
                       .gsub('end', 'left')
                       .gsub('flex-right', 'flex-start')
                       .gsub('flex-left', 'flex-end')

  template = File.open('test-template.html').read
  f = File.open('test.html', 'w')
  f.write sprintf(template, name, ltr_fixture, rtl_fixture, fixture)
  f.close
  FileUtils.copy('test.html', "#{name}.html") if $DEBUG

  browser.goto('file://' + Dir.pwd + '/test.html')
  logs = browser.driver.manage.logs.get(:browser)

  output.write eval(logs[0].message.sub(/^[^"]*/, '')) + "\n\n"

  print logs[1]
end

output.close

File.delete('test.html')
browser.close
