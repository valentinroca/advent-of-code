def calculate_total_distance(left_nums, right_nums):
    left = left_nums.copy()
    right = right_nums.copy()

    left.sort()
    right.sort()

    total_diff = sum(abs(l - r) for l, r in zip(left, right))
    return total_diff


def calculate_similarity_score(left_nums, right_nums):
    right_counts = {}
    for num in right_nums:
        right_counts[num] = right_counts.get(num, 0) + 1

    total_score = sum(num * right_counts.get(num, 0) for num in left_nums)
    return total_score


def main():
    left_nums = []
    right_nums = []

    try:
        with open("input") as file:
            for line in file:
                numbers = line.split()
                if len(numbers) != 2:
                    continue
                num1, num2 = map(int, numbers)
                left_nums.append(num1)
                right_nums.append(num2)

        total_distance = calculate_total_distance(left_nums, right_nums)
        similarity_score = calculate_similarity_score(left_nums, right_nums)

        print(f"Total distance: {total_distance}")
        print(f"Similarity score: {similarity_score}")

    except FileNotFoundError:
        print("Error: Could not open input file")


if __name__ == "__main__":
    main()
