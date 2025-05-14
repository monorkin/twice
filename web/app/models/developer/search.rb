# frozen_string_literal: true

module Developer::Search
  extend ActiveSupport::Concern

  class_methods do
    def search(query)
      return all if query.blank?

      query = query.strip.downcase
      where("email_address LIKE :query", query: "%#{query}%")
    end
  end
end
