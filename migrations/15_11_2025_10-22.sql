CREATE TABLE users (
  id SERIAL PRIMARY KEY DEFAULT,
  username TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE workouts (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE plans (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE plan_workout (
  plan_id INT REFERENCES plans(id) ON DELETE CASCADE,
  workout_id INT REFERENCES workouts(id) ON DELETE CASCADE,
  PRIMARY KEY (plan_id, workout_id)
);

CREATE TABLE trainer_clients (
  trainer_id INT REFERENCES users(id) ON DELETE CASCADE,
  client_id INT REFERENCES users(id) ON DELETE CASCADE,
  PRIMARY KEY (plan_id, workout_id)
);

CREATE TABLE requests (
  id SERIAL PRIMARY KEY,
  from_id INT REFERENCES users(id) ON DELETE CASCADE,
  to_id INT REFERENCES users(id) ON DELETE CASCADE,
  type TEXT NOT NULL,
  status TEXT DEFAULT 'pending',  -- 'pending', 'accepted', 'rejected'
  created_at TIMESTAMP DEFAULT NOW()
);