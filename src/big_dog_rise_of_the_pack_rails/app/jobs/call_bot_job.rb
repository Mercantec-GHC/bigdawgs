class CallBotJob < ApplicationJob
  queue_as :default

  def perform
    MarketBotService.perform

    next_run = rand(1.minute..2.hours)
    self.class.set(wait: next_run).perform_later
  end
end