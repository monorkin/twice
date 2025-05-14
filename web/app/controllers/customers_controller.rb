# frozen_string_literal: true

class CustomersController < ApplicationController
  before_action :set_customer, only: %i[show edit update destroy]

  def index
    customers = Customer.all
    customers = customers.search(params[:search][:query]) if params.dig(:search, :query).present?
    set_page_and_extract_portion_from customers, ordered_by: { email_address: :asc, id: :desc }
  end

  def show
  end

  def new
    @customer = Customer.new
  end

  def create
    @customer = Customer.new(customer_params)

    if @customer.save
      redirect_to @customer, status: :see_other, notice: "Saved"
    else
      render :new, status: :unprocessable_entity
    end
  end

  def edit
  end

  def update
    if @customer.update(customer_params)
      redirect_to @customer, status: :see_other, notice: "Saved"
    else
      render :edit, status: :unprocessable_entity
    end
  end

  def destroy
    @customer.destroy
    redirect_to customers_url, status: :see_other, notice: "Deleted"
  end

  private

    def set_customer
      @customer = Customer.find(params[:id])
    end

    def customer_params
      params.require(:customer).permit(:email_address)
    end
end

