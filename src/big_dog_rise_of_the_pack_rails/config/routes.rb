Rails.application.routes.draw do
  get "pack/show"
  get "marketplace/show"
  post "marketplace/create"
  get "marketplace/market_history"
  get "resource_bar/show"
  get "maps/tile_popup"
  post "sessions/logout", to: "sessions#logout"
  resources :building do 
    collection do
      get :show_more
    end
  end
  resource :map, only: %i[show new create]
  resource :session, only: %i[ new create show ]
  resources :users, only: %i[new create]
  root to: "home#index"
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Reveal health status on /up that returns 200 if the app boots with no exceptions, otherwise 500.
  # Can be used by load balancers and uptime monitors to verify that the app is live.
  get "up" => "rails/health#show", as: :rails_health_check

  # Render dynamic PWA files from app/views/pwa/* (remember to link manifest in application.html.erb)
  # get "manifest" => "rails/pwa#manifest", as: :pwa_manifest
  # get "service-worker" => "rails/pwa#service_worker", as: :pwa_service_worker

  # Defines the root path route ("/")
  # root "posts#index"

  
end
