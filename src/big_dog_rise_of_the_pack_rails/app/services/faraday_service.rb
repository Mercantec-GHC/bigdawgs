class FaradayService

  def self.post(url, connection_type: "ENGINE_BASE_URL", token: nil, body: nil)
    new(url, connection_type).post_request(token: token, body: body)
  end

  def self.fetch_data(url, connection_type: "ENGINE_BASE_URL", token: nil)
    new(url, connection_type).fetch_data(token: token)
  end

  attr_reader :url, :connection_type

  def initialize(url, connection_type)
    @url = url
    @connection_type = connection_type
  end

  def post_request(token: nil, body: nil)
    faraday_connection.post(url) do |request|
      request.headers["Authorization"] = "Bearer #{token}" if token
      request.body = body if body
    end
  end
  

  def faraday_connection
    @faraday_connection ||= Faraday.new(url: ENV.fetch(connection_type, 'localhost:5432')) do |faraday|
      faraday.request :json
      faraday.response :json
      faraday.adapter Faraday.default_adapter
    end
  end

  def fetch_data(token: nil)
    response = faraday_connection.get(url) do |request|
      request.headers["Authorization"] = "Bearer #{token}" if token
    end
    response&.body
  end

end