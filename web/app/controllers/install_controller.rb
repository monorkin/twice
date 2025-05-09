# frozen_string_literal: true

class InstallController < ApplicationController
  before_action :set_license
  helper_method :license_key

  def show
  end

  private

    def set_license
      @license = License.find_by_key!(params[:license_key])
    end
end
