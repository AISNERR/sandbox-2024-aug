def find_root(nodes: list[int]) -> int:
    picks = set()
    kids = set()
    if len(nodes) == 2:
        return nodes[0]
    i = 0
    while i < len(nodes) - 1:
        picks.add(nodes[i])
        i += 1
        shift = 0
        if nodes[i] > 0:
            kids_count = nodes[i]
            while shift < kids_count:
                i += 1
                kids.add(nodes[i])
                shift += 1
        i += 1

    print(picks)
    print(kids)
    for t in picks:
        if t not in kids:
            return t

    return 0





def check_the_longest_connection():
    for count in range(int(input())):
        for num in range(int(input())):
            input_list = input().split()

        print(input_list)

if __name__ == '__main__':
    test_case1 = [3, 0, 1, 0, 5, 2, 2, 6, 4, 3, 5, 1, 3, 2, 0, 6, 0]
    test_case2 = [1, 0, 2, 1, 1]
    test_case3 = [3, 1, 1, 1, 2, 4, 2, 4, 0, 2, 0]
    testik = [4, 1, 5, 5, 2, 2, 1, 2, 1, 3, 1, 0, 3, 0]
    # print(find_root(test_case1)==3)
    # print(find_root(test_case2) == 2)
    print(find_root(test_case1) == 4)
    print(find_root(testik) == 4)
