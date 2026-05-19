require "test_helper"

class PackControllerTest < ActionDispatch::IntegrationTest
  test "should get show" do
    get pack_show_url
    assert_response :success
  end
end
