# frozen_string_literal: true

class DevelopersController < ApplicationController
  before_action :set_developer, only: %i[show edit update destroy]

  def index
    set_page_and_extract_portion_from Developer.all, ordered_by: { email_address: :asc, id: :desc }
  end

  def show
  end

  def new
    @developer = Developer.new
  end

  def create
    @developer = Developer.new(developer_params)

    if @developer.save
      redirect_to @developer, status: :see_other, notice: "Saved"
    else
      render :new, status: :unprocessable_entity
    end
  end

  def edit
  end

  def update
    if @developer.update(developer_params)
      redirect_to @developer, status: :see_other, notice: "Saved"
    else
      render :edit, status: :unprocessable_entity
    end
  end

  def destroy
    @developer.destroy
    redirect_to developers_url, status: :see_other, notice: "Deleted"
  end

  private

    def set_developer
      @developer = Developer.find(params[:id])
    end

    def developer_params
      params.require(:developer).permit(:email_address, :password, :password_confirmation)
    end
end

