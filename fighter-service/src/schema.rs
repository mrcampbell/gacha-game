table! {
    fighters (id) {
        id -> Text,
        user_id -> Text,
        unit_id -> Text,
        level -> Int4,
    }
}

table! {
    moves (id) {
        id -> Text,
        name -> Text,
        element -> Int4,
        targets -> Int4,
        #[sql_name = "type"]
        type_ -> Int4,
    }
}

table! {
    unit_moves (id) {
        id -> Text,
        unit_id -> Text,
        move_id -> Text,
        level_learned_at -> Int4,
    }
}

table! {
    units (id) {
        id -> Text,
        name -> Text,
        element -> Int4,
        #[sql_name = "type"]
        type_ -> Int4,
    }
}

joinable!(fighters -> units (unit_id));
joinable!(unit_moves -> moves (move_id));
joinable!(unit_moves -> units (unit_id));

allow_tables_to_appear_in_same_query!(fighters, moves, unit_moves, units,);
