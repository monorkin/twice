# frozen_string_literal: true

class WebhooksController < ApplicationController
  def order_created
    head :ok
  end
end
