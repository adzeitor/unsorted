json.extract! board, :id, :title, :description, :sprint_id, :next_board_id, :created_at, :updated_at
json.url board_url(board, format: :json)