# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20170129172057) do

  create_table "boards", force: :cascade do |t|
    t.string   "title"
    t.text     "description"
    t.integer  "sprint_id"
    t.integer  "next_board_id"
    t.datetime "created_at",    null: false
    t.datetime "updated_at",    null: false
    t.index ["next_board_id"], name: "index_boards_on_next_board_id"
    t.index ["sprint_id"], name: "index_boards_on_sprint_id"
  end

  create_table "issues", force: :cascade do |t|
    t.string   "title"
    t.text     "description"
    t.integer  "duration_hours"
    t.integer  "sprint_id"
    t.integer  "board_id"
    t.integer  "assigned_to_id"
    t.datetime "completed_at"
    t.datetime "created_at",     null: false
    t.datetime "updated_at",     null: false
    t.index ["assigned_to_id"], name: "index_issues_on_assigned_to_id"
    t.index ["board_id"], name: "index_issues_on_board_id"
    t.index ["sprint_id"], name: "index_issues_on_sprint_id"
  end

  create_table "sprints", force: :cascade do |t|
    t.string   "title"
    t.text     "description"
    t.datetime "due_to"
    t.datetime "completed_at"
    t.datetime "created_at",   null: false
    t.datetime "updated_at",   null: false
  end

  create_table "users", force: :cascade do |t|
    t.string   "username"
    t.string   "name"
    t.string   "avatar_url"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

end
