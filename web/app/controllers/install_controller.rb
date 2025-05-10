# frozen_string_literal: true

class InstallController < ApplicationController
  rate_limit to: 5, within: 1.minute
  before_action :set_customer

  def install
  end

  def download
  end

  private

    def set_customer
      @customer = Customer.find_by_license_key!(params[:license_key])
    end
end
