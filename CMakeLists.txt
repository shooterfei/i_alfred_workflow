set(CMAKE_C_COMPILER clang)
set(CMAKE_CXX_COMPILER clang++)
project(alfred_workflow)
# 通用文件处理
file(COPY info.plist icon.png data DESTINATION ${PROJECT_BINARY_DIR}/bin)
file(COPY third_party DESTINATION ${PROJECT_BINARY_DIR}/bin)
add_subdirectory(go)
add_subdirectory(javascript)
add_subdirectory(python)
add_subdirectory(lua)
add_subdirectory(c)


# option(GO "test" ON)


# if(GO)
#   message("test is defined, value:${GO}")
# else()
#   message("test is not defined")
# endif()
