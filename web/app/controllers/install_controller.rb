# frozen_string_literal: true

class InstallController < ApplicationController
  allow_unauthenticated_access
  before_action :set_license
  rate_limit to: 5, within: 1.minute, if: -> { @license.blank? }
  rate_limit to: 30, within: 1.minute, if: -> { @license.present? }

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

  private

    def set_license
      @user = License.find_by_key!(params[:license_key])
    end

    def sanitize_file_name_part(value)
      return if value.blank?

      File.basename(value).gsub(/[^a-z0-9_\-]/, "").downcase.presence
    end
end
