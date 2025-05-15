Rails.application.routes.draw do
  root to: redirect("/customers")

  resources :products
  resources :developers
  resources :customers do
    resources :licenses, except: %i[ show edit update ]
  end

  match "/registry/auth", to: "registry#auth", as: :registry_auth, via: %i[ get post ]

  get "/install/:license_key", to: "install#install", as: :install, defaults: { format: :text }
  get "/install/:license_key/download", to: "install#download", as: :install_download

  post "webhooks/order/created", to: "webhooks#order_created", as: :order_created_webhook

  resource :session
  resources :passwords, param: :token

  get "up" => "rails/health#show", as: :rails_health_check
end
