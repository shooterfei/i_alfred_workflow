set(CMAKE_C_COMPILER clang)
set(CMAKE_CXX_COMPILER clang++)
project(alfred_workflow)
# 通用文件处理
file(
  COPY
  info.plist
  icon.png
  data
  icons
  "List Filter Images"
  third_party DESTINATION
  ${PROJECT_BINARY_DIR}/bin
)
file(COPY script DESTINATION ${PROJECT_BINARY_DIR})

add_subdirectory(go)
add_subdirectory(javascript)
add_subdirectory(python)
add_subdirectory(lua)
add_subdirectory(c)

# 通用处理
add_custom_target(
  base
  COMMAND ./plist_handle.py WORKING_DIRECTORY ${PROJECT_BINARY_DIR}/script
  COMMENT "+++ 通用处理"
)

# 打包为单一文件
add_custom_target(
  package
  COMMAND zip -r tools.alfredworkflow ./* WORKING_DIRECTORY ${PROJECT_BINARY_DIR}/bin
  COMMENT "+++ 打包成alfredworkflow 文件"
  DEPENDS base
)

# 使用go的时候的目录整理
add_custom_target(
  go
  COMMAND mv bin/* ${PROJECT_BINARY_DIR}/bin/ WORKING_DIRECTORY ${PROJECT_BINARY_DIR}/go
)

# javascript 方式实现
add_custom_target(
  javascript
  COMMAND mv bin/* ${PROJECT_BINARY_DIR}/bin/ WORKING_DIRECTORY ${PROJECT_BINARY_DIR}/javascript
)

# python 方式实现
add_custom_target(
  python
  DEPENDS base package
)

# lua 方式实现
add_custom_target(
  lua
  COMMAND mv bin/* ${PROJECT_BINARY_DIR}/bin/ WORKING_DIRECTORY ${PROJECT_BINARY_DIR}/lua
)

# c 方式实现
add_custom_target(
  c
  DEPENDS base package
)

# 安装脚本
add_custom_target(
  install
  DEPENDS package
  COMMAND ./install.py WORKING_DIRECTORY ${PROJECT_BINARY_DIR}/script
)
