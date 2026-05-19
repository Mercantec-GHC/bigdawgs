class TestJob < ApplicationJob
  queue_as :default

  def perform
    Rails.logger.info "TEST JOB WORKED"
    puts "TEST JOB WORKED"
  end
end