# frozen_string_literal: true

class DockerRegistry::AccessRequest < Data.define(:type, :name, :actions)
  def self.full_access_to_registry
    [
      new("repository", "*", ["pull", "push"]),
      new("registry", "*", ["pull", "push"]),
    ]
  end

  def self.pull_access_to_repository(repository)
    new("repository", repository, ["pull"])
  end

  def self.parse(scope)
    scope.split("&").map do |part|
      type, name, actions = part.split(":", 3)
      actions = actions.split(",") if actions
      new(type, name, actions)
    end
  end

  def to_s
    [type, name, actions&.join(",")].join(":")
  end
end

