class FaradayService

  def self.post(url, token: nil, body: nil)
    new(url).post_request(token: token, body: body)
  end

  def self.fetch_data(url, token: nil)
    new(url).fetch_data(token: token)
  end

  attr_reader :url

  def initialize(url)
    @url = url
  end

  def post_request(token: nil, body: nil)
    faraday_connection.post(url) do |request|
      request.headers["Authorization"] = "Bearer #{token}" if token
      request.body = body if body
    end
  end
  

  def faraday_connection
    @faraday_connection ||= Faraday.new(url: ENV.fetch('ENGINE_BASE_URL', 'localhost:3000')) do |faraday|
      faraday.request :json
      faraday.response :json
      faraday.adapter Faraday.default_adapter
    end
  end

  def fetch_data(token: nil)
    response = faraday_connection.get(url) do |request|
      request.headers["Authorization"] = "Bearer #{token}" if token
    end
    response.body
  end

end