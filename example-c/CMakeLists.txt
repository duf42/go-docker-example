cmake_minimum_required(VERSION 3.22)

project(
  GoIntegrationC
  VERSION 1.0
  LANGUAGES C)


add_subdirectory(sources)

# add_executable(MyExample simple_example.cpp)
# target_link_libraries(MyExample PRIVATE MyLibExample)

enable_testing()
add_subdirectory(tests)