Rails.application.routes.draw do
  resources :boards
  resources :sprints
  resources :issues do
    post :complete, on: :member
    post :move_to_board, on: :member
  end
  resources :users

  root 'sprints#index'
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
