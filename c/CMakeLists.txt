set(CMAKE_CXX_STANDARD 11)


set(BOOST_ROOT /usr/local/Cellar/boost/1.78.0_1)
find_package(Boost COMPONENTS regex system REQUIRED)


if(Boost_FOUND)
  include_directories(${Boost_INCLUDE_DIRS})
  message(STATUS "Boost_INCLUDE_DIRS = ${Boost_INCLUDE_DIRS}.")
  message(STATUS "Boost_LIBRARIES = ${Boost_LIBRARIES}.")
  message(STATUS "Boost_LIB_VERSION = ${Boost_LIB_VERSION}.")
  add_executable(http main.cpp)
  # target_link_libraries(boost ${Boost_LIBRARIES})
endif()
