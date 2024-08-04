# 5
# 11
# 7 1 4 1 9 1 1 9 1 7 9
# 5
# 1 2 3 4 5
# 5
# 5 5 5 5 5
# 9
# 10 20 10 10 30 10 20 10 40
# 1
# 1000000000


# def longest_sequence(lst):
#     if len(lst) < 2:
#         return len(lst)
#     max_length = 0
#     max_sequence = []
#     current_sequence = []
#     current_digits = []
#     for num in lst:
#         if len(current_digits) == 2:
#             if num in current_digits:
#                 current_sequence.append(num)
#             else:
#                 if len(current_sequence) > max_length:
#                     max_sequence = current_sequence.copy()
#                     max_length = len(current_sequence)
#                     current_sequence = []
#                     current_sequence.append(num)
#                     current_digits = []
#                     current_digits.add(num)
#                 else:
#                     current_digits = []
#                     current_sequence = []
#                     current_sequence.append(num)
#                     current_digits.append(num)
#         else:
#             current_sequence.append(num)
#             current_digits.add(num)
#
#     return max_sequence


def longest_subarray_with_two_distinct(nums):
    left = 0
    count = {}
    max_length = 0
    best_subarray = []

    for right in range(len(nums)):
        count[nums[right]] = count.get(nums[right], 0) + 1

        while len(count) > 2:
            count[nums[left]] -= 1
            if count[nums[left]] == 0:
                del count[nums[left]]
            left += 1

        current_length = right - left + 1
        if current_length > max_length:
            max_length = current_length
            best_subarray = nums[left:right + 1]

    return len(best_subarray)


def check_the_longest_connection():
    for count in range(int(input())):
        for num in range(int(input())):
            input_list = input().split()

        print(input_list)

if __name__ == '__main__':
    # check_the_longest_connection()
    test_case1 = [7, 1, 4, 1, 9, 1, 1, 9, 1, 7, 9]
    test_case2 = [1, 2, 3, 4, 5]
    test_case3 = [5, 5, 5, 5, 5]
    test_case4 = [10, 20, 10, 10, 30, 10, 20, 10, 40]
    test_case5 = [1, 1, 2]
    print(longest_subarray_with_two_distinct(test_case1))
    print(longest_subarray_with_two_distinct(test_case2))
    print(longest_subarray_with_two_distinct(test_case3))
    print(longest_subarray_with_two_distinct(test_case4))
    print(longest_subarray_with_two_distinct(test_case5))
