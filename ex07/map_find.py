def find_min_area(n, m, k, resources):
    min_area = float('inf')
    best_rectangle = None

    # Пробегаем все возможные прямоугольники
    for x1 in range(1, n + 1):
        for x2 in range(x1, n + 1):
            for y1 in range(1, m + 1):
                for y2 in range(y1, m + 1):
                    # Проверяем, есть ли все ресурсы в текущем прямоугольнике
                    found_resources = set()
                    for (x, y) in resources:
                        if x1 <= x <= x2 and y1 <= y <= y2:
                            found_resources.add((x, y))

                    # Проверяем, если нашли все ресурсы
                    if len(found_resources) == k:
                        area = (x2 - x1 + 1) * (y2 - y1 + 1)
                        if area < min_area:
                            min_area = area
                            best_rectangle = (x1, y1, x2, y2)

    if best_rectangle:
        return (min_area,)
    return None


# Пример входных данных
data = [
    (2, 3, 2, [(1, 3), (2, 2)])
]

for n, m, k, resources in data:
    result = find_min_area(n, m, k, resources)
    print(result)