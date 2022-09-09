
add_executable(firsttest test.c)
include_directories(firsttest ../sources)
target_link_libraries(firsttest PRIVATE SimpleLib)

add_test(success_test firsttest)
set_tests_properties(success_test PROPERTIES PASS_REGULAR_EXPRESSION "SUCCESS")