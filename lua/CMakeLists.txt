file(
  COPY
  app.sh
  lualib
  src
  main.lua
  DESTINATION
  ${PROJECT_BINARY_DIR}/lua/bin
)

