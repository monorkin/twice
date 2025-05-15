# frozen_string_literal: true

class InstallController < ApplicationController
  allow_unauthenticated_access
  before_action :set_license
  rate_limit to: 5, within: 1.minute, if: -> { @license.blank? }
  rate_limit to: 30, within: 1.minute, if: -> { @license.present? }

  def install
    respond_to do |format|
      format.text
      format.json
    end
  end

  def download
    platform = sanitize_file_name_part(params[:platform])
    arch = sanitize_file_name_part(params[:arch])

    filename = ["twice", platform, arch].compact.join("-")
    Rails.logger.info "Download filename: #{filename}"
    file_path = Pathname.new(Rails.root.join("storage", filename))

    if file_path.exist?
      send_file file_path, type: "application/octet-stream", filename: filename
    else
      head :not_found
    end
  end

  private

    def set_license
      @license = License.find_by_key!(params[:license_key])
    end

    def sanitize_file_name_part(value)
      return if value.blank?

      File.basename(value).gsub(/[^a-z0-9_\-]/i, "").downcase.presence
    end
end
