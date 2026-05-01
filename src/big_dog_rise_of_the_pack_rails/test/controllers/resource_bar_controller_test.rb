require "test_helper"

class ResourceBarControllerTest < ActionDispatch::IntegrationTest
  test "should get show" do
    get resource_bar_show_url
    assert_response :success
  end
end
