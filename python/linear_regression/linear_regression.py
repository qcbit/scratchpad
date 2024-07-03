import numpy as np

x = np.array([1, 2, 3, 4, 5])
y = np.array([1, 2, 3, 4, 5])

mean_x = np.mean(x)
mean_y = np.mean(y)

m = np.sum((x - mean_x) * (y - mean_y)) / np.sum((x - mean_x) ** 2 )
c = mean_y - m * mean_x

print(f"y = {m}x + {c}")