# frozen_string_literal: true

class InstallController < ApplicationController
  allow_unauthenticated_access
  rate_limit to: 5, within: 1.minute
  before_action :set_user

  def install
  end

  def download
  end

  def products
  end

  private

    def set_user
      @user = User.find_by_license_key!(params[:license_key])
    end
end
