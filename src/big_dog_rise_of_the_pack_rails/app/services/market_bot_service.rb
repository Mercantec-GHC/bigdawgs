class MarketBotService
  def self.perform
    user = User.first
    FaradayService.post("/market/bot/trade", connection_type: "MARKETPLACE_BASE_URL", token: user.encoded_json_web_token).body
  end

end