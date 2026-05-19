
Rails.application.config.after_initialize do
  next unless Rails.env.production? || Rails.env.development?
  next unless defined?(SolidQueue::Job)

  unless SolidQueue::Job.where(class_name: "CallBotJob", finished_at: nil).exists?
    CallBotJob.perform_later
    Rails.logger.info "CallBotJob started"
  end
end