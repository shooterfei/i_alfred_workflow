file(
  COPY
  app.sh
  DESTINATION
  ${PROJECT_BINARY_DIR}/go/bin
)

# go 编译 指令
execute_process(COMMAND go build -ldflags "-X main.Version=1.0.0" -o ${PROJECT_BINARY_DIR}/go/bin/ ./... WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR})
