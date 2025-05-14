# frozen_string_literal: true

module License::Search
  extend ActiveSupport::Concern

  class_methods do
    def search(term)
      return all if term.blank?

      term = term.strip.downcase
      left_joins(:product).where("key LIKE :term OR products.name LIKE :term", term: "%#{term}%")
    end
  end
end
