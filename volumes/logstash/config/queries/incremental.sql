db.todos.find({last_update: {$lte: new Date()}, last_update: {$gte: :sql_last_value}})

