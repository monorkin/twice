# frozen_string_literal: true

class LicensesController < ApplicationController
  before_action :set_customer
  before_action :set_license, only: %i[destroy]

  def index
    licenses = @customer.licenses
    licenses = licenses.search(params[:search][:query]) if params.dig(:search, :query).present?
    set_page_and_extract_portion_from licenses, ordered_by: { product_id: :asc, key: :asc }
  end

  def new
    @license = License.new(owner: @customer)
  end

  def create
    @license = License.new(license_params)

    if @license.save
      redirect_to customer_path(@license.owner), status: :see_other, notice: "License purchased"
    else
      render :new, status: :unprocessable_entity
    end
  end

  def destroy
    @license.destroy
    redirect_to action: :index, status: :see_other, notice: "License revoked"
  end

  private

    def set_customer
      @customer = Customer.find(params[:customer_id])
    end

    def set_license
      @license = @customer.licenses.find(params[:id])
    end

    def license_params
      params.require(:license).permit(:product_id).with_defaults(owner: @customer)
    end
end

