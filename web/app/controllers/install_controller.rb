# frozen_string_literal: true

class InstallController < ApplicationController
  allow_unauthenticated_access
  rate_limit to: 10, within: 1.minute
  before_action :set_user

  def install
  end

  def download
    platform = sanitize_file_name_part(params[:platform])
    arch = sanitize_file_name_part(params[:arch])

    filename = ["twice", platform, arch].compact.join("-")
    file_path = Pathname.new(Rails.root.join("storage", filename))

    if file_path.exist?
      send_file file_path, type: "application/octet-stream", filename: filename
    else
      head :not_found
    end
  end

  def products
    products = @user.products
    products = Product.all if @user.is_a?(Developer)
    set_page_and_extract_portion_from products, ordered_by: { name: :asc, id: :desc }
  end

  private

    def set_user
      @user = User.find_by_license_key!(params[:license_key])
    end

    def sanitize_file_name_part(value)
      return if value.blank?

      File.basename(value).gsub(/[^a-z0-9_\-]/, "").downcase.presence
    end
end
