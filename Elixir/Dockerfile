FROM elixir:1.16

WORKDIR /app

# Install Hex and Rebar
RUN mix local.hex --force && \
    mix local.rebar --force

# Install dependencies
COPY mix.exs mix.lock ./
RUN mix do deps.get, deps.compile

# Ensure the directory exists
RUN mkdir -p /app/_build/dev/lib/certifi/priv

# Copy all application files
COPY . .

CMD ["elixir", "+S", "24", "-S", "mix", "run", "--no-halt"]