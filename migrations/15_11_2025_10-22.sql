CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  role TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE workouts (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE plans (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE client_plans (
  id SERIAL PRIMARY KEY,
  trainer_id INT REFERENCES users(id) ON DELETE CASCADE,
  client_id INT REFERENCES users(id) ON DELETE CASCADE,
  plan_id INT REFERENCES plans(id) ON DELETE CASCADE,
  assigned_at TIMESTAMP DEFAULT NOW(),
  status TEXT DEFAULT 'active'  -- 'active', 'completed', 'cancelled'
);

CREATE TABLE trainer_clients (
  trainer_id INT REFERENCES users(id) ON DELETE CASCADE,
  client_id INT REFERENCES users(id) ON DELETE CASCADE,
  PRIMARY KEY (trainer_id, client_id)
);

CREATE TABLE requests (
  id SERIAL PRIMARY KEY,
  from_id INT REFERENCES users(id) ON DELETE CASCADE,
  to_id INT REFERENCES users(id) ON DELETE CASCADE,
  type TEXT NOT NULL,
  status TEXT DEFAULT 'pending',  -- 'pending', 'accepted', 'rejected'
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);