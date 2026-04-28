require "test_helper"

class BuildingControllerTest < ActionDispatch::IntegrationTest
  test "should get show" do
    get building_show_url
    assert_response :success
  end

  test "should get update" do
    get building_update_url
    assert_response :success
  end

  test "should get index" do
    get building_index_url
    assert_response :success
  end

  test "should get destroy" do
    get building_destroy_url
    assert_response :success
  end

  test "should get delete" do
    get building_delete_url
    assert_response :success
  end
end
