class FaradayService

  def self.post_with_success(url, token: nil, body: nil)
    new(url).post_request(token: token, body: body).success?
  end

  attr_reader :url

  def initialize(url)
    @url = url
  end

  def post_request(token: nil, body: nil)
    faraday_connection.post do |request|
      request.headers["Authorization"] = "Bearer #{token}" if token
      request.body = body if body
    end
  end
  

  def faraday_connection
    @faraday_connection ||= Faraday.new(url: url) do |faraday|
      faraday.request :json
      faraday.response :json
      faraday.adapter Faraday.default_adapter
    end
  end

end